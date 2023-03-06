package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var _server Server

// Choosing a file to send
func chooseFile(w fyne.Window, h *widget.Label) string {
	fileName := ""

	dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
		if nil != err {
			dialog.ShowError(err, w)
			os.Exit(1)
		}
		//if this check is not done the whole app crashes
		if file == nil {
			return
		}
		fileName = file.URI().Path()
		h.SetText(fileName)
	}, w)

	return fileName
}

// Home Page
func renderHomePage() {
	//To be changed to dynamically determine serverAddr and protocol based on file type and general layout/arhitecture of the
	_server.serverAddr = "127.0.0.1:5000"
	_server.protocol = "tcp"
	_server.dir = "/home/corazone/send/"

	//Fyne init
	application := app.New()
	window := application.NewWindow("Geass")
	//================================================================>>

	//Widgets and widget icons
	//_W stands for widget
	serverImg, err := fyne.LoadResourceFromPath("/home/corazone/programming/go/server_app/resources/web-server-icon.svg")
	if err != nil {
		fmt.Println("Could not load img", err)
	}

	serverW := widget.NewButtonWithIcon("", serverImg, func() {
		go runServer(_server)
	})

	fileImg, err := fyne.LoadResourceFromPath("/home/corazone/programming/go/server_app/resources/folder-icon.svg")
	if err != nil {
		fmt.Println("Could not load img", err)
	}

	clientW := widget.NewLabel("File :")

	openfileW := widget.NewButtonWithIcon("", fileImg, func() {
		chooseFile(window, clientW)
	})

	sendFileImg, err := fyne.LoadResourceFromPath("/home/corazone/programming/go/server_app/resources/send-message-icon.svg")
	if err != nil {
		fmt.Println("Could not load img", err)
	}

	sendfileW := widget.NewButtonWithIcon("", sendFileImg, func() {
		conn, err := connToServer(_server.protocol, _server.serverAddr)
		if err != nil {
			fmt.Println("No server to connect to")
			return
		}
		//If there is no server to connect to, don't crash go to home-screen
		//Bug fix
		if conn == nil {
			return
		}

		sendFile(conn, clientW.Text)
	})

	exitImg, err := fyne.LoadResourceFromPath("/home/corazone/programming/go/server_app/resources/emergency-exit-icon.svg")
	if err != nil {
		fmt.Println("Could not load img", err)
	}

	exitW := widget.NewButtonWithIcon("", exitImg, func() {
		os.Exit(1)
	})

	//Positions

	serverW.Resize(fyne.NewSize(60, 60))
	openfileW.Resize(fyne.NewSize(60, 60))
	sendfileW.Resize(fyne.NewSize(60, 60))
	exitW.Resize(fyne.NewSize(60, 60))

	serverW.Move(fyne.NewPos(200, 400))
	openfileW.Move(fyne.NewPos(300, 400))
	sendfileW.Move(fyne.NewPos(400, 400))
	exitW.Move(fyne.NewPos(500, 400))
	//================================================================>>
	//fmt.Println("Pos of widget is ", serverW.Position().X, serverW.Position().Y)
	//Rendering widgets in the window
	window.SetContent(container.NewWithoutLayout(
		clientW,
		openfileW,
		sendfileW,
		serverW,
		exitW,
	))
	//================================================================>>

	//General window settings
	window.Resize(fyne.NewSize(800, 500))
	window.ShowAndRun()
	//================================================================>>
}

func main() {
	renderHomePage()
}

//How fyne works:
//App instance must be created, inside every app there is a window that we can use to render contents.
//The contents of every fyne.Window is a fyne.Canvas
//Inside every fyne.Canvas is at least one fyne.CanvasObject
//The fyne.Container extends the fyne.CanvasObject type to include managing multiple child objects

//func renderClient(clientWidget) {
// render
//}

//func renderServer(serverWidget) {
//render
//}
//

//AirDrop
//
//
