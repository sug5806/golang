package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "File Upload Endpoint Hit")

	// 10MB limit
	r.ParseMultipartForm(10 << 20)

	// 주어진 키의 첫번째 파일을 가져온다
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Szie: %+v\n", handler.Size)
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)

	// temp-images 디렉토리에 해당 패턴으로 파일을 만든다
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	// form file을 읽는다
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 임시파일에 file을 쓴다
	_, err = tempFile.Write(fileBytes)

	if err != nil {
		log.Println("err : ", err)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoute() {
	http.HandleFunc("/upload", uploadFile)
	_ = http.ListenAndServe(":8000", nil)
}

func main() {
	fmt.Println("hello World")
	setupRoute()
}
