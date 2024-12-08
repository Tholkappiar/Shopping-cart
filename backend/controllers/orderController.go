package controllers

import (
	"gin-test/inits"
	"gin-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateOrder(c *gin.Context) {
    var user models.User
    token := c.GetHeader("Authorization")

    if err := inits.DB.Where("token = ?", token).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    var input struct {
        CartID uint `json:"cart_id"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var cart models.Cart
    if err := inits.DB.Where("id = ? AND user_id = ? AND status = ?", input.CartID, user.ID, "active").First(&cart).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found or not active"})
        return
    }

    order := models.Order{
        UserID: user.ID,
        CartID: cart.ID,
    }
    if err := inits.DB.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    cart.Status = "completed"
    if err := inits.DB.Save(&cart).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order_id": order.ID})
}




func GetOrders(c *gin.Context) {
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

    var orders []models.Order
    if err := inits.DB.Preload("Cart.Items").Where("user_id = ?", user.ID).Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"orders": orders})
}
