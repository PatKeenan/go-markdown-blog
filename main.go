package main 

import (
    "fmt"
    "os"
    "log"
    "net/http"
    "github.com/gomarkdown/markdown"
)


type Page struct {
    Title string
    Body []byte
}

func (p *Page) save() error{
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
    filename := title + ".md"
    body, err := os.ReadFile("./posts/" + filename)
    if err != nil {
        return nil 
    }
    convertedBody := markdown.ToHTML(body, nil, nil)
    return &Page{Title: title, Body: convertedBody}
}

func generator(){
 p1 := &Page{Title: "TestPage", Body:[]byte("This is a sample page.")}
 p1.save()
 p2 := loadPage("TestPage")
 fmt.Println(string(p2.Body))
}

func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/view/"):]
    p := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler (w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[len("/edit/"):]
    p := loadPage(title)
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)

}

func main(){
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    log.Fatal(http.ListenAndServe(":3000", nil))
}
