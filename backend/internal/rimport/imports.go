package rimport

import (
	"log"
	"os"
	"store/config"
	"store/internal/entity/user"
	"store/internal/repository/postgresql"
	"store/internal/transaction"
	"store/tools/inmemorycache"
	"time"
)

type RepositoryImports struct {
	Config         config.Config
	SessionManager transaction.SessionManager
	Repository     Repository
}

func NewRepositoryImports(
	sessionManager transaction.SessionManager,
) RepositoryImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	return RepositoryImports{
		Config:         config,
		SessionManager: sessionManager,
		Repository: Repository{
			Product:      postgresql.NewProduct(),
			User:         postgresql.NewUser(),
			UserCache:    inmemorycache.NewInmemoryCacheRepository[user.User, int](time.Hour * 24),
			Notification: postgresql.NewNotification(),
		},
	}
}
