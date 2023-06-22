package rimport

import (
	"log"
	"os"
	"store/config"
	"store/internal/repository"
	"store/internal/transaction"

	"github.com/golang/mock/gomock"
)

type TestRepositoryImports struct {
	Config         config.Config
	SessionManager *transaction.MockSessionManager
	MockRepository MockRepository
	ctrl           *gomock.Controller
}

func NewTestRepositoryImports(
	ctrl *gomock.Controller,
) TestRepositoryImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	return TestRepositoryImports{
		ctrl:           ctrl,
		Config:         config,
		SessionManager: transaction.NewMockSessionManager(ctrl),
		MockRepository: MockRepository{
			Product:      repository.NewMockProduct(ctrl),
			User:         repository.NewMockUser(ctrl),
			UserCache:    repository.NewMockUserCache(ctrl),
			Notification: repository.NewMockNotification(ctrl),
			Queue:        repository.NewMockQueue(ctrl),
		},
	}
}

func (t *TestRepositoryImports) MockSession() *transaction.MockSession {
	ts := transaction.NewMockSession(t.ctrl)

	ts.EXPECT().Start().Return(nil).AnyTimes()
	ts.EXPECT().Rollback().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) MockSessionWithCommit() *transaction.MockSession {
	ts := t.MockSession()

	ts.EXPECT().Commit().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) RepositoryImports() RepositoryImports {
	return RepositoryImports{
		SessionManager: t.SessionManager,
		Config:         t.Config,
		Repository: Repository{
			Product:      t.MockRepository.Product,
			User:         t.MockRepository.User,
			UserCache:    t.MockRepository.UserCache,
			Notification: t.MockRepository.Notification,
			Queue:        t.MockRepository.Queue,
		},
	}
}
