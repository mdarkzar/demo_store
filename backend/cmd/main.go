package main

import (
	"flag"
	"os"
	"store/bimport"
	"store/external/restapi"
	"store/external/worker"
	"store/internal/rimport"
	"store/tools/logger"
	"store/uimport"
)

var (
	version string = os.Getenv("VERSION")
	module         = "demo_store"
)

const (
	api         = "api"
	notifWorker = "notification_worker"
)

func main() {
	log := logger.NewFileLogger(module)
	log.Infoln("версия", version)

	ri := rimport.NewRepositoryImports()
	bi := bimport.NewEmptyBridge()

	ui := uimport.NewUsecaseImports(log, ri, bi)
	bi.InitBridge(
		ui.Usecase.Notification,
		ui.Usecase.Queue,
	)

	service := flag.String("service", "", "нужно выбрать сервис -service=<name>, доступные: -service=api, -service=notification_worker")
	flag.Parse()

	switch *service {
	case api:
		api := restapi.NewRestAPI(ui, log)
		api.RunServer()
	case notifWorker:
		w := worker.NewNotificationWorker(ui, log)
		w.Run()
	}

}
