package controllers

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"uploadDownloadFiles/dataBase"
	"uploadDownloadFiles/models"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fsFiles, _, ctx := dataBase.InitCollections()

		cursor, err := fsFiles.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
			return
		}
		defer cursor.Close(ctx)
		var files []models.FilesInDB
		for cursor.Next(ctx) {
			var result models.FilesInDB
			err := cursor.Decode(&result)
			if err != nil {
				log.Println("cursor.Next() error:", err)
				continue
			}
			files = append(files, result)
		}
		marshal, err := json.Marshal(files)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(marshal)
	}
}
