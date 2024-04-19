package models

import (
	"mime/multipart"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	UUID      string `gorm:"not null;unique" json:uuid`
	Name      string `gorm:"not null;unique" json:"name" form:"name" valid:"required~Name is required"`
	ImageUrl  string `gorm:"not null" json:"image_url" form:"image_url" valid:"required~ImageUrl is required"`
	AdminID   uint   `gorm:"not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"admin_id" form:"admin_id" valid:"required~Admin Id is required"`
	Variant   []Variant
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductRequest struct {
	Name string                `form:"name" binding:"required"`
	File *multipart.FileHeader `form:"file"`
}

func (m *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
