package routes

import (
	"appWeb/controlers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controlers.Index)
	http.HandleFunc("/new", controlers.New)
	http.HandleFunc("/insert", controlers.Insert)
	http.HandleFunc("/delete", controlers.Delete)
}
