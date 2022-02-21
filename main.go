package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
    "os"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("static/templates", "layout.html")
	fp := filepath.Join("static/templates", filepath.Clean(r.URL.Path))


    // Check if template is available, if not send 404
    
    info, err := os.Stat(fp)
    if err != nil {
        if os.IsNotExist(err){
            http.NotFound(w,r)
            return
        }
    }

    // Send 404 if the client is requesting a directory
    if info.IsDir() {
        http.NotFound(w,r)
        return
    }

	tmpl, err := template.ParseFiles(lp, fp)

    if err != nil {
        log.Println(err.Error())
        http.Error(w, http.StatusText(500), 500)
        return
    }
    
    
	err = tmpl.ExecuteTemplate(w, "layout", nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, http.StatusText(500), 500)
    }
}

