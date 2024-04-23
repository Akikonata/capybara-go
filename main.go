package main  
  
import (  
    "fmt"  
    "log"  
    "net/http"  
)  
  
func handler(w http.ResponseWriter, r *http.Request) {  
    fmt.Fprint(w, "Hello World!")  
}  
  
func main() {  
    http.HandleFunc("/", handler)  
    log.Println("Server starting on port 8080...")  
    if err := http.ListenAndServe(":8080", nil); err != nil {  
        log.Fatal("ListenAndServe: ", err)  
    }  
}