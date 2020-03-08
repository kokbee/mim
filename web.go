package main

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	http.ServeFile(w, r, "assets/template/base.html")
}

func webPage(port string) {
	// '/assets/' 경로에 접근시 파일서버 동작, 경로삭제후 'assets'으로 만 표시
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// main
	http.HandleFunc("/", Index)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
