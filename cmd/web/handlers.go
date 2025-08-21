package main

import (
	"fmt"
	"html/template"

	//"log"
	"net/http"
	//"os"
	"strconv"
)

// En caso de error envia señal de error interno
// func internalError(w http.ResponseWriter, err error) bool {
// 	if err != nil {
// 		//log.Print(err.Error())
// 		errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// 		errorLog.Print(err.Error())
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return true
// 	}

// 	return false

// }

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
		//http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, " Display a specific snippet ID %d :", id)

	//w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // http.MethodPost = "POST"
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	//fmt.Fprintf(w, " Display a specific snippet ID %d :", id)

	w.Write([]byte("Create a new snippet ...."))
}
