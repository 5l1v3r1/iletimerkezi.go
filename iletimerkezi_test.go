package iletimerkezi_test

import (
    "fmt"
)

func main() {
    user_name     := "your_phone_number"
    password      := "your password"
    msg_content   := "This is message content"
    receipents    := "number, number2, number3" // multiple receipents is allowed
    sender_title  := "SMS SNDR"
    data := send_sms(user_name, password, msg_content, receipents, sender_title)
    fmt.Println(data)
}
