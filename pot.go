package honey

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	delay     time.Duration
	loginTmpl *template.Template
)

const (
	startDelay = 600
	maxDelay   = 2000
)

func InitTemplates() error {
	var err error
	loginTmpl, err = template.New("").ParseFiles("html/login.html")
	if err != nil {
		return fmt.Errorf("Unable to parse template: %v", err)
	}

	return nil
}

// NewBee handles requests for wp-login.php
func NewBee(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodPost {
		return handleLogin(w, r)
	}
	return serveLoginPage(w, r)
}

func handleLogin(w http.ResponseWriter, r *http.Request) error {
	// Chill for a bit
	if delay == 0 || delay >= maxDelay {
		delay = startDelay
	}
	time.Sleep(delay * time.Millisecond)
	delay += 150

	// Note results
	log.Printf("user: \"%s\" pass: \"%s\"", r.FormValue("log"), r.FormValue("pwd"))

	return serveLoginPage(w, r)
}

func serveLoginPage(w http.ResponseWriter, r *http.Request) error {
	ld := struct {
		Username string
	}{
		Username: r.FormValue("log"),
	}

	w.Header().Set("Content-Type", "text/html")
	if err := loginTmpl.ExecuteTemplate(w, "login", ld); err != nil {
		return err
	}
	//fmt.Fprintf(w, "%s", lp)
	return nil
}

func openFile(d, n string) (*os.File, error) {
	return os.OpenFile(filepath.Join(d, n), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
}
