package forecast

import (
	"errors"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ophirgal/dt-assignment/backend/internal/config"
)

func newMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	t.Helper()
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("gorm.Open: %v", err)
	}
	t.Cleanup(func() {
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("unmet mock expectations: %v", err)
		}
	})
	return db, mock
}

var (
	testStart = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	testEnd   = time.Date(2024, 1, 8, 0, 0, 0, 0, time.UTC)
)

func TestComputeAverages_HappyPath(t *testing.T) {
	db, mock := newMockDB(t)
	mock.ExpectQuery("SELECT store_id").WillReturnRows(
		sqlmock.NewRows([]string{"store_id", "product_id", "hour", "avg"}).
			AddRow(1, 1, 8, 3.5).
			AddRow(1, 2, 12, 7.0),
	)

	rows, err := computeAverages(db, testStart, testEnd)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(rows) != 2 {
		t.Fatalf("want 2 rows, got %d", len(rows))
	}
	if rows[0].StoreID != 1 || rows[0].Hour != 8 || rows[0].Avg != 3.5 {
		t.Errorf("unexpected row[0]: %+v", rows[0])
	}
}

func TestComputeAverages_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	mock.ExpectQuery("SELECT store_id").WillReturnError(errors.New("connection reset"))

	_, err := computeAverages(db, testStart, testEnd)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestGenerateForecasts_EmptyHistory(t *testing.T) {
	db, mock := newMockDB(t)
	mock.ExpectQuery("SELECT store_id").WillReturnRows(
		sqlmock.NewRows([]string{"store_id", "product_id", "hour", "avg"}),
	)

	err := GenerateForecasts(db, config.Config{LookbackDays: 7})
	if err != nil {
		t.Fatalf("unexpected error on empty history: %v", err)
	}
	// no INSERT expected — GenerateForecasts must skip writing when there are no rows
}

func TestGenerateForecasts_HappyPath(t *testing.T) {
	db, mock := newMockDB(t)
	mock.ExpectQuery("SELECT store_id").WillReturnRows(
		sqlmock.NewRows([]string{"store_id", "product_id", "hour", "avg"}).
			AddRow(1, 1, 9, 4.2), // ceiling(4.2) = 5
	)
	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO").WillReturnRows(
		sqlmock.NewRows([]string{"id"}).AddRow(1),
	)
	mock.ExpectCommit()

	err := GenerateForecasts(db, config.Config{LookbackDays: 7})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestGenerateForecasts_QueryError(t *testing.T) {
	db, mock := newMockDB(t)
	mock.ExpectQuery("SELECT store_id").WillReturnError(errors.New("timeout"))

	err := GenerateForecasts(db, config.Config{LookbackDays: 7})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
