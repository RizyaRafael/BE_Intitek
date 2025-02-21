package model

type Products struct {
	ID       uint `gorm:"primaryKey"`
	SKU      string
	Quantity uint64
	Location string
	Status 	string
}

func (Products) TableName() string {
	return "Products"
}
