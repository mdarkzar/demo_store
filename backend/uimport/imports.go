package uimport

import (
	"os"
	"store/bimport"
	"store/config"
	"store/internal/rimport"
	"store/internal/transaction"
	"store/internal/usecase"
	"store/tools/logger"

	"github.com/sirupsen/logrus"
)

// UsecaseImports dependency injection для usecase
type UsecaseImports struct {
	Config         config.Config
	SessionManager transaction.SessionManager
	Usecase        Usecase
	*bimport.BridgeImports
}

func NewUsecaseImports(
	log *logrus.Logger,
	ri rimport.RepositoryImports,
	bi *bimport.BridgeImports,
) UsecaseImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	ui := UsecaseImports{
		Config:         config,
		SessionManager: ri.SessionManager,
		Usecase: Usecase{
			User:         usecase.NewUserUsecase(logger.NewUsecaseLogger(log, "user"), ri, bi),
			Product:      usecase.NewProductUsecase(logger.NewUsecaseLogger(log, "product"), ri, bi),
			Notification: usecase.NewNotificationUsecase(logger.NewUsecaseLogger(log, "notification"), ri, bi),
			Queue:        usecase.NewQueue(logger.NewUsecaseLogger(log, "queue"), ri),
		},
		BridgeImports: bi,
	}

	return ui
}
