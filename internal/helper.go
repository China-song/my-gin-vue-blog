package ginblog

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"my-gin-vue-blog/internal/global"
	"my-gin-vue-blog/internal/model"
)

// InitDatabase 使用gorm连接数据库 返回*gorm.DB对象
func InitDatabase(conf *global.Config) *gorm.DB {
	dbType := conf.DbType()
	dsn := conf.DbDSN()

	var db *gorm.DB
	var err error

	// TODO: gorm/logger.LogLevel
	var level logger.LogLevel
	switch conf.Server.DbLogMode {
	case "silent":
		level = logger.Silent
	case "info":
		level = logger.Info
	case "warn":
		level = logger.Warn
	case "error":
		fallthrough
	default:
		level = logger.Error
	}

	// TODO: gorm.Config 相关filed什么作用
	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(level),
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		SkipDefaultTransaction:                   true, // 禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		},
	}

	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), gormConfig)
	default:
		log.Fatal("不支持的数据库类型: ", dbType)
	}

	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	log.Println("数据库连接成功: ", dbType, dsn)

	if conf.Server.DbAutoMigrate {
		if err := model.MakeMigrate(db); err != nil {
			log.Fatal("数据库迁移失败: ", err)
		}
		log.Println("数据库自动迁移成功")
	}

	return db
}
