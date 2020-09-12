package main

import (
    "os/exec"
    "runtime"
    "encoding/json"
    "log"
    "fmt"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
        log.Fatal("Failed to login: Please open the following URL in your browser. " + url)
	}
}

func dumpToJson(value interface{}) (jsonData []byte) {
    jsonData, err := json.Marshal(value)
    if err != nil {
        log.Println(err)
    }
    return
}


