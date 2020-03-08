package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// 메인
func indexHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles(
		"assets/template/header.html",
		"assets/template/base.html",
		"assets/template/footer.html",
	))
	err := index.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// 파일에 따른 타입 값
func staticType(path string) string {
	var contentType string
	ext := filepath.Ext(path)
	switch ext {
	case ".html":
		contentType = "text/html"
	case ".css":
		contentType = "text/css"
	case ".js":
		contentType = "application/javascript"
	case ".jpg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	default:
		contentType = "text/plain"
	}

	return contentType
}

func webPage(port string) {
	// assets 폴더
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("assets/js"))))
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("assets/template"))))
	http.Handle("/bootstrap/", http.StripPrefix("/bootstrap/", http.FileServer(http.Dir("assets/bootstrap_3.3.2"))))
	// vendor 폴더
	http.Handle("/vendor/", http.StripPrefix("/vendor/", http.FileServer(http.Dir("assets/vendor"))))

	// main
	http.HandleFunc("/", indexHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
