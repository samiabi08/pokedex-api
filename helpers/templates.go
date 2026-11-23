package helpers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func Rendu(w http.ResponseWriter, etat int, nom string, données any) {
	var buffer bytes.Buffer
	if err := templates.ExecuteTemplate(&buffer, nom, données); err != nil {
		log.Println("rendu template:", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(etat)
	w.Write(buffer.Bytes())
}

type ErrorPage struct {
	Code    int
	Message string
}

func RenduErreur(w http.ResponseWriter, code int, message string) {
	page := ErrorPage{Code: code, Message: message}
	Rendu(w, code, "error", page)
}
