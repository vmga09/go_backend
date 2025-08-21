package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Estructura para los logs
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	//Definimos una variable addr como flag si se ingresa valor , el puerto por defecto es 8080
	addr := flag.String("addr", ":8080", "HTTP network address and port")

	//Usamos flag.Parse()
	flag.Parse()
	// GUARDAR EN UN ARCHIVO LOS LOGS /////////////////////////////
	//Se define el archivo de salida
	f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	//Si hay un error se detiene la aplicacion
	if err != nil {
		log.Fatal(err)
	}
	// EL defer hara que se ejecute al final
	defer f.Close()
	//////////////////////////////////////
	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}
	//Ahora utilizamos el flag, go run ./cmd/web -addr=":4000"
	infoLog.Printf("Startin server on %s", *addr)
	//err := http.ListenAndServe(*addr, mux)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
