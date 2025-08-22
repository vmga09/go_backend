package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/vmga09/go_backend/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// Estructura para los logs
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	//Definimos una variable addr como flag si se ingresa valor , el puerto por defecto es 8080
	addr := flag.String("addr", ":8080", "HTTP network address and port")
	dsn := flag.String("dsn", "web:pass@tcp(dbsimulador:3306)/snippetbox?parseTime=true", "MySQL data source name")

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

	// Conexion desde el cli

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	//Ahora utilizamos el flag, go run ./cmd/web -addr=":4000"
	infoLog.Printf("Startin server on %s", *addr)
	//err := http.ListenAndServe(*addr, mux)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB opens a database connection using the provided DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
