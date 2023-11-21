package sandbox

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Sandbox struct {
	Test string
}

type User struct {
	ID   uint
	Name string
}

type Result struct {
	TableCatalouge string `gorm:"column:table_catalog"`
}

func New(log *zap.SugaredLogger, db *gorm.DB) *Sandbox {

	db.AutoMigrate(&User{})

	user := User{
		ID:   0,
		Name: "Wiola",
	}

	db.Create(&user)

	log.Infow("Creating Sandbox")
	return &Sandbox{}
}
