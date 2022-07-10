// Understand routing
// Example 5 - Third party serveMux

package main

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func main() {

}
