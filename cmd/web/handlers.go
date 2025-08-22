package main

import (
	"errors"
	"fmt"
	"html/template"

	//"log"
	"net/http"
	//"os"
	"strconv"

	"github.com/vmga09/go_backend/internal/models"
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

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}

	fmt.Fprintf(w, "%+v", snippet)

	//w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost { // http.MethodPost = "POST"
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n–Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

	//fmt.Fprintf(w, " Display a specific snippet ID %d :", id)

	w.Write([]byte("Create a new snippet ...."))
}
