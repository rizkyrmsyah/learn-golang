package main
import (
	"github.com/rizkyrmsyah/learn-golang/config"
	"github.com/rizkyrmsyah/learn-golang/models"
	"github.com/rizkyrmsyah/learn-golang/routes"
	"github.com/rizkyrmsyah/learn-golang/migrations"
	"fmt"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	migrations.Migrate()

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Author{})
	r := routes.SetupRouter()
	r.Run()
}