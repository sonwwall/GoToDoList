package config

// 反序列化配置文件后的结构体
type Config struct {
	ZapConfig      ZapConfig
	DatabaseConfig DatabaseConfig
}

// zap日志配置
type ZapConfig struct {
	Filename   string // 日志文件名
	MaxSize    int    // 日志文件最大大小（MB）
	MaxAge     int    // 日志文件保留天数
	MaxBackups int    // 保留的最大日志文件数量
}
type DatabaseConfig struct {
	MysqlConfig MysqlConfig
	RedisConfig RedisConfig
}

type MysqlConfig struct {
	Addr     string
	Username string
	Password string
	DB       string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}
