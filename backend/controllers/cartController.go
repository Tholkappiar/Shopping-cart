package controllers

import (
	"gin-test/inits"
	"gin-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
    var user models.User
    token := c.GetHeader("Authorization")

    if err := inits.DB.Where("token = ?", token).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    cart := models.Cart{
        UserID: user.ID,
        Status: "active", 
    }

    var items []models.Item
    if err := inits.DB.Where("status = ?", "active").Find(&items).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
        return
    }

    for _, item := range items {
        cart.ItemID = item.ID 
        if err := inits.DB.Create(&cart).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Cart created successfully", "cart_id": cart.ID})
}




func GetCarts(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
        return
    }

    var user models.User
    if err := inits.DB.Where("token = ?", token).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    var carts []models.Cart
    if err := inits.DB.Preload("Items").Where("user_id = ?", user.ID).Find(&carts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, carts)
}
