package main

import "fmt"
import "time"


func ping(pings chan<- string, msg string) {
  pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
  msg := <-pings
  pongs <- msg
}


func main() {
  message := make(chan string)

  go func() {
    message <- "ping"
  }()

  messages := make(chan string, 2)

  messages <- "buffered"
  messages <- "channel"

  msg := <-message
  fmt.Println(msg)

  fmt.Println(<-messages)
  fmt.Println(<-messages)

  // channel buffers are synchronous. I can not print 2nd value directly

  done := make(chan bool, 1)
  go worker(done)

  <-done

  fmt.Println("What about me?")


  // Channel Directions

  pings := make (chan string, 1)
  pongs := make (chan string, 1)

  ping(pings, "passed message")
  pong(pings, pongs)

  fmt.Println(<-pongs)

}


// Channel Synchronisation
func worker(done chan bool) {
  fmt.Print("working")
  //fmt.Println(time.Second)
  time.Sleep(time.Second)
  fmt.Println("done")

  done <- false

  fmt.Println("would you like to print me?")
}
