package router

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"io"
	"kite/models/db"
	"kite/models/net"
	"kite/utilities"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func UploadUrImg(c *gin.Context) {
	uriobj := net.Url{}
	hashobj := db.Images{}
	err := c.ShouldBindJSON(&uriobj)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "No params",
		})
		return
	}
	resp, err := http.Get(uriobj.Url)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Internal Error",
		})
		sentry.CaptureException(err)
		log.Printf("Url File Get Error: %v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Internal Error",
			})
			sentry.CaptureException(err)
			log.Printf("Url File Get Error: %v", err)
			return
		}
	}(resp.Body)
	filename := ""
	for {
		rdk := utilities.RandomString(20)
		filename = filepath.Join(rdk + ".png")
		if !utilities.IfileExist(filename) {
			break
		}
	}
	if success := utilities.SaveImage2TmpImages(filename, resp); !success {
		c.JSON(500, gin.H{
			"Error": "Internal Error",
		})
		return
	}
	data, err := os.ReadFile("./data/TmpImages/" + filename)
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])
	utilities.GetDB().Model(&db.Images{}).Where(&db.Images{Hash: hashStr}).First(&hashobj)
	if hashobj.Hash != "" {
		if exist := utilities.IfileExist("./data/Images/" + hashobj.FileName); !exist {
			utilities.GetDB().Delete(&hashobj)
		} else {
			err := os.Remove("./data/TmpImages/" + filename)
			if err != nil {
				c.JSON(500, gin.H{
					"Error": "Internal Error",
				})
				sentry.CaptureException(err)
				log.Printf("Url File Remove Error: %v", err)
				return
			}
			c.JSON(200, gin.H{
				"Url": "https://kite.cloudyi.xyz/images/" + hashobj.FileName,
			})
			return
		}
	}
	err = os.Rename("./data/TmpImages/"+filename, "./data/Images/"+filename)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Internal Error",
		})
		sentry.CaptureException(err)
		log.Printf("Url File Move Error: %v", err)
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
