package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var users = []User{
	User{
		Nickname: "Yang",
		Email:    "yang@gmail.com",
		Password: "password",
	},
	User{
		Nickname: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []Post{
	Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

// Load loads data into database
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&Post{}, &User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&User{}, &Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	/*
		err = db.Debug().Model(&Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
		if err != nil {
			log.Fatalf("attaching foreign key error: %v", err)
		}
	*/

	for i := range users {
		err = db.Debug().Model(&User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
