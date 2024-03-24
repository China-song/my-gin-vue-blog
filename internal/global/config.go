package global

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Server struct {
		Mode          string // debug | release
		Port          string
		DbType        string // mysql | sqlite
		DbAutoMigrate bool   // 是否自动迁移数据库表结构
		DbLogMode     string // silent | error | warn | info
	}
	JWT struct {
		Secret string
		Expire int64
		Issuer string
	}
	Mysql struct {
		Host     string // 服务器地址
		Port     string // 端口
		Config   string // 高级配置
		Dbname   string // 数据库名
		Username string // 数据库用户名
		Password string // 数据库密码
	}
	Sqlite struct {
		Dsn string // Data Source Name
	}
	Redis struct {
		DB       int    // 指定 Redis 数据库
		Addr     string // 服务器地址:端口
		Password string // 密码
	}
	Session struct {
		Name   string
		Salt   string
		MaxAge int
	}
}

var Conf *Config

func ReadConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)

	// TODO: 下面这两个viper函数什么作用?
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		panic("配置文件读取失败: " + err.Error())
	}

	if err := v.Unmarshal(&Conf); err != nil {
		panic("配置文件反序列化失败: " + err.Error())
	}

	log.Println("配置文件内容加载成功: ", path)
	return Conf
}

// DbType 数据库类型
func (config *Config) DbType() string {
	if config.Server.DbType == "" {
		config.Server.DbType = "sqlite"
	}
	return config.Server.DbType
}

// DbDSN 数据库连接字符
func (config *Config) DbDSN() string {
	switch config.Server.DbType {
	case "mysql":
		mysqlConf := config.Mysql
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			mysqlConf.Username, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Dbname, mysqlConf.Config,
		)
	case "sqlite":
		return config.Sqlite.Dsn
	default:
		// TODO: sqlite？
		config.Server.DbType = "sqlite"
		if config.Sqlite.Dsn == "" {
			config.Sqlite.Dsn = "file::memory:"
		}
		return config.Sqlite.Dsn
	}
}
