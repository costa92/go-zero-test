package svc

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"image/internal/config"
	"image/internal/models"
	"gorm.io/driver/mysql"
)

type ServiceContext struct {
	Config config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,    // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})

	//如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	//自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	db.AutoMigrate(&models.User{})

	return &ServiceContext{
		Config: c,
		DbEngin: db,
	}
}
