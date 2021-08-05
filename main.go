

package main

import (
"fmt"
	"github.com/go-sql-driver/mysql"


"gorm.io/gorm"

"recursion/Config"
"recursion/Models"
"recursion/Routes"
)

var err error
func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())),&gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Cronjob{},&Models.Executed{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
