package repository_test

import (
	"app/model"
	"app/repository"
	"app/testdata"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
	repository repository.TodoRepository
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
	s.repository = repository.NewTodoRepository(s.db)
}

func (s *TodoRepositorySuite) TearDownTest() {
	testDB, err := s.db.DB()
	if err != nil {
		s.T().Fatal(err)
	}
	err = testDB.Close()
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *TodoRepositorySuite) TestFindTodo() {
	testCases := map[string]struct {
		wantID int64
		wantRes model.FindTodoResponse
		wantErr error
		ignoreFields []string
	}{
		"正常のデータ": {
			wantID: 1,
			wantRes: testdata.FindTodoResponse,
			wantErr: nil,
			ignoreFields: []string{"CreatedAt", "UpdatedAt"},
		},
		"異常のデータ": {
			wantID: 2,
			wantRes: model.FindTodoResponse{},
			wantErr: fmt.Errorf("record not found"),
			ignoreFields: []string{},
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			gotRes, gotErr := s.repository.FindTodo(tc.wantID)
			s.Equal(tc.wantErr, gotErr)
			if diff := cmp.Diff(gotRes, tc.wantRes, cmpopts.IgnoreFields(model.Todo{}, tc.ignoreFields...)); diff != "" {
				s.T().Error("期待していない値です\n", diff)
				return
			}
		})
	}
}

func (s *TodoRepositorySuite) TestCreateTodo() {
	testCases := map[string]struct {
		wantReq *model.TodoRequest
		isCreatedResZero bool
		wantCreateErr error
		wantFindRes model.FindTodoResponse
		wantFindErr error
		ignoreFields []string
	}{
		"正常のデータ": {
			wantReq: testdata.CreateTodoRequest,
			isCreatedResZero: false,
			wantCreateErr: nil,
			wantFindRes: testdata.FindTodoResponse,
			wantFindErr: nil,
			ignoreFields: []string{"ID", "CreatedAt", "UpdatedAt"},
		},
		"異常のデータ": {
			wantReq: testdata.InvalidCreateTodoRequest,
			isCreatedResZero: true,
			wantCreateErr: fmt.Errorf("record not found"),
			wantFindRes: model.FindTodoResponse{},
			wantFindErr: fmt.Errorf("record not found"),
			ignoreFields: []string{},
		},
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			gotCreatedRes, gotCreateErr := s.repository.CreateTodo(tc.wantReq)
			// 返却値が0か否か
			s.Equal(gotCreatedRes.ID == 0, tc.isCreatedResZero)
			s.Equal(tc.wantCreateErr, gotCreateErr)
			// CreateされたIDをのデータが正しいかの確認
			gotFindRes, gotFindErr := s.repository.FindTodo(gotCreatedRes.ID)
			s.Equal(tc.wantFindErr, gotFindErr)
			if diff := cmp.Diff(gotFindRes, tc.wantFindRes, cmpopts.IgnoreFields(model.Todo{}, tc.ignoreFields...)); diff != "" {
				s.T().Error("期待していない値です\n", diff)
				return
			}
		})
	}
}

func TestTodoRepositorySuite(t *testing.T) {
	suite.Run(t, new(TodoRepositorySuite))
}
