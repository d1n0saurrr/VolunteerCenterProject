package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
	"text/template"
)

type User struct {
	Id int
	FirstName, LastName, Email, Password string
}

var users []User

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html"/*, "templates/header.html", "templates/footer.html"*/)

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=VolunteerCenter sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("select * from users")

	if err != nil {
		panic(err)
	}

	defer res.Close()
	users = []User{}

	for res.Next() {
		var user User
		err = res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)

		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}

	t.ExecuteTemplate(w, "index", users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=VolunteerCenter sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	res, err := db.Query(fmt.Sprintf("select * from users where id = $1"), id)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Close()
	showUser := User{}

	for res.Next() {
		var user User
		err = res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)

		if err != nil {
			fmt.Println(err)
		}

		showUser = user
	}

	userJson, err := json.Marshal(showUser)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, string(userJson))
}

func userPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t, err := template.ParseFiles("templates/user.html"/*, "templates/header.html", "templates/footer.html"*/)

	if err != nil {
		fmt.Println(err)
	}

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=VolunteerCenter sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	res, err := db.Query(fmt.Sprintf("select * from users where id = $1"), id)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Close()
	showUser := User{}

	for res.Next() {
		var user User
		err = res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)

		if err != nil {
			fmt.Println(err)
		}

		showUser = user
	}

	t.ExecuteTemplate(w, "user", showUser)
}

func changeUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=VolunteerCenter sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	res, err := db.Query(fmt.Sprintf("update users set firstname = $1, lastname = $2 where id = $3"),
		r.FormValue("fname"), r.FormValue("lname"), id)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Close()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func hundleFunc() {
	rtr := mux.NewRouter()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	rtr.HandleFunc("/", index)
	rtr.HandleFunc("/user/{id:[0-9]+}", userPage)
	rtr.HandleFunc("/getuser/{id:[0-9]+}", getUser).Methods("GET")
	rtr.HandleFunc("/user/{id:[0-9]+}/change", changeUser).Methods("POST")

	http.Handle("/", rtr)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	hundleFunc()
}