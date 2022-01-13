package main

import (
	"clat-client/handling"
	"fmt"
    "flag"
)

func main() {
    fmt.Println("idk shit about what i am doing...")
    
    // setup flags
    var serverIp string
    var serverPort string
    
    flag.StringVar(&serverIp, "i", "127.0.0.1", "specify ip address. default= localhost/127.0.0.1")
    flag.StringVar(&serverPort, "p", "8642", "specify the port. default is 8642")
    flag.Parse()

    // connect to server
    conn := handling.ConnectToServer("127.0.0.1")
    fmt.Fprint(conn, "deez\n")
}
