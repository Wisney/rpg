package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"rpg/assets"
	"rpg/infra/db"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Templates
var navigationBarHTML string
var homepageTpl *template.Template
var secondViewTpl *template.Template
var thirdViewTpl *template.Template

func main() {
	serverCfg := Config{
		Host:         "0.0.0.0:" + os.Getenv("PORT"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	if !db.Exist() {
		fmt.Println("Tables not exist! Trying create...")
		err := db.CreateSchemas()
		if err != nil {
			fmt.Println(err)
			fmt.Println("main : shutting down")
			os.Exit(1)
		}
		fmt.Println("Tables Created!")
	}

	htmlServer := Start(serverCfg)
	defer htmlServer.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

	fmt.Println("main : shutting down")
}

func init() {
	navigationBarHTML = assets.MustAssetString("templates/navigation_bar.html")

	homepageHTML := assets.MustAssetString("templates/index.html")
	homepageTpl = template.Must(template.New("homepage_view").Parse(homepageHTML))

	secondViewHTML := assets.MustAssetString("templates/second_view.html")
	secondViewTpl = template.Must(template.New("second_view").Parse(secondViewHTML))

	thirdViewFuncMap := ThirdViewFormattingFuncMap()
	thirdViewHTML := assets.MustAssetString("templates/third_view.html")
	thirdViewTpl = template.Must(template.New("third_view").Funcs(thirdViewFuncMap).Parse(thirdViewHTML))
}

// Config provides basic configuration
type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// HTMLServer represents the web service that serves up HTML
type HTMLServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

// Start launches the HTML Server
func Start(cfg Config) *HTMLServer {
	// Setup Context
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup Handlers
	router := mux.NewRouter()

	//basic routes
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/favicon.ico", faviconHandler)
	router.HandleFunc("/img", imgHandler)

	//account routes
	router.HandleFunc("/", getHomeHandler).Methods("GET")
	router.HandleFunc("/logout", getLogoutHomeHandler).Methods("GET")
	router.HandleFunc("/", postHomeHandler).Methods("POST")

	router.HandleFunc("/character", characterHandler)
	router.HandleFunc("/manual", manualHandler)

	router.HandleFunc("/createaccount", getCreateAccountHandler).Methods("GET")
	router.HandleFunc("/createaccount", postCreateAccountHandler).Methods("POST")

	router.HandleFunc("/forgotpassword", getForgotPasswordHandler).Methods("GET")
	router.HandleFunc("/forgotpassword", postForgotPasswordHandler).Methods("POST")

	router.HandleFunc("/resetpassword/{token}", getResetPasswordHandler).Methods("GET")
	router.HandleFunc("/resetpassword/{token}", postResetPasswordHandler).Methods("POST")

	// Create the HTML Server
	htmlServer := HTMLServer{
		server: &http.Server{
			Addr:           cfg.Host,
			Handler:        router,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}

	// Add to the WaitGroup for the listener goroutine
	htmlServer.wg.Add(1)

	// Start the listener
	go func() {
		fmt.Printf("\nHTMLServer : Service started : Host=%v\n", cfg.Host)
		htmlServer.server.ListenAndServe()
		htmlServer.wg.Done()
	}()

	return &htmlServer
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, "static/favicon.ico")
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	http.ServeFile(w, r, "static/img.png")
}

// Stop turns off the HTML Server
func (htmlServer *HTMLServer) Stop() error {
	// Create a context to attempt a graceful 5 second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	fmt.Printf("\nHTMLServer : Service stopping\n")

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := htmlServer.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err := htmlServer.server.Close(); err != nil {
			fmt.Printf("\nHTMLServer : Service stopping : Error=%v\n", err)
			return err
		}
	}

	// Wait for the listener to report that it is closed.
	htmlServer.wg.Wait()
	fmt.Printf("\nHTMLServer : Stopped\n")
	return nil
}

// Render a template, or server error.
func render(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		fmt.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client.
func push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}

func characterHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	navbar(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	characterSheet, err := ioutil.ReadFile("./pages/character_sheet.html")
	if err != nil {
		fmt.Print(err)
	}

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(characterSheet),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

func manualHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	navbar(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	page, err := ioutil.ReadFile("./pages/manual.html")
	if err != nil {
		fmt.Print(err)
	}

	keys, ok := r.URL.Query()["topic"]

	link := ""

	if !ok || len(keys[0]) < 1 {
		link = "https://drive.google.com/file/d/1AAVaPK-2CpcObvVBJdf_6Cp4TSxP8jr-/preview"
	} else {

		switch string(keys[0]) {
		case "magic":
			link = "https://drive.google.com/file/d/1_ut_Mu82w6nRkc0R5wTh7TTEJVdj0seb/preview"
		case "skill":
			link = "https://drive.google.com/file/d/1ap3UTGOkc3ib6oBhI7iHSKo7wIycnE7z/preview"
		case "race":
			link = "https://drive.google.com/file/d/1dabUBI91i1Zb0x1Nqi1E-qOUYf3yodhd/preview"
		case "advantage":
			link = "https://drive.google.com/file/d/12Qk9P4dv88tnCoVnwlMlcmGoIyZCaB0y/preview"
		default:
			link = "https://drive.google.com/file/d/1AAVaPK-2CpcObvVBJdf_6Cp4TSxP8jr-/preview"
		}
	}

	manual := strings.Replace(string(page), "{{.Link}}", link, 3)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(manual),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}
