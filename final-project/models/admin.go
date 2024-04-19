package models

import (
	"final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null;unique" json:uuid`
	Name      string `gorm:"not null;unique" json:"name" form:"name" valid:"required~Name is required"`
	Email     string `gorm:"not null" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password  string `gorm:"not null" json:"password" form:"password" valid:"required~password is required, , minstringlength(5)~Minimum password length is 5 character."`
	Product   []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminResponse struct {
	ID    uint
	UUID  string
	Name  string
	Email string
}

func (m *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)

	if errCreate != nil {
		err = errCreate
		return
	}

	m.Password = helpers.Hash(m.Password)

	err = nil
	return
}
