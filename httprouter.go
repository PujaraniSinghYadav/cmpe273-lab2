package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "net/http"
    "io/ioutil"
)

type Request_Json_Obj struct{
 Name string
}

type Response_Json_Obj struct{
    Greeting string
}

var jsonObj Request_Json_Obj

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func post_hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
        body, err := ioutil.ReadAll(req.Body)
        if err != nil {
        panic(err)
        }
        if err := req.Body.Close(); err != nil {
        panic(err)

        }
        if err := json.Unmarshal(body, &jsonObj); 
        err != nil {
        rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
            if err := json.NewEncoder(rw).Encode(err); err != nil {
            panic(err)
            }
        }

        fmt.Println(jsonObj.Name)
        rw.WriteHeader(http.StatusCreated)
        Response_Json_Obj := Response_Json_Obj{"Hello,"+jsonObj.Name}
        js, err := json.Marshal(Response_Json_Obj)
        if err != nil {
        panic(err)
        }
        rw.Write(js)

}

func main() {
    mux := httprouter.New()
    mux.GET("/hello/:name", hello)
    mux.POST("/hello",post_hello)
    server := http.Server{
            Addr:        "127.0.0.1:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}
