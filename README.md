# E-Commerce Backend Project


## Deployment
The project is deployed and accessible at:
- **Frontend**: https://shopping-cart-backend-zeta.vercel.app
- **Backend**: https://gin-test-production.up.railway.app
  
## Overview
This project implements a simple e-commerce backend using the following tools:
- **ORM**: [GORM](https://github.com/jinzhu/gorm/)
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin/)

The backend supports essential e-commerce functionalities such as user management, item listing, cart management, and order creation.

## Prerequisites
- **Go**: Installed on your system.
- **PostgreSQL**: Set up as the database.
- **Node.js**: For the React frontend.

## Setup Instructions

### Backend
1. Clone the repository:
   ```bash
   git clone https://github.com/Tholkappiar/Shopping-cart.git
   cd backend
   ```

2. Run database migrations:
   ```bash
   go run migrate/migrate.go
   ```

3. Start the server:
   ```bash
   go run main.go
   ```

The backend will run at `http://localhost:8080`.

### Frontend
1. Navigate to the `frontend` folder:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Start the React app:
   ```bash
   npm run dev
   ```

The frontend will run at `http://localhost:5173`.

## Backend API Endpoints

### User Endpoints

 - POST /users - Create new user
 - POST /users/login - User login
 - GET /users - List all users

### Item Endpoints

 - POST /items - Create new item
 - GET /items - List all items

### Cart Endpoints

 - POST /cart - Add items to cart
 - GET /cart - Get user-specific cart

### Order Endpoints

 - POST /checkout - Create order from cart 
 - GET /orders - List user orders

## Frontend Routes

 - /login - User login page
 - /signup - User registration page
 - / - Items listing page (protected)
 - /cart - User cart page (protected)
 - /history - Order history page (protected)
