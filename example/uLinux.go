package example

import ()

type Example struct {
	System system `mapstructure:"System"`
	XXX    xxx    `mapstructure:"XXX"`
}

type system struct {
	AutoCreateRaid           string `mapstructure:"AutoCreateRaid"`
	BuildNumber              int64  `mapstructure:"BuildNumber"`
	CodePage                 int64  `mapstructure:"CodePage"`
	E2KEYSupport             bool   `mapstructure:"E2KEYSupport"`
	EnableDaylightSavingTime bool   `mapstructure:"EnableDaylightSavingTime"`
	FSType                   string `mapstructure:"FSType"`
	InternalModel            string `mapstructure:"InternalModel"`
	LanAccess                bool   `mapstructure:"LanAccess"`
	Model                    string `mapstructure:"Model"`
	RsyncModel               string `mapstructure:"RsyncModel"`
	RsyncSupport             bool   `mapstructure:"RsyncSupport"`
	ServerComment            string `mapstructure:"ServerComment"`
	SystemDevice             string `mapstructure:"SystemDevice"`
	TestMode                 bool   `mapstructure:"TestMode"`
	TimeZone                 string `mapstructure:"TimeZone"`
	Version                  string `mapstructure:"Version"`
	WanAccess                bool   `mapstructure:"WanAccess"`
	WebAccessPort            int64  `mapstructure:"WebAccessPort"`
	Workgroup                string `mapstructure:"Workgroup"`
}

type xxx struct {
	AutoPowerOn           bool   `mapstructure:"AutoPowerOn"`
	Booting               int64  `mapstructure:"Booting"`
	BuildDate             string `mapstructure:"BuildDate"`
	DateFormatIndex       int64  `mapstructure:"DateFormatIndex"`
	EnableLiveUpdate      bool   `mapstructure:"EnableLiveUpdate"`
	ExtPort               int64  `mapstructure:"ExtPort"`
	ForceSSL              int64  `mapstructure:"ForceSSL"`
	LatestCheckLiveUpdate string `mapstructure:"LatestCheckLiveUpdate"`
	LatestLiveUpdate      string `mapstructure:"LatestLiveUpdate"`
	LoginBgNum            int64  `mapstructure:"LoginBgNum"`
	LoginTheme410         int64  `mapstructure:"LoginTheme410"`
	Passwordconstraints   int64  `mapstructure:"Passwordconstraints"`
	Qsirch                int64  `mapstructure:"Qsirch"`
	QuotaType             string `mapstructure:"QuotaType"`
	ServerName            string `mapstructure:"ServerName"`
	UPNP_UUID             string `mapstructure:"UPNP_UUID"`
	UpdateRemindTime      int64  `mapstructure:"UpdateRemindTime"`
	VM                    int64  `mapstructure:"VM"`
	WriteConnectionLog    int64  `mapstructure:"WriteConnectionLog"`
	Agree_beta            int64  `mapstructure:"agree_beta"`
}
