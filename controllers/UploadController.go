package controllers

import (
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io/ioutil"
	"log"
	"net/http"
	"uploadDownloadFiles/dataBase"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20) // 32MB is the default used by FormFile
	file, header, err := r.FormFile("myFile")
	if err != nil {
		log.Println("Cant open file! ", err)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	bucket, err := gridfs.NewBucket(
		dataBase.DbConnect(),
	)
	if err != nil {
		log.Println(err)
		return
	}
	uploadStream, err := bucket.OpenUploadStream(
		header.Filename,
	)
	if err != nil {
		log.Println(err)
		return
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(fileBytes)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
}
