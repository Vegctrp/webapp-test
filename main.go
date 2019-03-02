package main
import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	fmt.Println("test")
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./doc"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}