package main

import (
    "time"
    "strconv"
    "fmt"
    "github.com/vyzigold/loki-client/pkg/loki"
)

func main () {
    client, err := loki.CreateClient("http://localhost:3100")
    if err == nil {
        fmt.Println("Client successfuly created")
        fmt.Println(client)
    }
    label1 := loki.Label {
        Key: "testkey2",
        Value: "testvalue5",
    }
    message1 := loki.Message {
        Time: strconv.FormatInt(time.Now().UnixNano(), 10),
        Message: "{\"simple\": [\"json\", \"string\"]}",
    }
    message2 := loki.Message {
        Time: strconv.FormatInt(time.Now().UnixNano(), 10),
        Message: "not json message",
    }
    labels := []loki.Label{label1}
    messages := []loki.Message{message1, message2}
    err = client.Send(labels, messages)
    if err == nil {
        fmt.Println("Message successfuly sent")
    } else {
        fmt.Println(err)
    }
    response, err := client.Query("{testkey2=~\"test.*\"}")
    fmt.Println(response)
}
