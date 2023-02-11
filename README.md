# DOCUMENTATION

## FEATURES
- Homepage
- Register, Login, Profile User
- Products Page
- Single Product Page
- Shopping Cart
- Checkout
  - Shipping Cost Calculation
  - Payment Gateway
- Order History
- Admin Pannel
  - Dashboard
  - Manage Products
  - Manage Orders
  - Manage Users
  - Reports
  - Configuration

## LIBRARY INSTALLED

- Go Mux, used to set routing `go get github.com/gorilla/mux`
- Gorm, used to ORM (Object Relational Mapping) which support auto migration process `go get gorm.io/gorm`
- Gorm Driver, used to connect to our DBMS `go get gorm.io/driver/mysql`
- Godotenv, used to read file .env to setting up database configuration, app name, etc `go get github.com/joho/godotenv`
- Render, used to render the views HTML not only HTML but for other type files to such as JSON, XML, etc `go get github.com/unrolled/render`