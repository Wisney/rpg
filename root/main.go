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
	"strconv"
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

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/second", SecondHandler)
	router.HandleFunc("/third/{number}", ThirdHandler)
	router.HandleFunc("/favicon.ico", faviconHandler)
	router.HandleFunc("/img", imgHandler)
	router.HandleFunc("/createaccount", getCreateAccountHandler).Methods("GET")
	router.HandleFunc("/createaccount", postCreateAccountHandler).Methods("POST")
	router.HandleFunc("/character", characterHandler)

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

// Route Handlers

// HomeHandler renders the homepage view template
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//get html file
	formSignin, err := ioutil.ReadFile("./pages/form_signin.html")
	if err != nil {
		fmt.Print(err)
	}

	page := strings.Replace(string(formSignin), "{{.Message}}", "E aew os caras! e.e", 1)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// SecondHandler renders the second view template
func SecondHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
	}
	render(w, r, secondViewTpl, "second_view", fullData)
}

// ThirdHandler renders the third view template
func ThirdHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	push(w, "/static/third_view.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var queryString string
	pathVariables := mux.Vars(r)
	queryNumber, err := strconv.Atoi(pathVariables["number"])
	if err != nil {
		queryString = pathVariables["number"]
	}
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Number":        queryNumber,
		"StringQuery":   queryString,
	}
	render(w, r, thirdViewTpl, "third_view", fullData)
}
