package main

import (
    "os"
    "path/filepath"
//    "strings"
    "time"
)

const maxFileSize = 1 * 1024 * 1024 * 1024 // 1 GB

func main() {
    // Get the current directory
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    for {
        // Walk the directory tree
        err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
            // Check if the file size is greater than 1 GB
            if info.Size() > maxFileSize {
                // Delete the file
                err = os.Remove(path)
                if err != nil {
                    return err
                }
            }

            return nil
        })
        if err != nil {
            panic(err)
        }

        // Sleep for 1 second before checking again
        time.Sleep(time.Second)
    }
}
