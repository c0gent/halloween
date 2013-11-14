package main

import (
	"encoding/gob"
	_ "github.com/bmizerany/pq"
	"github.com/gorilla/pat"
	"github.com/nsan1129/auctionLog/log"
	"github.com/nsan1129/unframed"
	"html/template"
	"net/http"
)

//This will be read from a text file eventually

var TS *templateStore
var DB *unframed.DB

func main() {

	DB = unframed.NewDB("postgres", "user=postgres password=postgres dbname=holloween sslmode=disable")
	prepareStatements()

	gob.Register(&Person{})
	gob.Register(&Session{})

	defer DB.Close()
	initTS()
	serve()
}

// ---------- ROUTES -----------

func initRouter() {
	a := pat.New()

	a.Get("/People/list", listPeople)
	a.Get("/Person/compose", listPeople)
	a.Get("/Person/show", showPerson)
	a.Post("/Person/update/vote", voteUpdatePerson)
	//a.Get("/Candy/Users", listUsers)
	a.Get("/Session/compose", composeSession)
	a.Get("/Session/delete", deleteSession)
	a.Get("/Session/create/fail", failCreateSession)
	a.Post("/Session/create", createSession)
	a.Post("/Person/create", listPeople)
	a.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	a.Get("/", listPeople)

	http.Handle("/", a)
}

func serve() {
	initRouter()
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

// -------- TEMPLATES ------------

type templateStore struct {
	listPeople        *template.Template
	composePerson     *template.Template
	composeSession    *template.Template
	failCreateSession *template.Template
	showPerson        *template.Template
}

func initTS() {
	TS = new(templateStore).init()
}

func (ts *templateStore) init() *templateStore {
	ts.loadTemplates()
	log.Message("Templates Loaded")
	return ts
}

func (ts *templateStore) loadTemplates() {
	ts.listPeople = template.Must(template.ParseFiles("tmpl/base.html.tmpl", "tmpl/People/listPeople.html.tmpl"))
	ts.composePerson = template.Must(template.ParseFiles("tmpl/base.html.tmpl"))
	ts.composeSession = template.Must(template.ParseFiles("tmpl/base.html.tmpl", "tmpl/People/composeSession.html.tmpl"))
	ts.failCreateSession = template.Must(template.ParseFiles("tmpl/base.html.tmpl", "tmpl/People/composeSession.html.tmpl"))
	ts.showPerson = template.Must(template.ParseFiles("tmpl/base.html.tmpl", "tmpl/People/showPerson.html.tmpl"))
}

/*
--TERMINOLOGY--
- Actions relating to Data Objects (Records). Must be Verbs. -
Go			Http		SQL				Purpose
------------------------------------------------------------
create		POST		INSERT			store new record
show		GET 		SELECT			display record
update		POST(PUT)	UPDATE			modify existing record
delete		DELETE		DELETE			destroy existing record

list		GET 		SELECT			display multiple records
compose		GET			(none)			display composition controls/tools
edit		GET			(none)			display editing controls/tools


- Other -
find
NewXXX()		Return a new instance of something. Customary GO shorthand for GetNewXXX().

-Other Terms-
List(noun) = Table of Data, rows(multiple), etc.

*/
