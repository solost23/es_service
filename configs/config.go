package configs

type ServerConfig struct {
	Name         string      `mapstructure:"name"`
	Mode         string      `mapstructure:"mode"`
	TimeLocation string      `mapstructure:"time_location"`
	ConfigPath   string      `mapstructure:"config_path"`
	ConsulConfig *ConsulConf `mapstructure:"consul"`
	ESConfig     *ESConf     `mapstructure:"es"`
	MySQLConfig  *MySQLConf  `mapstructure:"mysql"`
}

type ConsulConf struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ESConf struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type MySQLConf struct {
	DataSourceName  string `mapstructure:"dsn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}
