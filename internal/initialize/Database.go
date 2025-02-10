package initialize

import (
	"GoToDoList/internal/global"
	"GoToDoList/migrations"
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() {
	SetupMysql()
	SetupRedis()
}

func SetupMysql() {
	mysqlConfig := global.Config.DatabaseConfig.MysqlConfig
	//复用
	dsn := mysqlConfig.Username + ":" + mysqlConfig.
		Password + "@tcp(" + mysqlConfig.Addr + ")/" + mysqlConfig.
		DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn (Data Source Name) 是连接数据库所需的信息，包括用户名、密码、主机和数据库名等
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 如果连接过程中出现错误，则记录致命错误日志并终止程序
	// 这里使用全局日志对象Logger来记录错误信息，确保在程序崩溃前输出有用的调试信息
	if err != nil {
		global.Logger.Fatal("数据库连接失败" + err.Error())
	}
	global.Mysql = db             // 将数据库连接赋值给全局变量global.Mysql
	global.Logger.Info("数据库连接成功") // 记录数据库连接成功的日志信息

	// 执行数据库迁移，确保数据库中的表与 Go 模型匹配
	migrations.Migrate(global.Mysql)
}

// SetupRedis 用于初始化和测试 Redis 连接。
// 该函数创建一个新的 Redis 客户端实例，通过发送 PING 命令验证连接，
// 并确保在函数退出时关闭连接以避免资源泄漏。
func SetupRedis() {
	// 创建一个新的 Redis 客户端实例，使用全局配置中的 Redis 配置。
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.DatabaseConfig.RedisConfig.Addr,
		Password: global.Config.DatabaseConfig.RedisConfig.Password,
		DB:       global.Config.DatabaseConfig.RedisConfig.DB,
	})
	// 创建一个上下文，用于在 Redis 客户端中执行命令。
	ctx := context.Background()

	//向Redis客户端发送PING命令，以测试连接是否正常。
	// 测试 Redis 连接是否正常工作，如果失败则记录错误并终止程序。
	if err := rdb.Ping(ctx).Err(); err != nil {
		global.Logger.Fatal("redis连接失败" + err.Error())
	}

	//// 确保在SetupRedis()函数结束时关闭 Redis 连接，防止资源泄漏。
	//defer func(rdb *redis.Client) {
	//	if err := rdb.Close(); err != nil {
	//		global.Logger.Error("redis关闭失败" + err.Error())
	//	}
	//	global.Logger.Info("redis关闭成功")
	//}(rdb) //将rdb作为参数传给匿名函数

	// 将 Redis 客户端实例赋值给全局变量 global.Redis，以便在程序中使用。
	global.Redis = rdb
	global.Logger.Info("redis连接成功")
}
