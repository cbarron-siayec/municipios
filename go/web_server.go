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
type IncidenteSeguridadCatalogos struct {
	SubtiposIncidentesPoliciacosID []int    `json:"SubtiposIncidentesPoliciacosID"`
	SubtiposIncidentesPoliciacos   []string `json:"SubtiposIncidentesPoliciacos"`
	GeneroID                       []int    `json:"GeneroID"`
	Genero                         []string `json:"Genero"`
	EntidadFederativaID            []int    `json:"EntidadFederativaID"`
	EntidadFederativa              []string `json:"EntidadFederativa"`
	ColoniaID                      []int    `json:"EntidadFederativaID"`
	Colonia                        []string `json:"EntidadFederativa"`
	OcupacionID                    []int    `json:"OcupacionID"`
	Ocupacion                      []string `json:"Ocupacion"`
	EscolaridadID                  []int    `json:"Escolaridad"`
	Escolaridad                    []string `json:"Escolaridad"`
	PaisID                         []int    `json:"Pais"`
	Pais                           []string `json:"Pais"`
	CordinacionTerritorialID       []int    `json:"CordinacionTerritorialID"`
	CordinacionTerritorial         []string `json:"CordinacionTerritorial"`
}

func newIncidenteSeguridad(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		connString := "root:@(127.0.0.1:3306)/municipios"
		db, err := sql.Open("mysql", connString)
		defer db.Close()
		if err != nil {
			panic(err.Error())
		}
		var incidentePoliciacoCatalogos IncidenteSeguridadCatalogos
		//SubtiposIncidentesPoliciacos
		queryS, err := db.Query("SELECT id_subtipos_incidentes_policiacos, subtipos_incidentes_policiacos FROM subtipos_incidentes_policiacos")
		if err != nil {
			panic(err.Error())
		}
		var idSubtiposIncidentesPoliciacos int
		var subtiposIncidentesPoliciacos string
		for queryS.Next() {
			// for each row, scan the result into our tag composite object
			err := queryS.Scan(&idSubtiposIncidentesPoliciacos, &subtiposIncidentesPoliciacos)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			incidentePoliciacoCatalogos.SubtiposIncidentesPoliciacosID = append(incidentePoliciacoCatalogos.SubtiposIncidentesPoliciacosID, idSubtiposIncidentesPoliciacos)
			incidentePoliciacoCatalogos.SubtiposIncidentesPoliciacos = append(incidentePoliciacoCatalogos.SubtiposIncidentesPoliciacos, subtiposIncidentesPoliciacos)
		}
		//TERMINA SubtiposIncidentesPoliciacos
		//Genero
		queryG, err := db.Query("SELECT id_genero, genero FROM genero")
		if err != nil {
			panic(err.Error())
		}
		var idGenero int
		var genero string
		for queryG.Next() {
			// for each row, scan the result into our tag composite object
			err := queryG.Scan(&idGenero, &genero)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			incidentePoliciacoCatalogos.GeneroID = append(incidentePoliciacoCatalogos.GeneroID, idGenero)
			incidentePoliciacoCatalogos.Genero = append(incidentePoliciacoCatalogos.Genero, genero)
		}
		//TERMINA Genero
		//EntidadFederativa
		queryE, err := db.Query("SELECT id_entidad_federativa, entidad_federativa FROM entidad_federativa")
		if err != nil {
			panic(err.Error())
		}
		var idEntidadFederativa int
		var entidadFederativa string
		for queryE.Next() {
			// for each row, scan the result into our tag composite object
			err := queryE.Scan(&idEntidadFederativa, &entidadFederativa)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			incidentePoliciacoCatalogos.EntidadFederativaID = append(incidentePoliciacoCatalogos.EntidadFederativaID, idEntidadFederativa)
			incidentePoliciacoCatalogos.EntidadFederativa = append(incidentePoliciacoCatalogos.EntidadFederativa, entidadFederativa)
		}
		//TERMINA EntidadFederativa
		//Colonia
		queryC, err := db.Query("SELECT id_colonia, colonia FROM colonia")
		if err != nil {
			panic(err.Error())
		}
		var idColonia int
		var colonia string
		for queryC.Next() {
			// for each row, scan the result into our tag composite object
			err := queryC.Scan(&idColonia, &colonia)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			incidentePoliciacoCatalogos.ColoniaID = append(incidentePoliciacoCatalogos.ColoniaID, idColonia)
			incidentePoliciacoCatalogos.Colonia = append(incidentePoliciacoCatalogos.Colonia, colonia)
		}
		//TERMINA Colonia
		//Escolaridad
		queryEs, err := db.Query("SELECT id_escolaridad, escolaridad FROM escolaridad")
		if err != nil {
			panic(err.Error())
		}
		var idEscolaridad int
		var escolaridad string
		for queryEs.Next() {
			// for each row, scan the result into our tag composite object
			err := queryEs.Scan(&idEscolaridad, &escolaridad)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			incidentePoliciacoCatalogos.EscolaridadID = append(incidentePoliciacoCatalogos.EscolaridadID, idEscolaridad)
			incidentePoliciacoCatalogos.Escolaridad = append(incidentePoliciacoCatalogos.Escolaridad, escolaridad)
		}
		//TERMINA Escolaridad
		//Ocupacion
		queryO, err := db.Query("SELECT id_ocupacion, ocupacion FROM ocupacion")
		if err != nil {
			panic(err.Error())
		}
		var idOcupacion int
		var ocupacion string
		for queryO.Next() {
			// for each row, scan the result into our tag composite object
			err := queryO.Scan(&idOcupacion, &ocupacion)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			incidentePoliciacoCatalogos.OcupacionID = append(incidentePoliciacoCatalogos.OcupacionID, idOcupacion)
			incidentePoliciacoCatalogos.Ocupacion = append(incidentePoliciacoCatalogos.Ocupacion, ocupacion)
		}
		//TERMINA Ocupacion
		//Pais
		queryP, err := db.Query("SELECT id_pais, pais FROM pais")
		if err != nil {
			panic(err.Error())
		}
		var idPais int
		var pais string
		for queryP.Next() {
			// for each row, scan the result into our tag composite object
			err := queryP.Scan(&idPais, &pais)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			incidentePoliciacoCatalogos.PaisID = append(incidentePoliciacoCatalogos.PaisID, idPais)
			incidentePoliciacoCatalogos.Pais = append(incidentePoliciacoCatalogos.Pais, pais)
		}
		//TERMINA Pais
		t, _ := template.ParseFiles("../static/html/user/registro-incidente-seguridad.html")
		t.Execute(w, incidentePoliciacoCatalogos)
	} else {
		r.ParseForm()

	}
}
func userPanel(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../static/html/user/user_panel.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, t)
}

func responsibleAssigned(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../static/html/admin/usuario-asignado.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, t)
}
func asignResponsible(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../static/html/admin/asignar_responsable.html")
	if err != nil {
		panic(err.Error())
	}
	t.Execute(w, t)
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
				userPanel(w, nil)
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
	http.HandleFunc("/asignar-responsable", asignResponsible)
	http.HandleFunc("/usuario-asignado", responsibleAssigned)
	http.HandleFunc("/panel-usuario", userPanel)
	http.HandleFunc("/registro-incidente-seguridad", newIncidenteSeguridad)
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)
	log.Println("Listening...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Listen and Serve", err)
	}
}
