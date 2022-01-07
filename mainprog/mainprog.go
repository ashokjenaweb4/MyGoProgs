package main
import (
    "fmt"
    "timu/emailservice"
)

func main() {
    // Get a server ip address and print it.
    message := emailservice.GetServerIP("localhost")
    fmt.Println(message)
}