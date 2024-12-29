package stuff

import "net/http"

func Hello() string {
	return "Hello, world!"
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World!</h1>"))
}
