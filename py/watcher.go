package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/fsnotify/fsnotify"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run watcher.go <directory>")
        os.Exit(1)
    }

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                fmt.Printf("📁 %s: %s\n", event.Op, event.Name)
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("Error:", err)
            }
        }
    }()

    err = watcher.Add(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("👀 Watching: %s\n", os.Args[1])
    <-done
}
