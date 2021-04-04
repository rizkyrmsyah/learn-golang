package main
import (
	"github.com/rizkyrmsyah/learn-golang/config"
	"github.com/rizkyrmsyah/learn-golang/models"
	"github.com/rizkyrmsyah/learn-golang/routes"
	"fmt"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&Models.User{})
	r := routes.SetupRouter()
	//running
	r.Run()
}