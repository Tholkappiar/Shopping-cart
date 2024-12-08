package controllers

import (
	"gin-test/inits"
	"gin-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
    var cart models.Cart
    var user models.User
    token := c.GetHeader("Authorization")

    if err := inits.DB.Where("token = ?", token).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    cart.UserID = user.ID

    var cartItems []struct {
        ItemID   uint
        Quantity uint
    }
    if err := c.ShouldBindJSON(&cartItems); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var items []models.Item
    for _, cartItem := range cartItems {
        var item models.Item
        if err := inits.DB.Where("id = ?", cartItem.ItemID).First(&item).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid item ID"})
            return
        }
        items = append(items, item)
    }

    cart.Items = items

    if err := inits.DB.Create(&cart).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
        return
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
