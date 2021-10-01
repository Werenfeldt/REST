package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
   getCourses();
   fmt.Println();
   delete();
   fmt.Println();
   getCourse();
   fmt.Println();
   getCourses();
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
	resp, err := http.Get("http://localhost:8080/courses/6666")

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

func delete(){
    client := &http.Client{}
    req, err := http.NewRequest("DELETE", "http://localhost:8080/courses/6666", nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    // Read Response Body
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(respBody))
}
