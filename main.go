package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "github.com/heroku/x/hmetrics/onload"
)

type User struct {
	Username string `json:"username"`
	Follower int    `json:"followers"`
}

func getJson() []byte {

	json_resp, err := http.Get("https://jsonkeeper.com/b/DMXK")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	json_data, err := ioutil.ReadAll(json_resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return json_data

}

func getOneDataDetail(w http.ResponseWriter, r *http.Request) {
	dataName := mux.Vars(r)["userid"]

	user := map[string]User{}
	json.Unmarshal(getJson(), &user)
	json.NewEncoder(w).Encode(user[dataName])
}
func getOneDataFollowers(w http.ResponseWriter, r *http.Request) {
	dataName := mux.Vars(r)["Username"]

	User := map[string]User{}
	json.Unmarshal(getJson(), &User)

	for _, singleData := range User {
		if singleData.Username == dataName {
			json.NewEncoder(w).Encode(singleData.Follower)
		}
	}
}

func getAllDatas(w http.ResponseWriter, r *http.Request) {
	User := map[string]User{}
	json.Unmarshal(getJson(), &User)

	json.NewEncoder(w).Encode(User)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/datas", getAllDatas).Methods("GET")
	router.HandleFunc("/datas/{userid}/detail", getOneDataDetail).Methods("GET")
	router.HandleFunc("/datas/follower/{Username}", getOneDataFollowers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
