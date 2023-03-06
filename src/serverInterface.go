package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
)

type Server struct {
	serverAddr string
	dir        string
	protocol   string
}

func getLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		fmt.Println("Error", err)
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().String()

	return localAddr
}

func getPublicIp() string {
	command := "curl"
	argument := "ifconfig.me"

	cmd := exec.Command(command, argument)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	return string(output)
}

// Given the byte array, loop over and if byte[i] is not 0 that means there is a letter at that location,
// Combining all letters gives us the name of the file sent by the user, that is later used to save transfered file with the same name
func parseBytesToString(bytes []byte) string {
	var newByteArr []byte

	//len(bytes) could be slow, will look into that later
	// in C len() is traversing the whole string witch is slow af
	for i := 0; i < len(bytes); i++ {
		//Byte array contains ASCII chars, 0 indicates that there is nothing
		if bytes[i] != 0 {
			newByteArr = append(newByteArr, bytes[i])
		}
	}
	return string(newByteArr)
}

// Read the connection and save sent file in the location
// Rewrite later to add whole directories
func recieveFile(conn net.Conn, filePath string) {
	fmt.Println("Client" + conn.RemoteAddr().String() + " connected")

	defer conn.Close()

	fileBytes := make([]byte, 100)
	//This is where the file name is located
	conn.Read([]byte(fileBytes))

	fileName := parseBytesToString(fileBytes)
	fmt.Println("Client: " + conn.RemoteAddr().String() + " sending => " + fileName)

	fo, err := os.Create(filePath + fileName)
	if err != nil {
		fmt.Println("Error:", err)
	}

	defer fo.Close()

	//Copy the file from connection with the client to the targeted directory
	_, err = io.Copy(fo, conn)
	if err != nil {
		fmt.Println("Error", err)
	}
}

// Allow download of the file
// When user connects to the server, present the available files that can be downloaded
func downloadFile(file string) {
	return
}

// Main server function
// Keeping the server alive and listening for upcoming file transfers
// For now only over tcp, udp support is to be added
func runServer(server Server) {
	_server, err := net.Listen(server.protocol, server.serverAddr)

	if err != nil {
		fmt.Println("Error", err)
		fmt.Println("Server already in use")
		return
	}

	fmt.Println("Server " + _server.Addr().String() + " running")

	defer _server.Close()

	for {
		conn, err := _server.Accept()
		if err != nil {
			fmt.Println("Error:", err)
		}

		go recieveFile(conn, server.dir)
	}
}
