package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

var (
	SID   string
	TOKEN string
	PORT  string
)

func main() {
	SID = os.Getenv("TWILLIO_SID")
	TOKEN = os.Getenv("TWILLIO_TOKEN")
	PORT = os.Getenv("OMNI_PORT")

	router := httprouter.New()
	router.GET("/", HandleMessage)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
}

func HandleMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
