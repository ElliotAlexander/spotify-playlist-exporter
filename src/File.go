package main

import (
    "io/ioutil"
    "os"
)

func writeToFile(filename string, data []byte){
    err := ioutil.WriteFile(filename, data, 0644)
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

