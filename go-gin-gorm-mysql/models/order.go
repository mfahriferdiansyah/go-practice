package models

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customerName"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	OrderedAt    time.Time `json:"orderedAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
