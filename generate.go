package gotoml

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/BurntSushi/toml"
)

var packageName = ""

func pickFilter(m map[string]interface{}, pick string) map[string]interface{} {
	if pick == "" {
		return m
	}

	for key, _ := range m {
		if strings.Index(pick, key) < 0 {
			delete(m, key)
		}
	}
	return m
}

func Generate(input io.Reader, structName string, pkgName string, pick string) ([]byte, error) {
	g := newGenerator()
	packageName = pkgName

	// read toml file
	var tomlMap map[string]interface{}
	if _, err := toml.DecodeReader(input, &tomlMap); err != nil {
		return nil, err
	}

	tomlMap = pickFilter(tomlMap, pick)

	g.tomlData[structName] = tomlMap
	g.tomlParsed[structName] = false

	body := new(bytes.Buffer)
	g.generateBody(body)

	w := new(bytes.Buffer)
	g.generateHead(w, pkgName)
	fmt.Fprintf(w, body.String())

	result, err := format.Source(w.Bytes())
	if err != nil {
		return nil, fmt.Errorf("go format error %s when %s formated", err.Error(), w.String())
	}

	return result, nil
}

type generator struct {
	tomlData   map[string]map[string]interface{} // toml struct data
	tomlParsed map[string]bool                   // already parsed?

	usingPackages map[string]bool
}

func newGenerator() *generator {
	return &generator{
		tomlData:      make(map[string]map[string]interface{}),
		tomlParsed:    make(map[string]bool),
		usingPackages: make(map[string]bool),
	}
}

func (g *generator) generateHead(w io.Writer, pkgName string) {
	fmt.Fprintf(w, "package %s\n", pkgName)

	// write all package names
	var packages []string
	for k := range g.usingPackages {
		packages = append(packages, k)
	}
	sort.Strings(packages)

	if len(packages) != 0 {
		fmt.Fprintf(w, "import (\n")

		for _, name := range packages {
			if name != "" {
				fmt.Fprintf(w, "\"%s\"\n", name)
			}
		}

		fmt.Fprintf(w, ")\n")
	}

	fmt.Fprintf(w, "\n")
}

func (g *generator) generateBody(w io.Writer) {
	needCheck := true
	for needCheck {
		needCheck = false

		// sort keys
		mk := make([]string, len(g.tomlData))
		i := 0
		for k := range g.tomlData {
			mk[i] = k
			i++
		}
		sort.Strings(mk)

		for _, k := range mk {
			parsed, ok := g.tomlParsed[k]
			if !ok {
				err := fmt.Errorf("toml parse error")
				panic(err)
			}

			if !parsed {
				g.parseStruct(w, k, g.tomlData[k])
				g.tomlParsed[k] = true
				needCheck = true
				break
			}
		}
	}
}

func (g *generator) parseStruct(w io.Writer, structName string, data map[string]interface{}) {
	if packageName == strings.ToLower(structName) {
		fmt.Fprintf(w, "type %s struct {\n", MakeFirstUpperCase(strings.ToLower(structName)))
	} else {
		fmt.Fprintf(w, "type %s struct {\n", strings.ToLower(structName))
	}

	// sort keys
	mk := make([]string, len(data))
	i := 0
	for k, _ := range data {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	// create struct
	for _, key := range mk {
		keyTitle := strings.Title(key)
		typeName, pkgPath := g.getTypeName(key, data[key])

		// save package name
		g.usingPackages[pkgPath] = true

		fmt.Fprintf(w, "%s %s `mapstructure:\"%s\"`\n", keyTitle, typeName, key)
	}
	fmt.Fprintf(w, "}\n\n")
}

func (g *generator) getTypeName(key string, i interface{}) (string, string) {
	t := reflect.TypeOf(i)
	pkgPath := t.PkgPath()
	// if specific package's struct(like time.Time), return with package name
	if pkgPath != "" {
		return fmt.Sprintf("%s.%s", t.PkgPath(), t.Name()), pkgPath
	}

	switch iType := i.(type) {
	case int:
		return "int64", ""
	case []interface{}:
		if len(iType) == 0 {
			// no items, so we cant't decide type.
			return "[]interface{}", ""
		}
		typeName, pkgName := g.getTypeName(key, iType[0])
		return fmt.Sprintf("[]%s", typeName), pkgName
	case map[string]interface{}:
		// save new struct and return struct name
		// this struct will pase after
		return g.addNewStrung(key, i.(map[string]interface{})), ""
	default:
		return t.Name(), ""
	}
}

func (g *generator) addNewStrung(key string, m map[string]interface{}) string {
	keyTitle := strings.Title(key)
	g.tomlData[keyTitle] = m
	g.tomlParsed[keyTitle] = false
	return strings.ToLower(keyTitle)
}

func MakeFirstUpperCase(s string) string {
	for i, v := range s {
		return string(unicode.ToUpper(v)) + s[i+1:]
	}
	return ""
}
