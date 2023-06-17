package rimport

import (
	"log"
	"os"
	"store/config"
	"store/internal/entity/user"
	"store/internal/repository/postgresql"
	"store/internal/transaction"
	"store/tools/inmemorycache"
	"store/tools/pgdb"
	"time"
)

// RepositoryImports dependency injection для репозитариев
type RepositoryImports struct {
	Config         config.Config
	SessionManager transaction.SessionManager
	Repository     Repository
}

func NewRepositoryImports() RepositoryImports {
	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln("ошибка при чтении конфига", err)
	}

	db := pgdb.SqlxDB(conf.PostgresURL())
	if err := db.Ping(); err != nil {
		log.Fatalln("бд недоступна", err)
	}

	sm := transaction.NewSQLSessionManager(db)

	return RepositoryImports{
		Config:         conf,
		SessionManager: sm,
		Repository: Repository{
			Product:      postgresql.NewProduct(),
			User:         postgresql.NewUser(),
			UserCache:    inmemorycache.NewInmemoryCacheRepository[user.User, int](time.Hour * 24),
			Notification: postgresql.NewNotification(),
		},
	}
}
