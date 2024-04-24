package main  
  
import (  
    "fmt"  
    "log"  
    "net/http"  
)  
  
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}
  
func main() {  
    http.HandleFunc("/", HelloWorldHandler)  
    log.Println("Server starting on port 8080...")  
    if err := http.ListenAndServe(":8080", nil); err != nil {  
        log.Fatal("ListenAndServe: ", err)  
    }  
}