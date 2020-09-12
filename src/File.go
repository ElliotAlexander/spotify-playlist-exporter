package main

import (
    "io/ioutil"
    "strings"
    "os"
)

func writeToFile(path string, filename string, data []byte) {
    // Playlists with forward slashes in their name will break without this.
    filepath := path + stripSlashes(filename)
    err := ioutil.WriteFile(filepath, data, 0644)
    if err != nil {
        panic(err)
    }
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }

    return !info.IsDir()
}

func stripSlashes(filename string) (filenameStripped string){
    filenameStripped = strings.Replace(filename, "/", "", -1)
    return
}
