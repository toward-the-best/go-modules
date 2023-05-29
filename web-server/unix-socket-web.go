package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func systemCall(cmd *exec.Cmd) {
	//cmd := exec.Command("ls", "-l")

	// Execute the command
	err := cmd.Run()

	// Error handling
	if err != nil {
		// Try to get the exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			// The command failed to run
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				fmt.Println("Exit Status:", status.ExitStatus())
			}
		} else {
			log.Fatal(err)
		}
	} else {
		// The command was successful
		if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
			fmt.Println("Exit Status:", status.ExitStatus())
		}
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		println("OK!!! Request")
		time.Sleep(5 * time.Second)
		cmd := exec.Command("ls", "-l")
		systemCall(cmd)
		w.Write([]byte("Hello from Unix socket server!"))
	})

	unixSocketFile := "/tmp/unix.sock"

	// Make sure the socket file does not already exist
	if err := os.Remove(unixSocketFile); err != nil && !os.IsNotExist(err) {
		log.Fatalf("failed to remove previous unix socket file %s: %v", unixSocketFile, err)
	}

	// Listen on the Unix domain socket
	listener, err := net.Listen("unix", unixSocketFile)
	if err != nil {
		log.Fatalf("failed to listen on unix socket file %s: %v", unixSocketFile, err)
	}

	// Serve HTTP over the Unix domain socket
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
