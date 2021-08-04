package controllers

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io"
	"log"
	"net/http"
	"uploadDownloadFiles/dataBase"
	"uploadDownloadFiles/models"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	_, fsTokens, ctx := dataBase.InitCollections()

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

	var token models.Token

	token.Value = tokenGenerator()
	returnId, err := primitive.ObjectIDFromHex(fileId[0])
	if err != nil {
		log.Println(err)
		return
	}

	_, err = fsTokens.InsertOne(ctx, bson.D{
		{"idFile", returnId},
		{"token", token.Value},
	})
	if err != nil {
		log.Println(err)
		return
	}

	marshal, err := json.Marshal(token)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(marshal)
}

func DownloadByToken(w http.ResponseWriter, r *http.Request) {
	fsFiles, fsTokens, ctx := dataBase.InitCollections()

	token, ok := r.URL.Query()["token"]
	if !ok {
		log.Println(ok)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var tokenFromDB models.TokenFromDB

	err := fsTokens.FindOne(ctx, bson.M{"token": token[0]}).Decode(&tokenFromDB)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var results models.FilesInDB

	err = fsFiles.FindOne(ctx, bson.M{"_id": tokenFromDB.IdFile}).Decode(&results)
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

	_, err = fsTokens.DeleteOne(ctx, bson.M{"token": token[0]})
	if err != nil {
		log.Println(err)
	}

}

func tokenGenerator() string {
	token := make([]byte, 32)
	rand.Read(token)
	return fmt.Sprintf("%x", token)
}
