package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"snippetbox.samuel/internal/models"

	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	logger         *slog.Logger
	snippets       models.SnippetModelInterface
	templateCache  map[string]*template.Template
	sessionManager *scs.SessionManager
	users          models.UserModelInterface
	debug          bool
}

func main() {
	// Parse command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	debug := flag.Bool("debug", false, "Debug mode")
	flag.Parse()

	// Create logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Create DB connection
	db, err := openDB("./lite.db")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Create template cache
	cache, err := newTemplateCache()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// Session Manager
	sessionManager := scs.New()
	sessionManager.Store = sqlite3store.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	// Create struct with dependencies
	app := &application{
		logger:         logger,
		snippets:       &models.SnippetModel{DB: db},
		templateCache:  cache,
		sessionManager: sessionManager,
		users:          &models.UserModel{DB: db},
		debug:          *debug,
	}

	logger.Info("starting server", "addr", *addr)
	// start web server

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
