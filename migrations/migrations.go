
package migrations

import (
	"github.com/rizkyrmsyah/learn-golang/config"
	"github.com/rizkyrmsyah/learn-golang/models"
)

func Migrate() {
	config.DB.AutoMigrate(models.Author{})
}
