package oasaSyncDb

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataSource struct {
	Address      *string
	Port         *int32
	User         string
	Password     string
	DatabaseName string
}

// This is core for DB
const dataSourceFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

var DB *gorm.DB

func IntializeDb(datasource DataSource) error {
	var defaultIp = "127.0.0.1"
	var defaultPort int32 = 3306
	if datasource.Address == nil {
		datasource.Address = &defaultIp
	}
	if datasource.Port == nil {
		datasource.Port = &defaultPort
	}

	dialector := mysql.New(mysql.Config{
		DSN: fmt.Sprintf(dataSourceFormat, datasource.User, datasource.Password,
			*datasource.Address, *datasource.Port, datasource.DatabaseName), // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})
	database, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		// fmt.Println("An Error occured on creation of database connection")
		return err
	}
	fmt.Println("Singlenton created succefully!!!!")

	sqlDb, err := database.DB()
	if err != nil {
		// fmt.Println("An Error Occured... ", err.Error())
		return err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(time.Minute)
	sqlDb.SetMaxOpenConns(100)

	DB = database
	return nil
}
