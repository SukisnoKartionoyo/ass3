package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type status struct {
	Status wether `json:"status"`
}

type wether struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	address := "localhost:9090"
	http.HandleFunc("/", mainPage)
	log.Printf("Your service is up and running at : " + address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("Error running service: ", err)
	}

}

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		fmt.Print(err)
	}

	var obj status

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	writeToFile()

	fmt.Fprintln(w, "The Weather is :")

	fmt.Fprintln(w, "wind :", obj.Status.Wind, "kmph")
	fmt.Fprintln(w, "water :", obj.Status.Water, "m")

	if obj.Status.Wind <= 6 {
		fmt.Fprintln(w, "Wind :aman")
	}
	if obj.Status.Wind >= 7 && obj.Status.Wind <= 15 {
		fmt.Fprintln(w, "Wind :status siaga")
	}
	if obj.Status.Wind > 15 {
		fmt.Fprintln(w, "Wind :bahaya")
	}

	if obj.Status.Water <= 5 {
		fmt.Fprintln(w, "Water :aman")
	}
	if obj.Status.Water >= 6 && obj.Status.Water <= 8 {
		fmt.Fprintln(w, "Water :status siaga")
	}
	if obj.Status.Water > 8 {
		fmt.Fprintln(w, "Water :bahaya")
	}
}

func writeToFile() {
	wind := rand.Intn(100)
	water := rand.Intn(100)
	dataStatus := status{
		Status: wether{Wind: wind,
			Water: water},
	}

	fmt.Println(dataStatus)
	/*
		rankingsJson, _ := json.Marshal(rankings)
		err = ioutil.WriteFile("output.json", rankingsJson, 0644)
		fmt.Printf("%+v", rankings)
	*/
	file, _ := json.MarshalIndent(dataStatus, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)
}

func randNumber() (angka int) {
	min := 1
	max := 15
	return rand.Intn((max - min) + min)
}
