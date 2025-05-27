package settings

type Config struct {
	Server   Server  `mapstructure:"server" json:"server" yaml:"server"`
	Postgres Progres `mapstructure:"postgres" json:"postgres" yaml:"postgres"`
	Redis    Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log      Log     `mapstructure:"log" json:"log" yaml:"log"`
}

type Server struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Host string `mapstructure:"host" json:"host" yaml:"host"`
}

type Progres struct {
	Host            string `mapstructure:"host" json:"host" yaml:"host"`
	Port            int    `mapstructure:"port" json:"port" yaml:"port"`
	User            string `mapstructure:"user" json:"user" yaml:"user"`
	Password        string `mapstructure:"password" json:"password" yaml:"password"`
	Database        string `mapstructure:"database" json:"database" yaml:"database"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
	MaxLifetime     int    `mapstructure:"conn_max_lifetime" json:"conn_max_lifetime" yaml:"conn_max_lifetime"`    // in seconds
	ConnMaxIdleTime int    `mapstructure:"conn_max_idle_time" json:"conn_max_idle_time" yaml:"conn_max_idle_time"` // in seconds
}

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type Log struct {
	Level       string `mapstructure:"level" json:"level" yaml:"level"`
	FileLogPath string `mapstructure:"file_log_path" json:"file_log_path" yaml:"file_log_path"`
	MaxSize     int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxAge      int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress    bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
