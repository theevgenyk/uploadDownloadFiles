package main

import (
	"log"
	"net/http"
	"uploadDownloadFiles/controllers"
)

func main() {
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/upload", controllers.UploadHandler)
	http.HandleFunc("/files", controllers.ListHandler)
	http.HandleFunc("/download", controllers.DownloadHandler)
	http.HandleFunc("/createToken", controllers.CreateToken)
	http.HandleFunc("/downloadByToken", controllers.DownloadByToken)
	err := http.ListenAndServe(":8080", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}
