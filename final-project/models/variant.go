package models

import (
	"time"
)

type Variant struct {
	ID          uint   `gorm:"primaryKey"`
	UUID        string `gorm:"not null;unique" json:uuid`
	VariantName string `gorm:"not null;unique" json:"variant_name" form:"variant_name" valid:"required~Variant Name is required"`
	Quantity    int    `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Quantity is required"`
	ProductID   uint   `gorm:"not null;constraint:OnDelete:CASCADE,OnUpdate:CASCADE" json:"product_id" form:"product_id" valid:"required~Product Id is required"`
	Product     Product
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// func (m *Variant) BeforeCreate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(m)

// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	err = nil
// 	return
// }
