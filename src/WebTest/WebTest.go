// WebTest project WebTest.go
package WebTest

import (
	"io"
	"log"
	"net/http"
)

type Hello struct {
	Name string
}

func (hello Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hello.Name)
}

type World struct {
	Name string
}

func (world World) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, world.Name)
}

func HelloWeb() {

	mux := http.NewServeMux()
	mux.Handle("/hello", &Hello{"我是hello"})
	mux.Handle("/world", &World{"我是world"})
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}
