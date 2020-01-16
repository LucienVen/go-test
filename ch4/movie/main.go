package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type Movie struct {
    Title string
    Year int `json:"released"`
    Color bool `json:"color,omitempty"`
    Actors []string
}
var movies = []Movie{
    {Title: "Casablanca", Year: 1942, Color: false,
        Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
    {Title: "Cool Hand Luke", Year: 1967, Color: true,
        Actors: []string{"Paul Newman"}},
    {Title: "Bullitt", Year: 1968, Color: true,
        Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
    // ...
}

// json编码
func jsonMarshal1(movie []Movie) ([]byte, error) {
    return json.Marshal(movie)
}

// json 编码（格式化）
func jsonMarshal2(movie []Movie) ([]byte, error) {
    return json.MarshalIndent(movie, "", "  ")
}

// json 解码
//func jsonUnMarshal(data []byte) error {
//    return json.Unmarshal(data)
//}

func main()  {
    data, err := jsonMarshal1(movies)

    if err != nil{
        log.Fatalf("JSON marshaling failed: %s", err)
    }

    fmt.Printf("jsonMarshal1 --- %s\n", data)

    data2, err2 := jsonMarshal2(movies)
    if err2 != nil {
        log.Fatalf("JSON marshaling failed: %s", err2)
    }
    fmt.Printf("jsonMarshal2 --- \n%s\n", data2)

}
