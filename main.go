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
	Nombre := "Tienda"
	conexion, err := sql.Open(Driver, Usuario+":"+Pass+"@tcp(localhost:8080)/"+Nombre)
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
	plantillas.ExecuteTemplate(w, "carrito", nil)
}

func Contacto(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "contacto", nil)
}

func Pagar(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "pagar", nil)
}
