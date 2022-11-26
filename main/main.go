package main

import (
	"bufio" //Read File by line
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify" //Notifications
)

var watcher *fsnotify.Watcher
var fileToWatch *os.File

func sendFileData(fileName string, cameraId string) {
	body := map[string]interface{}{
		"filename": fileName,
		"state":    "uploaded",
		"cameraId": cameraId,
	}
	requestURL := os.Getenv("MANAGER") + "/api/events/file"
	byts, _ := json.Marshal(body)
	jsonBody := []byte(byts)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}
	fmt.Printf("Response: %v\n", res)
}

func main() {
	//Check if received folder to watch
	folderToWatch := os.Args[1]

	//Creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	//Check if the argument is a valid folder
	if err := filepath.Walk(folderToWatch, watchDir); err != nil {
		fmt.Println("ERROR", err)
	}

	defer fileToWatch.Close()

	done := make(chan bool)

	go func() {
		fileLineScanner := bufio.NewReader(fileToWatch)
		cameraId := os.Args[2]
		for {
			select {
			//Watch for events
			case event := <-watcher.Events:
				fmt.Println("INFO: Event ", event)
				if event.Op == fsnotify.Write {
					fileName, _, _ := fileLineScanner.ReadLine()
					if string(fileName) != "" {
						sendFileData(string(fileName), cameraId)
					}
				}
			//Watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

// This gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {
	//Check if the file exists
	fmt.Println("INFO: Watching directory: ", path)
	if _, fileNotExistsError := os.Stat(path); fileNotExistsError != nil || os.IsNotExist(fileNotExistsError) {
		_, fileCreationError := os.Create(path)
		if fileCreationError != nil {
			return nil
		}
	}
	//Now that we're sure the file exists, we can open it to read it in the goroutine
	file, fileOpenError := os.Open(path)
	if fileOpenError != nil {
		return nil
	}
	//Save the file reference to the file that will be accessed within the goroutine
	fileToWatch = file
	//We add the file to the watcher
	return watcher.Add(path)
}
