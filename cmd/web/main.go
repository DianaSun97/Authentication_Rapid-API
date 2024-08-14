package main

import (
	"encoding/gob"
	"fmt"
	"github.com/DianaSun97/elluliin_booking/internal/config"
	"github.com/DianaSun97/elluliin_booking/internal/driver"
	"github.com/DianaSun97/elluliin_booking/internal/handlers"
	"github.com/DianaSun97/elluliin_booking/internal/helpers"
	"github.com/DianaSun97/elluliin_booking/internal/models"
	"github.com/DianaSun97/elluliin_booking/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// Register your models in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.RoomRestriction{})

	// Set the application in development mode
	app.InProduction = false

	// Set up logging
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// Connect to the database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=tcs password=your_password sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
		return nil, err
	}

	log.Println("Connected to database!")

	// Create a template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	// Set up the repository and handlers
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
