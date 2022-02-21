package main 

import (
    "fmt"
    "os"
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
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil 
    }
    return &Page{Title: title, Body: body}
}

func main(){
 p1 := &Page{Title: "TestPage", Body:[]byte("This is a sample page.")}
 p1.save()
 p2 := loadPage("TestPage")
 fmt.Println(string(p2.Body))
}
