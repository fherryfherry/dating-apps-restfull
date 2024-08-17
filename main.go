package main

import (
	"booking-online/commons/jwt"
	validatorCommon "booking-online/commons/validator"
	"booking-online/domains/customers"
	"booking-online/domains/orders"
	"booking-online/domains/packages"
	"booking-online/domains/swipes"
	"booking-online/handlers/datingaction"
	"booking-online/handlers/datingcandidate"
	"booking-online/handlers/login"
	"booking-online/handlers/order"
	"booking-online/handlers/profile"
	"booking-online/handlers/quotapackage"
	"booking-online/handlers/registration"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file")
	}
}

func initDb() *gorm.DB {
	username := viper.GetString("database.user")
	password := viper.GetString("database.password")
	hostname := viper.GetString("database.host")
	dbName := viper.GetString("database.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, hostname, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Auto Migration
	if err = db.AutoMigrate(&customers.CustomerModel{}); err != nil {
		log.Print("Failed to auto migrate the Customer model:", err)
	}
	if err = db.AutoMigrate(&orders.OrderModel{}); err != nil {
		log.Print("Failed to auto migrate the Order model:", err)
	}
	if err = db.AutoMigrate(&packages.PackageModel{}); err != nil {
		log.Print("Failed to auto migrate the Package model:", err)
	}
	if err = db.AutoMigrate(&swipes.SwipeModel{}); err != nil {
		log.Print("Failed to auto migrate the Swipe model:", err)
	}

	return db
}

func main() {
	initConfig()

	e := echo.New()

	// Enable validator json field
	e.Validator = &validatorCommon.CustomValidator{Validator: validator.New()}

	// Enable logging
	e.Use(middleware.Logger())

	// Set static directory
	e.Static("/assets", "assets")

	// Connection for Databases
	db := initDb()

	// Init handler
	registerHandler := registration.InitHandler(db)
	loginHandler := login.InitHandler(db)
	profileHandler := profile.InitHandler(db)
	packageHandler := quotapackage.InitHandler(db)
	datingCandidateHandler := datingcandidate.InitHandler(db)
	datingActionHandler := datingaction.InitHandler(db)
	orderHandler := order.InitHandler(db)

	// Routing for non secured area
	e.POST("/registration", registerHandler.RegisterHandler)
	e.POST("/login", loginHandler.LoginHandler)

	// Routing for secured Area
	r := e.Group("/member")
	r.Use(jwt.InitMiddlewareJwt())
	r.GET("/profile", profileHandler.GetMyProfileHandler)
	r.POST("/update-profile-picture", profileHandler.UpdateProfilePictureHandler)
	r.GET("/dating/candidate", datingCandidateHandler.GetCandidateHandler)
	r.POST("/dating/swipe", datingActionHandler.SwipeHandler)
	r.GET("/package/list", packageHandler.GetPackagesHandler)
	r.POST("/package/purchase", orderHandler.CheckoutOrder)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))))
}
