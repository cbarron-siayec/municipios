package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type SystemUser struct {
	Email   []string `json:"email"`
	IsAdmin []bool   `json:"isAdmin"`
	Regdate []string `json:"regdate"`
}

func listIndicators(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../static/html/admin/lista_indicadores.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, t)
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	connString := "root:@(127.0.0.1:3306)/usuarios"
	db, err := sql.Open("mysql", connString)
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	query, err := db.Query("SELECT email, is_admin, regdate_rv FROM users")
	var systemUser SystemUser
	for query.Next() {
		// for each row, scan the result into our tag composite object
		var emailTemp string
		var isAdminTemp bool
		var regdateTemp string
		err := query.Scan(&emailTemp, &isAdminTemp, &regdateTemp)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		systemUser.Email = append(systemUser.Email, emailTemp)
		systemUser.IsAdmin = append(systemUser.IsAdmin, isAdminTemp)
		systemUser.Regdate = append(systemUser.Regdate, regdateTemp)
	}
	t, err := template.ParseFiles("../static/html/admin/lista_usuarios.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, systemUser)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method) //Request Method
	if r.Method == "GET" {
		t, err := template.ParseFiles("../static/html/admin/crear_usuario.html")
		if err != nil {
			panic(err.Error())
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		checkbox := r.Form.Get("isAdmin")
		isAdmin := false
		if checkbox == "on" {
			isAdmin := 1
			_ = isAdmin
		} else {
			isAdmin := 0
			_ = isAdmin
		}
		println(isAdmin)
		connString := "root:@(127.0.0.1:3306)/usuarios"
		db, err := sql.Open("mysql", connString)
		defer db.Close()
		if err != nil {
			panic(err.Error())
		}
		query, err := db.Query("SELECT email FROM users where email=?", r.Form.Get("email"))
		var user string
		for query.Next() {
			// for each row, scan the result into our tag composite object
			err := query.Scan(&user)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
		}
		println(user)
		if user == "" {
			stmt, err := db.Prepare("INSERT users SET email=?,password=?,is_admin=?")
			if err != nil {
				panic(err.Error())
			}
			res, err := stmt.Exec(email, password, isAdmin)
			if err != nil {
				panic(err.Error())
			}
			id, err := res.LastInsertId()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(id)
			listUsers(w, nil)
		} else {
			t, err := template.ParseFiles("../static/html/admin/user_exists.html")
			if err != nil {
				panic(err.Error())
			}
			t.Execute(w, nil)
			println("User ", r.Form.Get("email"), " already exists")
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method) //Request Method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../static/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		connString := "root:@(127.0.0.1:3306)/usuarios"
		db, err := sql.Open("mysql", connString)
		defer db.Close()
		if err != nil {
			panic(err.Error())
		}
		query, err := db.Query("SELECT password, is_admin FROM users where email=?", r.Form.Get("username"))
		if err != nil {
			panic(err.Error())
		}
		var password string
		var isAdmin bool
		for query.Next() {
			// for each row, scan the result into our tag composite object
			err := query.Scan(&password, &isAdmin)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
		}
		if password == r.Form.Get("password") {
			if isAdmin {
				listUsers(w, nil)
			} else {
				t, err := template.ParseFiles("../static/html/user/user_panel.html")
				if err != nil {
					panic(err.Error())
				}
				t.Execute(w, nil)
			}
		} else {
			fmt.Println("Warning ", r.Form.Get("username"), " - Wrong Pasword")
			t, _ := template.ParseFiles("../static/forgot-password.html")
			t.Execute(w, nil)
		}
	}
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/crear-usuario", createUser)
	http.HandleFunc("/lista-usuarios", listUsers)
	http.HandleFunc("/lista-indicadores", listIndicators)
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)
	log.Println("Listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}
