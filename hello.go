package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krenaut1/goconfig"
	"github.com/krenaut1/oauthhelper"
)

// Config this structure defines the application properties
type Config struct {
	ServerAddr       string
	ServerPort       int
	ClientID         string
	ClientSecret     string
	TokenEndPoint    string
	CertEndPoint     string
	UserInfoEndPoint string
}

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

var oauth *oauthhelper.Oauthhelper
var config = Config{}

func main() {

	// load app properties from config file based on value of PROFILE env variable
	loadAppProperites()

	// initialize the Oauth Helper
	initOauthHelper()

	// map all supported request
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink) // don't do this, always specify method
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/event/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/event/{id}", updateEvent).Methods("PUT")
	router.HandleFunc("/event/{id}", deleteEvent).Methods("DELETE")
	router.Use(authenticateUser) // add middleware function to every handler to ensure caller is authenticated with Ping
	listenAddrPort := fmt.Sprintf("%v:%v", config.ServerAddr, config.ServerPort)
	fmt.Printf("listening on: %v:%v", config.ServerAddr, config.ServerPort)
	// start listening for requests
	log.Fatal(http.ListenAndServe(listenAddrPort, router))
}

// this is a middleware function that authenticates using Ping authorization header and bearer token
func authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHdr := r.Header.Get("Authorization")
		if len(authHdr) == 0 {
			log.Println("Forbidden,no authorization header found!")
			http.Error(w, "Forbidden, no authorization header found!", http.StatusForbidden)
			return
		}
		if oauth.IsValid(authHdr) {
			next.ServeHTTP(w, r)
		} else {
			log.Println("Forbidden, authentication failed!")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	})
}

func loadAppProperites() {
	err := goconfig.GoConfig(&config) // populate my configstructure from ./config directory using PROFILE environment variable
	if err != nil {
		log.Fatalf("Error loading application properties: %v", err.Error())
	}
}

func initOauthHelper() {
	// initialize my oauth helper object
	oauth = &oauthhelper.Oauthhelper{
		MyClientID:         config.ClientID,                    // client id
		MyClientSecret:     config.ClientSecret,                // client secret
		MyTokenEndPoint:    config.TokenEndPoint,               // token end point
		MyCertEndPoint:     config.CertEndPoint,                // cert end point
		MyUserInfoEndPoint: config.UserInfoEndPoint,            // user info end point
		MyAccessToken:      "",                                 // this must be an empty string
		MyAccessTokenExp:   0,                                  // this must be zero
		MyCerts:            make(map[string]oauthhelper.Certs), // this must be make(map[string]oauthhelper.Certs)
		MyUsers:            make(map[string]oauthhelper.Users), // this must be make(map[string]oauthhelper.Users)
	}
}
