package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("exec Command `git pull`")
		exec.Command("git", "pull").Output()
	})

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "5555"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}
