package router

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"io"
	"kite/models/db"
	"kite/utilities"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

func UpLoadImage(c *gin.Context) {
	image, err := c.FormFile("image")
	hashobj := db.Images{}
	if err != nil {
		c.JSON(http.StatusPaymentRequired, gin.H{
			"Error": "No params",
		})
		return
	}
	fileDot := strings.Index(image.Filename, ".")

	fileType := image.Filename[fileDot:]

	imageobj, err := image.Open()
	if err != nil {
		sentry.CaptureException(err)
		log.Printf("File Open Error: %v", err)
		return
	}

	defer func(imageobj multipart.File) {
		if err := imageobj.Close(); err != nil {
			sentry.CaptureException(err)
			log.Printf("File Close Error: %v", err)
			return
		}

	}(imageobj)

	nh := sha256.New()

	if _, err := io.Copy(nh, imageobj); err != nil {
		c.JSON(500, gin.H{
			"Error": "Internal Error",
		})
		sentry.CaptureException(err)
		log.Printf("Handle Error: %v", err)
		return
	}
	hash := nh.Sum(nil)
	hashStr := hex.EncodeToString(hash)

	utilities.GetDB().Model(&db.Images{}).Where(&db.Images{Hash: hashStr}).First(&hashobj)

	if hashobj.Hash != "" {
		if exist := utilities.IfileExist("./data/Images/" + hashobj.FileName); !exist {
			utilities.GetDB().Delete(&hashobj)
		} else {
			c.JSON(200, gin.H{
				"Url": "https://kite.cloudyi.xyz/images/" + hashobj.FileName,
			})
			return
		}
	}
	filename := ""

	for {
		rdk := utilities.RandomString(20)
		filename = filepath.Join(rdk + fileType)
		if !utilities.IfileExist(filename) {
			break
		}
	}
	err = c.SaveUploadedFile(image, "./data/Images/"+filename)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Internal Error",
		})
		sentry.CaptureException(err)
		log.Printf("File Save Error: %v", err)
		return
	}

	utilities.GetDB().Create(&db.Images{
		Hash:     hashStr,
		FileName: filename,
	})

	c.JSON(200, gin.H{
		"Url": "https://kite.cloudyi.xyz/images/" + filename,
	})
	return
}
