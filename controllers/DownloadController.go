package controllers

import (
	"bytes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io"
	"log"
	"net/http"
	"uploadDownloadFiles/dataBase"
	"uploadDownloadFiles/models"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	fsFiles, _, ctx := dataBase.InitCollections()

	fileId, ok := r.URL.Query()["id"]
	if !ok {
		log.Println(ok)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !ok || len(fileId[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	var results models.FilesInDB

	returnId, err := primitive.ObjectIDFromHex(fileId[0])
	if err != nil {
		log.Println(err)
		return
	}

	err = fsFiles.FindOne(ctx, bson.M{"_id": returnId}).Decode(&results)
	if err != nil {
		log.Println(err)
		return
	}

	bucket, _ := gridfs.NewBucket(
		dataBase.DbConnect(),
	)

	var buf bytes.Buffer
	dStream, err := bucket.DownloadToStreamByName(results.FileName, &buf)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("File size to download: %v\n", dStream)

	w.Header().Set("Content-Disposition", "attachment; filename="+results.FileName)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	io.Copy(w, &buf)
}
