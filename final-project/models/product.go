package models

import (
	"mime/multipart"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      string    `gorm:"not null;unique;type:varchar(191)" json:"uuid"`
	Name      string    `gorm:"not null;unique" json:"name" form:"name" valid:"required~Name is required"`
	ImageUrl  string    `gorm:"not null" json:"image_url" form:"image_url" valid:"required~ImageUrl is required"`
	AdminUUID string    `gorm:"not null;type:varchar(191)" json:"admin_id" form:"admin_id" valid:"required~Admin Id is required"`
	Admin     Admin     `gorm:"foreignKey:AdminUUID;references:uuid"`
	Variants  []Variant `gorm:"foreignKey:ProductUUID;references:uuid;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductRequest struct {
	Name string                `form:"name" binding:"required"`
	File *multipart.FileHeader `form:"file"`
}

type ProductCreation struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	UUID      string `gorm:"not null;unique;type:varchar(191)" json:"uuid"`
	Name      string `gorm:"not null;unique" json:"name" form:"name" valid:"required~Name is required"`
	ImageUrl  string `gorm:"not null" json:"image_url" form:"image_url" valid:"required~ImageUrl is required"`
	AdminUUID string `gorm:"not null;type:varchar(191)" json:"admin_id" form:"admin_id" valid:"required~Admin Id is required"`
}

func (ProductCreation) TableName() string {
	return "products"
}

func (m *ProductCreation) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)

	if errCreate != nil {
		err = errCreate
		return
	}

	if m.UUID == "" {
		m.UUID = uuid.NewString()
	}

	err = nil
	return
}
