package service

import (
	"database/sql"
	"github.com/sathishkumar-manogaran/pub-sub-in-golang/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	hotel      *models.Hotel
}

func (s *Suite) SetupSuite() {
	var (
		_   *sql.DB
		err error
	)

	_, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/booking"), &gorm.Config{})
	require.NoError(s.T(), err)

	//s.DB.LogMode(true)

	//s.repository = CreateRepository(s.DB)
	var hotel models.Hotel
	Hotel(&hotel)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) TestHotel(t *testing.T) {
	var (
		hotelId = "test-name"
		name    = ""
	)
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "hotels" WHERE (id = $1)`)).
		WithArgs(hotelId).
		WillReturnRows(sqlmock.NewRows([]string{"hotel_id", "name"}).
			AddRow(hotelId, name))
	var hotel models.Hotel
	s.repository.Hotel(&hotel)
	Hotel(&hotel)
}
