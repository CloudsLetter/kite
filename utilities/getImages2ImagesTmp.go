package utilities

import (
	"github.com/getsentry/sentry-go"
	"io"
	"log"
	"net/http"
	"os"
)

func SaveImage2TmpImages(filename string, resp *http.Response) bool {
	imageobj, err := os.Create("./data/TmpImages/" + filename)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Url File Create Error: %v", err)
		return false
	}
	_, err = io.Copy(imageobj, resp.Body)
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Url File Write Error: %v", err)
		return false
	}
	err = imageobj.Close()
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("Url File Close Error: %v", err)
		return false
	}
	return true
}
