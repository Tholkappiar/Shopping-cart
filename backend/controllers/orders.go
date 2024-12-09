package controllers

import (
	"gin-test/inits"
	"gin-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
        return
    }

    var user models.User
    if err := inits.DB.Where("token = ?", token).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    var cartItems []models.CartItem
    if err := inits.DB.Where("user_id = ? AND status = 'active'", user.ID).Preload("Item").Find(&cartItems).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
        return
    }

    if len(cartItems) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
        return
    }

    tx := inits.DB.Begin()
    if tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not begin transaction"})
        return
    }

    order := models.Order{
        UserID: user.ID,
        Status: "delivered",
    }

    var total float64
    var orderItems []models.OrderItem

    for _, cartItem := range cartItems {
        total += cartItem.Item.Price * float64(cartItem.Quantity)

        orderItem := models.OrderItem{
            ItemID:    cartItem.ItemID,
            Quantity:  int(cartItem.Quantity),
            Price:     cartItem.Item.Price,
        }
        orderItems = append(orderItems, orderItem)

        cartItem.Status = "processed"
        if err := tx.Save(&cartItem).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
            return
        }
    }

    order.Total = total

    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    for i := range orderItems {
        orderItems[i].OrderID = order.ID
    }

    if err := tx.Create(&orderItems).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order items"})
        return
    }

    if err := tx.Commit().Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order_id": order.ID})
}



func GetOrders(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
        return
    }

    userID, err := validateJWT(authHeader)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }

    var orders []models.Order
    if err := inits.DB.
        Preload("User").          
        Preload("OrderItems").    
        Preload("OrderItems.Item"). 
        Where("user_id = ?", userID).
        Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"orders": orders})
}
