package main

import (
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
)

type Music struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Artist  string  `json:"artist"`
}

func main() {
    file, _ := os.Open("music.csv")
    reader := csv.NewReader(bufio.NewReader(file))

    var music []Music

    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }

        music = append(music, Music{
            ID:         line[0],
            Title:      line[1],
            Artist:      line[2],
        })
    }

    result, _ := json.Marshal(music)
    fmt.Println(string(result))
}