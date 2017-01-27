# iletimerkezi.go

Send SMS with Ileti Merkezi API

## Requirements

  - Ileti Merkezi Account: https://www.iletimerkezi.com
  - Valid Sender

## URL Scheme:

 - Variables:
    - username `your phone number`
    - password `your password`
    - text   `sms content`
    - receipents `receipents can be multiple, is allowed`
    - sender `required valid sender ( sms title )`

```text
https://api.iletimerkezi.com/v1/send-sms/get/?
username=KULLANICI_ADI&
password=KULLANICI_SIFRESI&
text=Lorem%20ipsum%20dolor%20sit%20amet&
receipents=5301234569, 5301234570&
sender=ILETI%20MRKZI
```


## Usage

```go
package iletimerkezi_test

import (
    "fmt"
    "./iletimerkezi"
)

func main() {
    user_name     := "your_phone_number"
    password      := "your password"
    msg_content   := "This is message content"
    receipents    := "number, number2, number3" // multiple receipents is allowed
    sender_title  := "SMS SNDR"
    
    // function always return XML
    data := send_sms(user_name, password, msg_content, receipents, sender_title)
    
    fmt.Println(data)
}
```
