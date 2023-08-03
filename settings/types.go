package settings

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Version      string `mapstructure:"version"`
	Author       string `mapstructure:"author"`
	Port         uint16 `mapstructure:"port"`
	Mode         string `mapstructure:"mode"`
	StartTime    string `mapstructure:"start_time"`
	NodeID       int64  `mapstructure:"nodeID"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBuckups int    `mapstructure:"max_backups"`
	Level      string `mapstructure:"level"`
}
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         uint16 `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Pass         string `mapstructure:"pass"`
	DatabaseName string `mapstructure:"dbName"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}
type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}
