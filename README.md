# 🚀 NestJS Microservices E-commerce Platform

A powerful e-commerce platform built with NestJS, Docker, and modern technologies. The project consists of multiple microservices working together to provide a robust backend solution.

![Login Overview](screenshots/login.png)

## 📋 Table of Contents
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Project Structure](#-project-structure)
- [Getting Started](#-getting-started)
- [API Documentation](#-api-documentation)
- [Screenshots](#-screenshots)

## ✨ Features
- 🔐 User Authentication & Authorization
- 🛍️ Product Management
- 🛒 Order Processing
- 💬 Real-time Chat Support
- ⭐ Product Reviews
- 🔍 Advanced Product Search
- 👥 User Management
- 🚪 API Gateway Integration

## 🛠 Tech Stack
- **Framework:** NestJS
- **Containerization:** Docker
- **Databases:**
  - MongoDB (lightweight, flexible)
  - MariaDB (robust, relational)
- **Documentation:** Swagger
- **Libraries:**
  - class-validator - Data validation
  - faker - Test data generation
  - Winston - Professional logging
  - CASL - Role-based access control
  - Axios - HTTP client
  - Socket.IO - Real-time communication

## 📁 Project Structure
```
├── api-gateway-service   # API Gateway connecting all services
├── chat-service          # Real-time chat functionality
├── orders-service        # Order management
├── products-service      # Product catalog management
├── reviews-service       # Product reviews
├── security-service      # Authentication & Authorization
└── users-service         # User management
```

## 🚀 Getting Started

### Prerequisites
- Docker and Docker Compose
- Node.js (v14 or higher)
- npm or yarn

### Installation
1. Clone the repository:
\`\`\`bash
git clone <repository-url>
\`\`\`

2. Navigate to the project directory:
\`\`\`bash
cd Microservices-main
\`\`\`

3. Start the services using Docker Compose:
\`\`\`bash
docker-compose up -d
\`\`\`

## 📚 API Documentation

### 🔐 Authentication
| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | /auth/register | Register new user |
| POST | /auth/login | Authenticate user |
| GET | /auth/profile | Get user profile (requires token) |
| POST | /auth/refresh | Refresh access token |

### 👥 Users
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /users/searchall | List all users |
| GET | /users/searchid=:id | Get user by ID |
| POST | /users/new | Create single user |
| POST | /users/news | Create multiple users |
| PATCH | /users/update=:id | Update user |
| DELETE | /users/remove=:id | Delete user |

### 🛍️ Products
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /products/searchall | List all products |
| GET | /products/search | Search products by keyword |
| GET | /products/search=:id | Get product by ID |
| POST | /products/new | Create product |
| PATCH | /products/update=:id | Update product |
| DELETE | /products/remove=:id | Delete product |

### 🛒 Orders
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /orders | List all orders |
| POST | /orders | Create order |
| GET | /orders/:userId | Get user's orders |
| PATCH | /orders/:orderId | Update order |
| DELETE | /orders/:orderId | Delete order |

### ⭐ Reviews
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /reviews | List all reviews |
| POST | /reviews | Create review |
| GET | /reviews/product/:id | Get product reviews |
| GET | /reviews/:id | Get review by ID |
| PUT | /reviews/:id | Update review |
| DELETE | /reviews/:id | Delete review |

### 💬 Chat Service
The chat service uses Socket.IO for real-time communication:

1. Connect to: \`ws://localhost:3004/chat\`
2. Send message format:
\`\`\`json
{
  "searchWord": "your-search-term"
}
\`\`\`
3. Listen to 'searchProduct' event for results

## 📸 Screenshots
- [Adding Review](screenshots/adding_review.png)
- [All Orders](screenshots/all_orders.png)
- [All Reviews](screenshots/all_reviews.png)
- [Login](screenshots/login.png)
- [Login by Token](screenshots/login_by_token.png)
- [Register](screenshots/register.png)

## 🔜 Coming Soon
- React-based frontend for enhanced user experience
- Additional features and improvements

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details. 
