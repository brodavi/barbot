package main

import (
        "fmt"
        "net/http"
        "net/url"
        "os"
        "time"
        "github.com/stianeikeland/go-rpio"
)

var pins [8]rpio.Pin
var Multiple = 2

func toggle(pin int) {
  pins[pin].Toggle()
}

func pour(pin int, ms int) {
  pins[pin].Toggle()
  time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
  pins[pin].Toggle()
}

func makeDrink(drinkInt string) {
  switch drinkInt {
    case "legspreader": //legspreader
        fmt.Println("making legspreader")
        pour(0, 1000)
        pour(1, 1000)
        pour(2, 1000)
        pour(3, 1000)
        fmt.Println("made legspreader")
    case "flyingdutchman": //flyingdutchman
        fmt.Println("making flyingdutchman")
        pour(2, 2000)
        pour(4, 500)
        fmt.Println("made flyingdutchman")
    case "southbank": //southbank
        fmt.Println("making southbank")
        pour(3, 1000)
        pour(2, 1000)
        fmt.Println("made southbank")
    case "elgringo": //elgringo
        fmt.Println("making elgringo")
        pour(3, 1000)
        pour(0, 1000)
        pour(4, 1000)
        fmt.Println("made elgringo")
  	case "rumncoke": //rumncoke
        fmt.Println("making rumncoke")
        pour(3, 2000)
        pour(7, 2000)
        fmt.Println("made rumncoke")
    case "jackncoke": //jackncoke
        fmt.Println("making jackncoke")
        pour(5, 2000)
        pour(7, 2000)
        fmt.Println("made jackncoke")
    case "longisland": //longisland
        fmt.Println("made longisland")
        pour(1, 1000)
        pour(0, 1000)
        pour(3, 1000)
        pour(4, 1000)
        pour(6, 1500)
        fmt.Println("made longisland")
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  if (r.Method == "OPTIONS") {
    w.Header().Set("Access-Control-Allow-Headers", "Authorization")
  } else {
    var u = url.Parse(r.URL)
    fmt.Println(u)
    fmt.Println(u.Path)
    fmt.Println(r.URL.Path[1:])
    fmt.Println(r.URL.Path[2:])
    // if (r.URL.Path[1:] == "make") {
    //   makeDrink(r.URL.Path[2:])
    // } else if (r.URL.Path[1:] == "test") {
    //   test(r.URL.Path[2:])
    // }
    // fmt.Fprintf(w, "Drink complete")
  }
}

func initBoard() {
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer rpio.Close()
}

func main() {
  pins[0] = rpio.Pin(21) // 40 // Tequila
  pins[1] = rpio.Pin(20) // 38 // Vodka
  pins[2] = rpio.Pin(16) // 36 // Gin
  pins[3] = rpio.Pin(12) // 32 // Rum
  pins[4] = rpio.Pin(6) // 31 // Triplesec
  pins[5] = rpio.Pin(13) // 33 // Whiskey
  pins[6] = rpio.Pin(19) // 35 // Sweetnsour
  pins[7] = rpio.Pin(26) // 37 // Coke

  if err := rpio.Open(); err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  defer rpio.Close()

  fmt.Println("0")
  pins[0].Output()
  pins[0].High()

  fmt.Println("1")
  pins[1].Output()
  pins[1].High()

  fmt.Println("2")
  pins[2].Output()
  pins[2].High()

  fmt.Println("3")
  pins[3].Output()
  pins[3].High()

  fmt.Println("4")
  pins[4].Output()
  pins[4].High()

  fmt.Println("5")
  pins[5].Output()
  pins[5].High()

  fmt.Println("6")
  pins[6].Output()
  pins[6].High()

  fmt.Println("7")
  pins[7].Output()
  pins[7].High()

  fmt.Println("Listening on port 8080")

  //start httplistener
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}