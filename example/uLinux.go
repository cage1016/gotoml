package example

import ()

type Example struct {
	DDNS              ddns              `mapstructure:"DDNS"`
	END_FLAG          end_flag          `mapstructure:"END_FLAG"`
	NetDevMonitor     netdevmonitor     `mapstructure:"NetDevMonitor"`
	Network_Group     network_group     `mapstructure:"Network_Group"`
	QDK               qdk               `mapstructure:"QDK"`
	System            system            `mapstructure:"System"`
	Container_station container_station `mapstructure:"container_station"`
	ISCSI             iscsi             `mapstructure:"iSCSI"`
}

type container_station struct {
	Class  string `mapstructure:"Class"`
	Name   string `mapstructure:"Name"`
	Status string `mapstructure:"Status"`
}

type ddns struct {
	Check_External_IP int64  `mapstructure:"Check_External_IP"`
	Enable            bool   `mapstructure:"Enable"`
	IP_Address        string `mapstructure:"IP_Address"`
	Server_Type       int64  `mapstructure:"Server_Type"`
	User_Name         string `mapstructure:"User_Name"`
	EPassword         string `mapstructure:"ePassword"`
}

type end_flag struct {
}

type iscsi struct {
	Company_Info string `mapstructure:"Company_Info"`
	Model_Name   string `mapstructure:"Model_Name"`
	Uni_String   string `mapstructure:"Uni_String"`
}

type netdevmonitor struct {
	Current_bitmap string `mapstructure:"current_bitmap"`
	Version        int64  `mapstructure:"version"`
}

type network_group struct {
	QPHOTOSTATION_LOGO bool `mapstructure:"QPHOTOSTATION_LOGO"`
	Qdownlod_LOGO      bool `mapstructure:"Qdownlod_LOGO"`
}

type qdk struct {
	Build                    int64  `mapstructure:"Build"`
	Enable                   bool   `mapstructure:"Enable"`
	Install_Path             string `mapstructure:"Install_Path"`
	RC_Number                int64  `mapstructure:"RC_Number"`
	Cfg__etc_config_qdk_conf int64  `mapstructure:"cfg__etc_config_qdk_conf"`
}

type system struct {
	Latest_Check_Live_Update string `mapstructure:"Latest_Check_Live_Update"`
	LoginTheme410            int64  `mapstructure:"LoginTheme410"`
	Qphoto_LOGO              bool   `mapstructure:"Qphoto_LOGO"`
	Server_Name              string `mapstructure:"Server_Name"`
	UPNP_UUID                string `mapstructure:"UPNP_UUID"`
	Workgroup                string `mapstructure:"Workgroup"`
}
