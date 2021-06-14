package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var plantillas = template.Must(template.ParseGlob("Plantillas/*"))

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Pass := "root"
	Nombre := "tienda"
	conexion, err := sql.Open(Driver, Usuario+":"+Pass+"@tcp(localhost:3306)/"+Nombre)
	if err != nil {
		panic(err.Error())
	}
	return conexion
}

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/empresa", Empresa)
	http.HandleFunc("/galeria", Galeria)
	http.HandleFunc("/carrito", Carrito)
	http.HandleFunc("/contacto", Contacto)
	http.HandleFunc("/pagar", Pagar)
	http.HandleFunc("/registrar", Registrar)
	http.HandleFunc("/insertarUsuario", InsertarUsuario)
	http.HandleFunc("/insertarContacto", InsertarContacto)
	http.Handle("/Imagenes/", http.StripPrefix("/Imagenes/", http.FileServer(http.Dir("Imagenes/"))))
	log.Println("Servidor Encendido...")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Empresa(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "empresa", nil)
}

func Galeria(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "galeria", nil)
}
func Carrito(w http.ResponseWriter, r *http.Request) {
	//conexionBD()
	plantillas.ExecuteTemplate(w, "carrito", nil)
}

func Pagar(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "pagar", nil)
}

func Contacto(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "contacto", nil)
}

func Registrar(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "registrar", nil)
}

func InsertarUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("Nombre")
		apellido := r.FormValue("Apellido")
		correo := r.FormValue("Correo")
		telefono := r.FormValue("Telefono")

		conexionEst := conexionBD()
		insertarRegistro, err := conexionEst.Prepare("INSERT INTO contactos(Nombre, Apellido, Correo, Telefono) VALUES(?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		insertarRegistro.Exec(nombre, apellido, correo, telefono)

		http.Redirect(w, r, "/carrito", 301)
	}
}

func InsertarContacto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		Nombre := r.FormValue("Nombre")
		Apellido := r.FormValue("Apellido")
		Correo := r.FormValue("Correo")
		Telefono := r.FormValue("Telefono")
		Descripcion := r.FormValue("Descripcion")

		conexionEstablecida := conexionBD()

		InsertarRegistro, err := conexionEstablecida.Prepare("INSERT INTO contactos(Nombre, Apellido, Correo, Telefono, Descripcion) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		InsertarRegistro.Exec(Nombre, Correo, Apellido, Telefono, Descripcion)
		http.Redirect(w, r, "/contacto", 301)
	}
}
