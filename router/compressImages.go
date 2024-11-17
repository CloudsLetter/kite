package router

//
//import (
//	"kite/models/net"
//	"github.com/gin-gonic/gin"
//	"image"
//	"net/http"
//)
//
//func CompressImages(c *gin.Context) {
//	imageObj, err := c.FormFile("Image")
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{
//			"Error": "No params",
//		})
//		return
//	}
//
//	compressImagesObj := net.CompressImages{}
//	err = c.ShouldBindJSON(&compressImagesObj)
//	if err != nil {
//		c.JSON(500, gin.H{
//			"Error": "No params",
//		})
//		return
//	}
//	imgBuf, err := imageObj.Open()
//
//	_, format, err := image.DecodeConfig(imgBuf)
//}
