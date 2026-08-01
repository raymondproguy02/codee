// Daily GitHub contribution count
package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run csv2json.go <file.csv>")
        os.Exit(1)
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    headers, err := reader.Read()
    if err != nil {
        panic(err)
    }

    var result []map[string]string

    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            panic(err)
        }

        row := make(map[string]string)
        for i, value := range record {
            row[headers[i]] = value
        }
        result = append(result, row)
    }

    jsonData, _ := json.MarshalIndent(result, "", "  ")
    fmt.Println(string(jsonData))
}
