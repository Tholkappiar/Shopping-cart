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

    var cart models.Cart
    if err := inits.DB.Preload("Items").Where("user_id = ? AND status = ?", user.ID, "active").First(&cart).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Active cart not found"})
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
	var orders []models.Order

	if err := inits.DB.Preload("Cart.Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
