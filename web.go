package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// LoadTemplates 함수는 템플릿을 로딩합니다.
func LoadTemplates() (*template.Template, error) {
	t := template.New("")
	t, err := vfstemplate.ParseGlob(assets, t, "/template/*.html")
	return t, err
}

// 메인
func indexHandler(w http.ResponseWriter, r *http.Request) {
	// 템플릿으로 렌더링한다.
	w.Header().Set("Content-Type", "text/html")
	err = TEMPLATES.ExecuteTemplate(w, "base", rcp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
