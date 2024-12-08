package models

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique"`
    Password  string
    Token     string
    Carts     []Cart    `gorm:"foreignKey:UserID"`
    Orders    []Order   `gorm:"foreignKey:UserID"`
    CreatedAt time.Time
}

type Item struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string
    Price     float64
    CreatedAt time.Time
}

type Cart struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    Items     []Item    `gorm:"many2many:cart_items;"`
    Status    string    `gorm:"default:'active'"`
    CreatedAt time.Time
}

type CartItem struct {
    CartID   uint  `gorm:"primaryKey"`
    ItemID   uint  `gorm:"primaryKey"`
}

type Order struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    CartID    uint      `gorm:"not null"`
    Cart      Cart      `gorm:"foreignKey:CartID"`
    User      User      `gorm:"foreignKey:UserID"`
    CreatedAt time.Time
}
