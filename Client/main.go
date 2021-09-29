package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
   //getCourses();
   //getCourse();
   newCourse();
}

func getCourses(){
	resp, err := http.Get("http://localhost:8080/courses")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}

func getCourse(){
	resp, err := http.Get("http://localhost:8080/courses/9456")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}

func newCourse(){
	Data := map[string]string{"id": "0356","name": "BDSA","ects": "15","rating": "34"}
	json_data, err := json.Marshal(Data)


    if err != nil {
        log.Fatal(err)
    }

	resp, err := http.Post("http://localhost:8080/courses", "application/json", bytes.NewBuffer(json_data))
	
    if err != nil {
        log.Fatal(err)
    }

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res)
}
