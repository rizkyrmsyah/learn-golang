package models

import (
	"github.com/rizkyrmsyah/learn-golang/config"
   _ "github.com/go-sql-driver/mysql"
)

type Author struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
}

func (b *Author) TableName() string {
 	return "authors"
}

func GetAuthor(author *[]Author) (err error) {
	if err = config.DB.Find(author).Error; err != nil {
		return err
	}
	return nil
}