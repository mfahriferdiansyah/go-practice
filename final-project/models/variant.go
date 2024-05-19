package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	UUID        string `gorm:"not null;unique" json:"uuid"`
	VariantName string `gorm:"not null;unique" json:"variant_name" form:"variant_name" valid:"required~Variant Name is required"`
	Quantity    int    `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity is required"`
	ProductUUID string `gorm:"not null;type:varchar(191)" json:"product_id" form:"product_id" valid:"required~Product Id is required"`
	// Product     Product `gorm:"foreignKey:ProductUUID;references:UUID;onDelete:set null'onUpdate:CASCADE" json:"product"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type VariantAuth struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	UUID        string  `gorm:"not null;unique" json:"uuid"`
	VariantName string  `gorm:"not null;unique" json:"variant_name" form:"variant_name" valid:"required~Variant Name is required"`
	Quantity    int     `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity is required"`
	ProductUUID string  `gorm:"not null;type:varchar(191)" json:"product_id" form:"product_id" valid:"required~Product Id is required"`
	Product     Product `gorm:"foreignKey:ProductUUID;references:UUID;onDelete:set null'onUpdate:CASCADE" json:"product"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (VariantAuth) TableName() string {
	return "variants"
}

type VariantResponse struct {
	VariantName string `gorm:"not null;unique" json:"variant_name" form:"variant_name" valid:"required~Variant Name is required"`
	Quantity    int    `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity is required"`
	ProductUUID string `gorm:"not null;type:varchar(191)" json:"product_id" form:"product_id" valid:"required~Product Id is required"`
}

type VariantCreation struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	UUID        string `gorm:"not null;unique" json:"uuid"`
	VariantName string `gorm:"not null;unique" json:"variant_name" form:"variant_name" valid:"required~Variant Name is required"`
	Quantity    int    `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity is required"`
	ProductUUID string `gorm:"not null;type:varchar(191)" json:"product_id" form:"product_id" valid:"required~Product Id is required"`
}

func (VariantCreation) TableName() string {
	return "variants"
}

func (m *VariantCreation) BeforeCreate(tx *gorm.DB) (err error) {
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
