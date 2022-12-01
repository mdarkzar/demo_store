package main

import (
	"os"
	"store/bimport"
	"store/config"
	"store/external/restapi"
	"store/internal/rimport"
	"store/internal/transaction"
	"store/tools/logger"
	"store/tools/pgdb"
	"store/uimport"
)

var (
	version string = os.Getenv("VERSION")
	module         = "demo_store"
)

func main() {
	log := logger.NewFileLogger(module)
	log.Infoln("версия", version)

	conf, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln("ошибка при чтении конфига", err)
	}

	db := pgdb.SqlxDB(conf.PostgresURL())
	if err := db.Ping(); err != nil {
		log.Fatalln("бд недоступна", err)
	}

	sm := transaction.NewSQLSessionManager(db)

	ri := rimport.NewRepositoryImports(sm)
	bi := bimport.NewEmptyBridge()

	ui := uimport.NewUsecaseImports(log, ri, bi, sm)
	bi.InitBridge(
		ui.Usecase.Notification,
	)

	api := restapi.NewRestAPI(ui, log)
	api.RunServer(conf.ApiURL())
}
