package handling

import (
	"fmt"
	"net"
)

func ConnectToServer(ip string) net.Conn {
    address := fmt.Sprint(ip, ":8642")

    //connect to server
    conn, err := net.Dial("tcp", address)
    if err != nil {
        panic(err)
    }
    return conn
}

func GetMessages() {
    
}
