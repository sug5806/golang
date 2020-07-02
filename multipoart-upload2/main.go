package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	parseErr := r.ParseMultipartForm(32 << 20)
	if parseErr != nil {
		http.Error(w, "failed to parse multipart message", http.StatusBadRequest)
		return
	}

	// Key 한개 여러개 파일 업로드
	log.Println("Value: ", r.MultipartForm.Value)
	log.Println("File: ", r.MultipartForm.File)
	log.Println("bb: ", r.MultipartForm.Value["bb"])
	for _, h := range r.MultipartForm.File["aa"] {
		file, err := h.Open()
		if err != nil {
			log.Println("err : ", err)
			http.Error(w, "sdamfkjfas", http.StatusInternalServerError)
			return
		}
		log.Println("fileName : ", h.Filename)
		tempFile, err := ioutil.TempFile("./", h.Filename[:len(h.Filename)-4]+"-*.png")
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		tempFile.Write(fileBytes)
	}

	//contentType, params, error := mime.ParseMediaType(r.Header.Get("Content-Type"))
	//if error != nil || !strings.HasPrefix(contentType, "multipart/") {
	//	http.Error(w, "expecting a multipart message", http.StatusBadRequest)
	//	return
	//}
	//
	//log.Println("contentType : ", contentType)
	//log.Println("params: ", params)
	//
	//multipartReader := multipart.NewReader(r.Body, params["boundary"])
	//defer r.Body.Close()
	//
	//for {
	//	part, err := multipartReader.NextPart()
	//	if err == io.EOF {
	//		log.Println("EOF")
	//		break
	//	}
	//
	//	if err != nil {
	//		http.Error(w, "unexpected error when retrieving a part of the message", http.StatusInternalServerError)
	//		log.Println("err : ", err)
	//		return
	//	}
	//	defer part.Close()
	//
	//	fileBytes, err := ioutil.ReadAll(part)
	//	if err != nil {
	//		http.Error(w, "failed to read content of the part", http.StatusInternalServerError)
	//		return
	//	}
	//
	//	switch part.Header.Get("Content-ID") {
	//	case "metadata":
	//		log.Println(string(fileBytes))
	//	case "media":
	//		log.Printf("filesize : %d\n", len(fileBytes))
	//		f, _ := os.Create(part.Header.Get("Content-Filename"))
	//		f.Write(fileBytes)
	//		f.Close()
	//	}
	//}

	log.Println("file upload success")

}

func main() {
	log.Print("gogogogogoo")
	http.HandleFunc("/upload", uploadFile)
	_ = http.ListenAndServe(":8000", nil)
}
