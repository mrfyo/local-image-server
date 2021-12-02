package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var baseDir = "D:\\reasoures\\image"

func main() {
	router := gin.Default()

	InitHandler(router)

	router.Run(":9090")
}

func InitHandler(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")
	router.MaxMultipartMemory = 8 << 20

	router.GET("/image", ImageUpdatePage)
	router.POST("/image/upload", ImageUploadHandler)
}

func ImageUpdatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})

}

func ImageUploadHandler(c *gin.Context) {

	file, _ := c.FormFile("file")

	filename := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10) + "." + GetFileType(file.Filename)

	filepath := baseDir + "\\" + filename

	c.SaveUploadedFile(file, filepath)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"filename": filename,
		},
	})
}

func GetFileType(filename string) (fileType string) {

	b := []byte(filename)

	for i := 0; i < len(b)-1; i++ {
		if b[i] == '.' {
			return string(b[i+1:])
		}
	}
	return ""
}
