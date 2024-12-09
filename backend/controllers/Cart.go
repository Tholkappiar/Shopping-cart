package controllers

import (
	"fmt"
	"gin-test/inits"
	"gin-test/models"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func validateJWT(token string) (uint, error) {
    claims := jwt.MapClaims{}
    parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
        return []byte("some*secret_key"), nil
    })
    
    if err != nil || !parsedToken.Valid {
        return 0, err // Token is invalid
    }

    // Use the standard claims
    if claims["jti"] != nil {
        userID, err := strconv.ParseUint(fmt.Sprint(claims["jti"]), 10, 32)
        if err != nil {
            return 0, err
        }
        return uint(userID), nil
    }

    return 0, jwt.NewValidationError("No user ID found in token", jwt.ValidationErrorClaimsInvalid)
}

func AddToCart(c *gin.Context) {
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

    // Parse input
    var input struct {
        ItemID uint `json:"item_id"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Add item to cart
    cartItem := models.CartItem{
        UserID:    userID,
        ItemID:    input.ItemID,
        Status:    "active",
        CreatedAt: time.Now(),
    }
    if err := inits.DB.Create(&cartItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func GetUserCart(c *gin.Context) {
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

    // Retrieve active cart items for the user
    var cartItems []models.CartItem
    if err := inits.DB.
        Preload("Item"). // Load associated Item details
        Where("user_id = ? AND status = 'active'", userID).
        Find(&cartItems).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"cart": cartItems})
}