package main

import (
	"fmt"
	"log"
	"net/http"
  "./models"
	"github.com/gorilla/mux"
  daos "./dao"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
	"golang.org/x/net/context"
	 "golang.org/x/oauth2"
	 "golang.org/x/oauth2/google"

)

var dao = daos.TopicsDAO{}


func AllTopicsEndPoint(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("MONGO connection")
  topics, err := dao.FindAll()
	if err != nil {
      fmt.Printf("MONGO error")
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, topics)
}

func CreateTopicsEndPoint(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  var topic models.Topic
  fmt.Printf("%s \n",r.Body)
	if err := json.NewDecoder(r.Body).Decode(&topic); err != nil {
    fmt.Printf("%s could not decode",err.Error())
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	topic.ID = bson.NewObjectId()
	if err := dao.Insert(topic); err != nil {
    fmt.Printf("%s could not insert",err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, topic)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {

	dao.Server = "localhost"
	dao.Database = "alumbridge"
	dao.Connect()
}

func LoginEndPoint(w http.ResponseWriter, r *http.Request) {
	  defer r.Body.Close()
		var tokenBlob models.TokenBlob
		if err := json.NewDecoder(r.Body).Decode(&tokenBlob); err != nil {
	    fmt.Printf("%s could not decode in login\n",err.Error())
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
	ctx := context.Background()
conf := &oauth2.Config{
    ClientID:     "320180218476-272l5ct9t9c517151r19fj7qj3g2i9kf.apps.googleusercontent.com",
    ClientSecret: "9U8Ga4Yhf4J2GbertBZMRKTH",
    Scopes:       []string{"profile"},
    Endpoint: google.Endpoint,
}

// // Redirect user to consent page to ask for permission
// // for the scopes specified above.
// url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
// fmt.Printf("Visit the URL for the auth dialog: %v", url)
//
// // Use the authorization code that is pushed to the redirect
// // URL. Exchange will do the handshake to retrieve the
// // initial access token. The HTTP Client returned by
// // conf.Client will refresh the token as necessary.
fmt.Printf("%s is tokenBlob \n ",tokenBlob.AccessToken)
tok, err := conf.Exchange(oauth2.NoContext, tokenBlob.AccessToken)
if err != nil {
	fmt.Printf("%s is the error in exchange \n",err.Error())
}
fmt.Printf("about to perform get\n")

client := conf.Client(ctx, tok)
info,err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
if err != nil {
	fmt.Printf("%s is the error in exchange \n",err.Error())

}
fmt.Printf("%s is the info \n",info)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/topics", AllTopicsEndPoint).Methods("GET")
	r.HandleFunc("/createtopic", CreateTopicsEndPoint).Methods("POST")
	r.HandleFunc("/login", LoginEndPoint).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
