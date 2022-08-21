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
		err := exec.Command("git", "pull").Run()
		if err != nil {
			fmt.Println(err)
		}
	})

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "5555"
	}

	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
	if err != nil {
		fmt.Println(err)
	}
}
