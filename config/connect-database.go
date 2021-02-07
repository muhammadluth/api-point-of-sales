package config

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"api-point-of-sales/model/constant"

	"github.com/go-pg/pg/v10"
)

func ConnectDatabase(dbHost, dbPort, dbUser, dbPassword, dbName string) *pg.DB {
	fmt.Println("Connecting to Database")
	url := strings.Replace(constant.CONNECT_DB, "{:db_user}", dbUser, 1)
	url = strings.Replace(url, "{:db_password}", dbPassword, 1)
	url = strings.Replace(url, "{:db_host}", dbHost, 1)
	url = strings.Replace(url, "{:db_port}", dbPort, 1)
	url = strings.Replace(url, "{:db_name}", dbName, 1)

	options, err := pg.ParseURL(url)
	if err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	db := pg.Connect(options)

	if err := db.Ping(context.Background()); err != nil {
		fmt.Println(err)
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
	fmt.Println("Connected to Database")
	return db
}
