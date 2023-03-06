package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
)

type Client struct {
	serverAddr string
	protocol   string
}

func connToServer(protocol string, server string) (net.Conn, error) {
	conn, err := net.Dial(protocol, server)
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("No server to connect to")
	}
	return conn, nil
}

func sendFile(conn net.Conn, filePath string) {
	defer conn.Close()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()

	//Write file name to the connection, it will be saved as a last file in the path
	_, err = conn.Write([]byte(filepath.Base(filePath)))
	if err != nil {
		return
	}

	//Copy the opened file to the connection
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
}
