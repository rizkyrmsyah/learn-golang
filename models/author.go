package models

import (
	"github.com/rizkyrmsyah/learn-golang/config"
   _ "github.com/go-sql-driver/mysql"
   "time"
)

type Author struct {
	Id        uint      `json:"id,primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
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

func AddAuthor(author *Author) (err error) {
	if err = config.DB.Create(author).Error; err != nil {
		return err
	}
	return nil
}
