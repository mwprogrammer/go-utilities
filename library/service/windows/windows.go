package windows

import (
	"log"

	"github.com/mwprogrammer/go-utilities/models"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
)

var service models.WindowsService

func SetFunction(function func()) {
	service.Function = function
}

func SetServiceName(name string) {
	service.Name = name
}

func SetInterval(seconds int64) {
	service.Interval = seconds
}

func UseDebugMode() {
	service.IsInDebugMode = true
}

func Run() {

	service.Function()
	
	if service.IsInDebugMode {

		err := debug.Run(service.Name, &service)

		if err != nil {
			log.Fatalln("Error running service in Debug Mode.")
		}
	}

	err := svc.Run(service.Name, &service)

	if err != nil {
		log.Fatalln("Error running service in Service Control Mode.")
	}

}