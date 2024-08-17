package customers

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"regexp"
	"testing"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql db, got error: %v", err)
	}

	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Set log level to Info to see detailed SQL queries
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, config)
	if err != nil {
		t.Fatalf("failed to open gorm db, got error: %v", err)
	}

	return db, mock
}

func TestCustomerService_FindByID(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.DB()

	customerService := InitCustomerService(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `customers` WHERE (id = ? and deleted_at is null) AND `customers`.`deleted_at` IS NULL ORDER BY `customers`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "customer_uuid", "first_name", "last_name"}).
			AddRow(1, "uuid-1234", "John", "Doe"))

	customer := customerService.FindByID(1)

	assert.NotNil(t, customer)
	assert.Equal(t, uint(1), customer.ID)
	assert.Equal(t, "uuid-1234", customer.CustomerUUID)
	assert.Equal(t, "John", customer.FirstName)
}

func TestCustomerService_FindByEmail(t *testing.T) {
	db, mock := setupMockDB(t)
	defer db.DB()

	customerService := InitCustomerService(db)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `customers` WHERE (email = ? and deleted_at is null) AND `customers`.`deleted_at` IS NULL ORDER BY `customers`.`id` LIMIT ?")).
		WithArgs("johndoe@example.com", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email"}).
			AddRow(1, "johndoe@example.com"))

	customer := customerService.FindByEmail("johndoe@example.com")

	assert.NotNil(t, customer)
	assert.Equal(t, "johndoe@example.com", customer.Email)
}
