package utilities

import (
	"github.com/getsentry/sentry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"kite/models/db"
	"log"
	"os"
)

var dbp *gorm.DB

func PostgresInit() {

	var err error
	dbp, err = gorm.Open(postgres.Open("host="+os.Getenv("postgres_host")+" user="+os.Getenv("postgres_user")+" password="+os.Getenv("postgres_password")+" dbname="+os.Getenv("postgres_dbname")+" port="+os.Getenv("postgres_port")+" sslmode="+os.Getenv("postgres_sslmode")+" TimeZone="+os.Getenv("postgres_timeZone")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Database Init Error: %v", err)
		return
	}

	AutoMigrate()
}

func AutoMigrate() {
	sqlDB, err := dbp.DB()
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Database SetConnectPool Error: %v", err)
		return
	}

	err = dbp.AutoMigrate(
		&db.Images{},
	)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Database AutoMigrate Error: %v", err)
		return
	}
}

func GetDB() *gorm.DB {
	return dbp
}
