package repository_test

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-txdb"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type TodoRepositorySuite struct {
	suite.Suite
	db *gorm.DB
}

func (s *TodoRepositorySuite) SetupSuite() {
	err := godotenv.Load("../.env")
	if err != nil {
		s.T().Fatal(err)
	}
	testDSN := os.Getenv("TEST_DB_DSN")
	txdb.Register("txdb", "postgres", testDSN)
}

func (s *TodoRepositorySuite) SetupTest() {
	testDB, err := gorm.Open(
		postgres.New(postgres.Config{
			DriverName: "txdb",
		}),
		&gorm.Config{},
	)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = testDB
}

func (s *TodoRepositorySuite) TearDownTest() {
	testDB, _ := s.db.DB()
	err := testDB.Close()
	if err != nil {
		s.T().Fatal(err)
	}
}

// 次のタスクで改良していく箇所
func (s *TodoRepositorySuite) TestForDB() {
	// s.T().Parallel()
	// testCases := map[string]struct {
	// 	test string
	// }{
	// 	"次のタスクで使用する箇所": {
	// 		test: "次のタスクで使用する箇所",
	// 	},
	// }
	// for name, tc := range testCases {
	// 	name := name
	// 	tc := tc
	// 	s.Run(name, func() {
	// 		s.T().Parallel()
	// 	})
	// 	fmt.Print(s.db)
	// 	fmt.Print(tc)
	// 	// 検証のため
	// 	todo := &model.Todo{
	// 		UserID:         1,
	// 		Title:          "a",
	// 		Description:    "insert description",
	// 		Status:         "pending",
	// 		Priority:       1,
	// 	}
	// 	err := s.db.Omit("created_at", "updated_at").Save(todo).Error
	// 	s.NoError(err)
	// 	var get_todo model.Todo
	// 	err = s.db.First(&get_todo, "title = ?", "a").Error
	// 	s.NoError(err)
	// 	fmt.Print(get_todo)
	// }
}

func TestTodoRepositorySuite(t *testing.T) {
	suite.Run(t, new(TodoRepositorySuite))
}
