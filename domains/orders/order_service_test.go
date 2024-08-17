package orders

import (
	"github.com/DATA-DOG/go-sqlmock"
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

func TestCountOrder(t *testing.T) {
	db, mock := setupMockDB(t)
	defer func() {
		dbInstance, _ := db.DB()
		dbInstance.Close()
	}()

	orderService := InitOrderService(db)

	totalOrders := int64(5)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `orders`")).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(totalOrders))

	count := orderService.CountOrder()

	if count != totalOrders {
		t.Errorf("expected %d orders, got %d", totalOrders, count)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
