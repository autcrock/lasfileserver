package main
import (
  "encoding/json"
  "io/ioutil"
  "net/http"
)

const DataDirectory  = "data"

func ping(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("pong"))
}

func enableCors(w *http.ResponseWriter) {
  (*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func listLas(w http.ResponseWriter, r *http.Request) {
  // WARNING allows access to local file system - TESTING ONLY
  enableCors(&w)
  files, err := ioutil.ReadDir(DataDirectory)
  if err != nil {
    w.Write([]byte("LAS files not found."))
  } else {
    var list [] string
    for _, file := range files {
      list = append(list,file.Name())
    }
    var jsonOut, _ = json.Marshal(list)
    w.Write([]byte(jsonOut))
  }
}

func lasFile(w http.ResponseWriter, r *http.Request) {
  // WARNING allows access to local file system - TESTING ONLY
  enableCors(&w)
  fileName := DataDirectory + "/" + r.URL.Path[len("/las/"):]
  file, err := ioutil.ReadFile(fileName)
  if err != nil {
    w.WriteHeader(http.StatusNotFound)
  } else {
    w.Write([]byte(file))
  }
}


func main() {
  http.Handle("/", http.FileServer(http.Dir("./" + DataDirectory)))
  http.HandleFunc("/las", listLas)
  http.HandleFunc("/las/", lasFile)
  http.HandleFunc("/ping", ping)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
