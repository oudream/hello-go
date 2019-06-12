package multi_server1

import (
	"fmt"
	"net"
	"net/http"
)

func HelloMultiServer1() {

	http.HandleFunc("/hi", myHandler1)
	//因为会阻塞，所以需要新开goroutine进行监听
	go http.ListenAndServe(":8080", nil)

	server, err := net.Listen("tcp", ":5080")
	if err != nil {
		return
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	//do somethings
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

func myHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
