
package main

import(
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"errors"
)

type Page struct{
	Title string
	Body []byte // a byte slice
}

func (p *Page) save() error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page,error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// allow users to view a wiki page
// it will handle urls prefixed with /view/
const lenPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")
// cache for temples
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w,r)
	if err != nil{
		return
	}
	p, err := loadPage(title)
	if err != nil{
		http.Redirect(w,r,"/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	title, err := getTitle(w,r)
	if err != nil{
		return
	}
	p, err := loadPage(title)
	if err != nil{
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("save hdlr")
	title, err := getTitle(w,r)
	if err != nil{
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	fmt.Println(p.Title)
	err = p.save()
	if err != nil{
		fmt.Println("could not save", p.Title, err)
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("could save", p.Title)
	http.Redirect(w,r,"/view/"+title, http.StatusFound)
}

func getTitle(w http.ResponseWriter, r *http.Request)(title string, err error){
	title = r.URL.Path[lenPath:]
	if !titleValidator.MatchString(title){
		http.NotFound(w,r)
		err = errors.New("Invalid Page Title")
	}
	fmt.Println("getTitle:",title)
	return
}

func main(){
	/*p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))*/
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":4000",nil)
}










