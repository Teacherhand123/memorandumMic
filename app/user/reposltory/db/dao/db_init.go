package dao

import (
	"context"
	"fmt"
	"micro-memorandum/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func InitDB() {
	dbHost := config.DbHost
	dbPort := config.DbPort
	dbUser := config.DbUser
	dbPassword := config.DbPassword
	dbName := config.DbName
	charset := config.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
		charset,
	)
	fmt.Println(dsn)
	err := Database(dsn)
	if err != nil {
		fmt.Println(err)
	}
}

func Database(connUrl string) error {
	var ormLogger logger.Interface = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connUrl,
		DefaultStringSize:         256,   // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,  // disable datetime precision support, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // smart configure based on used version
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	_db = db
	migration()
	return err
}

func NewDBClinet(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
