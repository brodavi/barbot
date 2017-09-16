package main

import (
        "fmt"
        "net/http"
        "os"
        "time"
        "github.com/stianeikeland/go-rpio"
)

var pins [8]uint8

pins[0] = rpio.Pin(21) // 40 // Tequila
pins[1] = rpio.Pin(20) // 38 // Vodka
pins[2] = rpio.Pin(16) // 36 // Gin
pins[3] = rpio.Pin(12) // 32 // Rum
pins[4] = rpio.Pin(6) // 31 // Triplesec
pins[5] = rpio.Pin(13) // 33 // Whiskey
pins[6] = rpio.Pin(19) // 35 // Sweetnsour
pins[7] = rpio.Pin(26) // 37 // Coke

func toggle()

func pour(ingredient string, ms int) {
  switch ingredient {
    case "tequila":
      Tequila.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Tequila.Toggle()
    case "vodka":
      Vodka.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Vodka.Toggle()
    case "gin":
      Gin.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Gin.Toggle()
    case "rum":
      Rum.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Rum.Toggle()
    case "triplesec":
      Triplesec.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Triplesec.Toggle()
    case "whiskey":
      Whiskey.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Whiskey.Toggle()
    case "sweetnsour":
      Sweetnsour.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Sweetnsour.Toggle()
    case "coke":
      Coke.Toggle()
      time.Sleep(time.Millisecond * time.Duration(ms * Multiple))
      Coke.Toggle()
  }
}

func makeDrink(drinkInt string) {

    switch drinkInt {
        case "legspreader": //legspreader
            fmt.Println("making legspreader")
            pour("tequila", 1000)
            pour("vodka", 1000)
            pour("gin", 1000)
            pour("rum", 1000)
            fmt.Println("made legspreader")
        case "flyingdutchman": //flyingdutchman
            fmt.Println("making flyingdutchman")
            pour("gin", 2000)
            pour("triplesec", 500)
            fmt.Println("made flyingdutchman")
        case "southbank": //southbank
            fmt.Println("making southbank")
            pour("rum", 1000)
            pour("gin", 1000)
            fmt.Println("made southbank")
        case "elgringo": //elgringo
            fmt.Println("making elgringo")
            pour("rum", 1000)
            pour("tequila", 1000)
            pour("triplesec", 1000)
            fmt.Println("made elgringo")
      	case "rumncoke": //rumncoke
            fmt.Println("making rumncoke")
            pour("rum", 2000)
            pour("coke", 2000)
            fmt.Println("made rumncoke")
        case "jackncoke": //jackncoke
            fmt.Println("making jackncoke")
            pour("whiskey", 2000)
            pour("coke", 2000)
            fmt.Println("made jackncoke")
        case "longisland": //longisland
            fmt.Println("made longisland")
            pour("vodka", 1000)
            pour("tequila", 1000)
            pour("rum", 1000)
            pour("triplesec", 1000)
            pour("sweetnsour", 1500)
            fmt.Println("made longisland")
        case "test01": //test tequila
            pour("tequila", 1500)
            fmt.Println("tested tequila")
        case "test02": //test vodka
            pour("vodka", 1500)
            fmt.Println("tested vodka")
        case "test03": //test gin
            pour("gin", 1500)
            fmt.Println("tested gin")
        case "test04": //test rum
            pour("rum", 1500)
            fmt.Println("tested rum")
        case "test05": //test triplesec
            pour("triplesec", 1500)
            fmt.Println("tested triplesec")
        case "test06": //test whisky
            pour("whisky", 1500)
            fmt.Println("tested whisky")
        case "test07": //test sweetnsour
            pour("sweetnsour", 1500)
            fmt.Println("tested sweetnsour")
        case "test08": //test coke
            pour("coke", 1500)
            fmt.Println("tested coke")
    }
}

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if (r.Method == "OPTIONS") {
      w.Header().Set("Access-Control-Allow-Headers", "Authorization")
    } else {
      makeDrink(r.URL.Path[1:])
      fmt.Fprintf(w, "Drink complete")
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
    //init
    if err := rpio.Open(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer rpio.Close()

    fmt.Println("Tequila")
    Tequila.Output()
    Tequila.High()

    fmt.Println("Vodka")
    Vodka.Output()
    Vodka.High()

    fmt.Println("Gin")
    Gin.Output()
    Gin.High()

    fmt.Println("Rum")
    Rum.Output()
    Rum.High()

    fmt.Println("Triplesec")
    Triplesec.Output()
    Triplesec.High()

    fmt.Println("Whiskey")
    Whiskey.Output()
    Whiskey.High()

    fmt.Println("Sweetnsour")
    Sweetnsour.Output()
    Sweetnsour.High()

    fmt.Println("Coke")
    Coke.Output()
    Coke.High()

    fmt.Println("Listening on port 8080")

    //start httplistener
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
