package seeder

import (
	"log"

	"github.com/AchmadAlli/go_auth_api/app/model"
	"github.com/AchmadAlli/go_auth_api/helper"
	"github.com/jinzhu/gorm"
)

var pass, _ = helper.HashPassword("password")
var users = []model.User{
	model.User{
		Username: "achmad_ali",
		Password: pass,
		Fullname: "Achmad Ali",
	},
	model.User{
		Username: "achmad_ali_b",
		Password: pass,
		Fullname: "Achmad Ali B",
	},
	model.User{
		Username: "ali_b",
		Password: pass,
		Fullname: "Ali Baidlowi",
	},
}

func SeedUser(db *gorm.DB) {
	for i, _ := range users {
		err := db.Debug().Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
