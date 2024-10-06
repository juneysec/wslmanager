package main

import (
	"os/exec"
	"wslmanager/http"
)

func main() {
	exec.Command("cmd.exe", "/C", "start", "http://localhost:8080").Start()
	http.Serve()
}

