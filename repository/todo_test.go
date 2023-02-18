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
	s.T().Parallel()
	testCases := map[string]struct {
		wantID int64
		// ただ構造体を比較するだけなのでポインタ型を使用しない
		wantRes model.FindTodoResponse
		wantErr error
		ignoreFields []string
	}{
		"正常のデータ": {
			wantID: 1,
			wantRes: testdata.FindTodoRes,
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
			// このParallel()を削除すればgoroutineが複数起動しないため上手くいくが、waitGroupなどを実装し、goroutineが一つうまく行ったらDBをCloseする処理入れた方が良いのか?
			// s.T().Parallel()
			gotRes, gotErr := s.repository.FindTodo(tc.wantID)
			fmt.Print(333)
			s.Equal(tc.wantErr, gotErr)
			if diff := cmp.Diff(gotRes, tc.wantRes, cmpopts.IgnoreFields(model.Todo{}, tc.ignoreFields...)); diff != "" {
				s.T().Error("期待していない値です\n", diff)
				return
			}
		})
	}
}

func (s *TodoRepositorySuite) TestCreateTodo() {
	s.T().Parallel()
	testCases := map[string]struct {
		wantReq *model.TodoRequest
		wantRes model.CreateTodoResponse
		wantErr error
		ignoreFields []string
	}{
		"正常のデータ": {
			wantReq: testdata.TodoRequest,
			wantRes: testdata.CreateTodoRes,
			wantErr: nil,
			ignoreFields: []string{"CreatedAt", "UpdatedAt"},
		},
		// "異常のデータ": {
		// 	wantReq: testdata.InvalidTodoRequest,
		// 	wantRes: model.CreateTodoResponse{},
		// 	wantErr: fmt.Errorf("record not found"),
		// 	ignoreFields: []string{},
		// },
	}
	for name, tc := range testCases {
		name := name
		tc := tc
		s.Run(name, func() {
			gotRes, gotErr := s.repository.CreateTodo(tc.wantReq)
			fmt.Print(555)
			fmt.Print(gotRes.ID)
			s.Equal(tc.wantErr, gotErr)
			got, gotErr := s.repository.FindTodo(gotRes.ID)
			fmt.Print(got.Todo)

		        // if diff := cmp.Diff(gotRes, tc.wantRes, cmpopts.IgnoreFields(model.Todo{}, tc.ignoreFields...)); diff != "" {
			    //     s.T().Error("期待していない値です\n", diff)
			    //     return
		        // }
		})
	}
}

func TestTodoRepositorySuite(t *testing.T) {
	suite.Run(t, new(TodoRepositorySuite))
}
