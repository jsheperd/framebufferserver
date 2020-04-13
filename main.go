package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func fbset(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("fbset")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	} else {
		fmt.Fprintf(w, "%s\n", out)
	}
}

func main() {
	http.HandleFunc("/fbset", fbset)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
