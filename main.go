package main

import (
	"net/http"
	"os"
	"sort"
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

	router.GET("/image", ImageHomePage)
	router.GET("/image/list", ListUploadedImage)
	router.POST("/image/upload", UploadImage)
}

func ImageHomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})

}

func ListUploadedImage(c *gin.Context) {
	fs, err := os.ReadDir(baseDir)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "base directory is not exist",
		})
		return
	}

	// 按修改时间倒序
	sort.Slice(fs, func(i, j int) bool {
		a, _ := fs[i].Info()
		b, _ := fs[j].Info()
		return a.ModTime().After(b.ModTime())
	})

	var filenames []string
	for i := 0; i < 5 && i < len(fs); i++ {
		filenames = append(filenames, fs[i].Name())
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"filenames": filenames,
		},
	})

}

func UploadImage(c *gin.Context) {

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
