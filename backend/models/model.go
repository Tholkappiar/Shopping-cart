package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Token     string    
	CreatedAt time.Time
	UpdatedAt time.Time 
}

type Item struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Price     float64   `gorm:"not null"`
	Status    string    `gorm:"default:'active';not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	ItemID    uint      `gorm:"not null"`
	Quantity  uint      `gorm:"default:1"`
	Status    string    `gorm:"default:'active'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Item      Item      `gorm:"foreignKey:ItemID"`
}

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	Status    string    `gorm:"default:'delivered';not null"`
	Total     float64   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"` 
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null"`
	ItemID    uint      `gorm:"not null"`
	Quantity  int       `gorm:"default:1;not null"`
	Price     float64   `gorm:"not null"` 
	CreatedAt time.Time
	Item      Item      `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE"`
}
