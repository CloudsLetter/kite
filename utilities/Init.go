package utilities

import (
	"github.com/getsentry/sentry-go"
	"log"
	"os"
)

func Init() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("sentryUrl"),
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	if exist := IfileExist("./data/Images"); !exist {
		err := os.MkdirAll("./data/Images", os.ModePerm)
		if err != nil {
			sentry.CaptureException(err)
			log.Printf("CreateFolder Error: %v", err)
			return
		}
	}

	if exist := IfileExist("./data/TmpImages/"); !exist {
		err := os.MkdirAll("./data/TmpImages/", os.ModePerm)
		if err != nil {
			sentry.CaptureException(err)
			log.Printf("CreateFolder Error: %v", err)
			return
		}
	}
}
