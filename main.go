package main

import (
    "time"
    "strconv"
    "fmt"
    "github.com/vyzigold/loki-client/pkg/loki"
)

var count int = 0
func sendSomething(client *loki.LokiClient) {
    count++
    labels := make(map[string]string)
    labels["labelkey"]="labelvalue"
    labels["level"]=strconv.Itoa(count%4)
    message1 := loki.Message {
        Time: strconv.FormatInt(time.Now().UnixNano(), 10),
        Message: "{\"simple\": [\"json\", \"string\"]}",
    }
    message2 := loki.Message {
        Time: strconv.FormatInt(time.Now().UnixNano(), 10),
        Message: "not json message, cnt=" + strconv.Itoa(count) + ", level=" + strconv.Itoa(count%4),
    }
    messages := []loki.Message{message1, message2}
    client.AddStream(labels, messages)
}

func main () {
    client, err := loki.CreateClient("http://10.62.117.69:3100", 4, 1000000000)
    if err == nil {
        fmt.Println("Client successfuly created")
    }
    for i := 100000000000; i > 0; i-- {
        sendSomething(client)
        time.Sleep(time.Second)
    }
    response, err := client.Query("{labelkey=~\"lab.*\"}")
    fmt.Println("\nthe response of the query currently is:")
    fmt.Println(response)
    client.Shutdown()
}
