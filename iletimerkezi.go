package iletimerkezi

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
)

func send_sms(username string, password string, message string, receipents string, sender_title string) string {
    var url = "https://api.iletimerkezi.com/v1/send-sms/get/?"
    var text = "username="+ username +"&password="+ password +"&text="+ strings.Replace(message, " ", "%20", -1) +""
    var info = "&receipents="+ receipents +"&sender="+ strings.Replace(sender_title, " ", "%20", -1) +""
    
    var build_variables = url + text + info

    response, err := http.Get(build_variables)
    defer response.Body.Close()

    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }

    return fmt.Sprintf("%s", body)
}
