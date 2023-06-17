package main

import (
	"os"
	"store/bimport"
	"store/external/restapi"
	"store/internal/rimport"
	"store/tools/logger"
	"store/uimport"
)

var (
	version string = os.Getenv("VERSION")
	module         = "demo_store"
)

func main() {
	log := logger.NewFileLogger(module)
	log.Infoln("версия", version)

	ri := rimport.NewRepositoryImports()
	bi := bimport.NewEmptyBridge()

	ui := uimport.NewUsecaseImports(log, ri, bi)
	bi.InitBridge(
		ui.Usecase.Notification,
	)

	api := restapi.NewRestAPI(ui, log)
	api.RunServer()
}
