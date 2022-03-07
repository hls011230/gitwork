package main

import (
	"fmt"
	"net/http"
	"os"
)
func handler(w http.ResponseWriter, r *http.Request) {
	target := "欢迎使用AllSmile管理！"
	fmt.Fprintf(w, "%s!\n", target)
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
}
