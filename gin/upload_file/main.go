/**
 * @author wei.song
 * @since 2023/10/30 16:34
 */
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("------------- upload file --------------")
	// uploadSingleFile()

	fmt.Println("------------- upload file --------------")
	uploadMultipleFile()
}

func uploadSingleFile() {
	engine := gin.Default()
	engine.POST("/upload/single", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		// context.SaveUploadedFile(file, file.Filename)
		context.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	_ = engine.Run(":9999")
}

func uploadMultipleFile() {
	engine := gin.Default()
	engine.POST("/upload/multiple", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["file"]

		result := make(map[string]string)

		for _, file := range files {
			// context.SaveUploadedFile(file, file.Filename)
			result[file.Filename] = fmt.Sprintf("'%s' uploaded!", file.Filename)
		}

		context.JSON(http.StatusOK, result)
	})

	_ = engine.Run(":9999")
}
