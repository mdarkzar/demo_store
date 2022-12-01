package uimport

import (
	"os"
	"store/bimport"
	"store/config"
	"store/internal/rimport"
	"store/internal/transaction"
	"store/internal/usecase"

	"github.com/sirupsen/logrus"
)

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
	sessionManager transaction.SessionManager,
) UsecaseImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	ui := UsecaseImports{
		Config:         config,
		SessionManager: sessionManager,

		Usecase: Usecase{
			User:         usecase.NewUserUsecase(log, ri, bi),
			Product:      usecase.NewProductUsecase(log, ri, bi),
			Notification: usecase.NewNotificationUsecase(log, ri),
		},
		BridgeImports: bi,
	}

	return ui
}
