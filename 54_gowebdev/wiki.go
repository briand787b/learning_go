package main

import (
	"net/http"
	"io/ioutil"
	"html/template"
	"fmt"
	"errors"
	"path/filepath"
	"os"
	"strings"
	"github.com/satori/go.uuid"
)

const (
	articleDirectory  = "articles"
	sessionCookieName = "sessionId"
)

type page struct {
	Title 	string
	Body 	[]byte
}

type user struct {
	Username string
	Password string
}

// The templates for the program
var tpl = template.Must(template.ParseFiles("edit.html", "view.html", "index.html", "login.html"))

// Holds all the titles of the articles, may want to make this a slice
var articles = make(map[string]bool)

// Maps usernames to a particular user.  Must be a 1:1 relationship
// [username]user
var usernameUsers = make(map[string]*user)

// Maps session Ids to usernames.  One username may have many session Ids
// [sessionId]username
var sessionUsernames = make(map[string]string)

func init() {
	// go through existing articles and add them to the articles map
	filepath.Walk("articles", func(p string, i os.FileInfo, e error) error {
		if i.Name() != articleDirectory {
			articles[strings.Trim(i.Name(), ".txt")] = true
		}
		return e
	})
}

// May want to modify this method to return User and Bool
func validateUser(r *http.Request) bool {
	c, err := r.Cookie(sessionCookieName)
	if err != nil {
		return false
	}

	if _, ok := sessionUsernames[c.Value]; !ok {
		return false
	}
	return true
}

// May want to modify this method to return User and Bool
func creatUser(usr *user) bool {
	if _, ok := usernameUsers[usr.Username]; ok {
		// The username is already taken
		return false
	}

	usernameUsers[usr.Username] = usr
	return true
}

func createSession(un string) (sID string) {
	sID = uuid.NewV4().String()
	sessionUsernames[sID] = un
	return
}

func loadPage(title string) (*page, error) {
	// check for article existence in articles map first (no prefix)
	exists := articles[title]
	if !exists {
		return nil, errors.New("Article does not exist")
	}

	// TODO: Use filepath.join instead of manual concatenation
	relPath := articleDirectory + "/" + title + ".txt"
	body, err := ioutil.ReadFile(relPath)
	if err != nil {
		return nil, err
	}
	return &page{Title: title, Body: body}, nil
}

func savePage(p *page) error {
	rp := articleDirectory + "/" + p.Title + ".txt"
	err := ioutil.WriteFile(rp, p.Body, 0600)
	if err != nil {
		fmt.Println(err)
		return err
	}
	articles[p.Title] = true
	return nil
}

func renderTemplate(w http.ResponseWriter, title string, i interface{}) {
	if err := tpl.ExecuteTemplate(w, title+".html", i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "login", nil)
		return
	case http.MethodPost:
		usr := &user{
			Username: r.FormValue("username"),
			Password: r.FormValue("password")}
		if !creatUser(usr) {
			renderTemplate(w, "login", "Username already taken")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name: sessionCookieName,
			Value: createSession(usr.Username),
			Path: "/",
		})
		// http.Redirect(w, r, "/index/", http.StatusSeeOther)
		return
	default:
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	if !validateUser(r) {
		http.Redirect(w, r, "/login/", http.StatusForbidden)
		return
	}
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &page{Title: title}
	}
	fmt.Println(p.Title)
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	err := savePage(&page{Title: title, Body: []byte(body)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	for k := range articles {
		fmt.Println(k)
	}
	renderTemplate(w, "index", articles)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/index/", indexHandler)
	http.HandleFunc("/login/", loginHandler)

	http.ListenAndServe(":8080", nil)
}
