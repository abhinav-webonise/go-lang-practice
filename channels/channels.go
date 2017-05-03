package main

import "fmt"

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

}
