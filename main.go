package main

import (
	"log"
	"net/http"
	"os"
	"html/template"
    "github.com/russross/blackfriday"
)


type Page struct {
    Title string
    Body template.HTML
}


func loadPage(title string) *Page {
    filename := title + ".md"
    body, err := os.ReadFile("./posts/" + filename)
    if err != nil {
        return nil 
    }
    //convertedToHtml := markdown.ToHTML(body, nil, nil)
    newPost := template.HTML(blackfriday.MarkdownCommon([]byte(body)))
    return &Page{Title: title, Body: newPost}
}

func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/posts/"):]
    p := loadPage(title)
    t,_ := template.ParseFiles("./templates/posts.html")
    t.Execute(w,p)
}

func main(){
    http.HandleFunc("/posts/", viewHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Fatal(http.ListenAndServe(":3000", nil))
}
