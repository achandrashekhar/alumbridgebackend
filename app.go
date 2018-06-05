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
    "strconv"
)

var dao = daos.TopicsDAO{}


func AllTopicsEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateTopicsEndPoint(w http.ResponseWriter, r *http.Request) {

  r.ParseForm()
  fmt.Printf("Name is %s \n", r.Form.Get("Name"))
  defer r.Body.Close()
	var topic models.Topic
	if err := json.NewDecoder(r.Body).Decode(&topic); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	topic.ID = bson.NewObjectId()

  var yearEstablished, _ = strconv.ParseInt(r.Form.Get("year"), 10, 32)
  topic.YearEstablished = int32(yearEstablished)
	if err := dao.Insert(topic); err != nil {
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

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/topics", AllTopicsEndPoint).Methods("GET")
	r.HandleFunc("/createtopic", CreateTopicsEndPoint).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
