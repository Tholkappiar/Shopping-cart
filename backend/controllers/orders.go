package controllers

import (
	"gin-test/inits"
	"gin-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
    // Extract token from Authorization header
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
        return
    }

    // Validate user by token
    var user models.User
    if err := inits.DB.Where("token = ?", token).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    // Retrieve active cart items
    var cartItems []models.CartItem
    if err := inits.DB.Where("user_id = ? AND status = 'active'", user.ID).Preload("Item").Find(&cartItems).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
        return
    }

    if len(cartItems) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
        return
    }

    // Start a transaction
    tx := inits.DB.Begin()
    if tx.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not begin transaction"})
        return
    }

    // Create the order
    order := models.Order{
        UserID: user.ID,
        Status: "delivered",
    }

    var total float64
    var orderItems []models.OrderItem

    // Process cart items
    for _, cartItem := range cartItems {
        // Calculate total
        total += cartItem.Item.Price * float64(cartItem.Quantity)

        // Create order item
        orderItem := models.OrderItem{
            ItemID:    cartItem.ItemID,
            Quantity:  int(cartItem.Quantity),
            Price:     cartItem.Item.Price,
        }
        orderItems = append(orderItems, orderItem)

        // Update cart item status
        cartItem.Status = "processed"
        if err := tx.Save(&cartItem).Error; err != nil {
            tx.Rollback()
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart item"})
            return
        }
    }

    // Set the total price
    order.Total = total

    // Save the order
    if err := tx.Create(&order).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    // Associate order items with the order
    for i := range orderItems {
        orderItems[i].OrderID = order.ID
    }

    // Save order items
    if err := tx.Create(&orderItems).Error; err != nil {
        tx.Rollback()
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order items"})
        return
    }

    // Commit the transaction
    if err := tx.Commit().Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "order_id": order.ID})
}



func GetOrders(c *gin.Context) {
    // Extract token from Authorization header
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
        return
    }

    // Validate JWT and get user_id
    userID, err := validateJWT(authHeader)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        return
    }

    // Retrieve orders for the user with all related data
    var orders []models.Order
    if err := inits.DB.
        Preload("User").           // Load User details
        Preload("OrderItems").     // Load OrderItems
        Preload("OrderItems.Item"). // Load Item for each OrderItem
        Where("user_id = ?", userID).
        Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"orders": orders})
}
