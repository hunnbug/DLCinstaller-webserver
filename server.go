package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func getFiles(c *gin.Context) {

	type File struct {
		FileName 	string `json:"Filename"`
		Bytes		[]byte `json:"Bytes"`
	}
	
	var files []File

	currentDir, err := os.Getwd()
	check(err)
	allFilePaths, err := os.ReadDir(currentDir + "/files")
	check(err)

	for _, item := range allFilePaths {

		filePath := currentDir + "/files/" + item.Name()

		newBytes, err := os.ReadFile(filePath)
		check(err)

		var newFile File
		newFile.FileName = item.Name()
		newFile.Bytes = newBytes

		fmt.Println(newFile.FileName)

		files = append(files, newFile)
	}

	c.IndentedJSON(http.StatusOK, files)

}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.GET("/files", getFiles)
	router.Run("localhost:8080")
}
