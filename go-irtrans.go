package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	irtrans_ip := "10.0.0.7"
	irtrans_port := "21000"
	//retVal := []string{}
	conn, err := net.Dial("tcp", irtrans_ip+":"+irtrans_port)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	} else {
		if len(os.Args) == 1 {
			fmt.Println("list remotes and cmds: go run go-irtrans.go list\nrun cmd for remote: go run go-irtrans.go remote cmd")
		} else {
			if len(os.Args) == 2 {
				if os.Args[1] == "list" {
					gR := getRemotes(conn)
					gR_arr := strings.Split(gR, ",")
					for _, remote := range gR_arr {
						gC := getCMDs(conn, remote)
						fmt.Println(remote + ":" + gC)
					}
				} else {
					fmt.Println("list remotes and cmds: go run go-irtrans.go list\nrun cmd for remote: go run go-irtrans.go remote cmd")
				}

			} else if len(os.Args) == 3 {
				arg_remote := os.Args[1]
				arg_cmd := os.Args[2]
				sC := sendCMD(conn, arg_remote, arg_cmd)
				fmt.Println(sC)
			}
		}

	}
	conn.Close()
}

func getRemotes(conn net.Conn) string {
	arr := []string{}
	retVal := ""
	_, err := conn.Write([]byte("Agetremotes 0"))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	line = strings.TrimRight(line, " \t\r\n")
	//fmt.Println(line)
	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		arr = strings.Split(line, " ")
		arr = strings.Split(arr[2], ",")

		o, _ := strconv.Atoi(arr[0])
		m, _ := strconv.Atoi(arr[1])
		r, _ := strconv.Atoi(arr[2])

		for index, remote := range arr[3:] {
			retVal = retVal + remote
			if index < (r - 1) {
				retVal = retVal + ","
			}
		}

		if (o + r) != m {
			_, err := conn.Write([]byte("Agetremotes " + arr[2]))
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}
			reader := bufio.NewReader(conn)
			line, err := reader.ReadString('\n')
			line = strings.TrimRight(line, " \t\r\n")
			//fmt.Println(line)
			if err != nil {
				fmt.Println(err)
				return "error"
			} else {
				arr = strings.Split(line, " ")
				arr = strings.Split(arr[2], ",")
				o, _ = strconv.Atoi(arr[0])
				m, _ = strconv.Atoi(arr[1])
				r, _ = strconv.Atoi(arr[2])
				for _, remote := range arr[3:] {
					retVal = retVal + "," + remote
				}
			}
		}

		return retVal
	}
}

func getCMDs(conn net.Conn, remote string) string {
	arr := []string{}
	retVal := ""
	_, err := conn.Write([]byte("Agetcommands " + remote + ",0"))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	line = strings.TrimRight(line, " \t\r\n")
	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		arr = strings.Split(line, " ")
		arr = strings.Split(arr[2], ",")

		o, _ := strconv.Atoi(arr[0])
		m, _ := strconv.Atoi(arr[1])
		r, _ := strconv.Atoi(arr[2])

		for index, cmd := range arr[3:] {
			retVal = retVal + cmd
			if index < (r - 1) {
				retVal = retVal + ","
			}
		}
		if (o + r) != m {
			_, err := conn.Write([]byte("Agetcommands " + remote + "," + arr[2]))
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}
			reader := bufio.NewReader(conn)
			line, err := reader.ReadString('\n')
			line = strings.TrimRight(line, " \t\r\n")
			if err != nil {
				return "error"
			} else {
				arr = strings.Split(line, " ")
				arr = strings.Split(arr[2], ",")
				o, _ = strconv.Atoi(arr[0])
				m, _ = strconv.Atoi(arr[1])
				r, _ = strconv.Atoi(arr[2])
				for _, cmd := range arr[3:] {
					retVal = retVal + "," + cmd
				}
			}
		}
		return retVal
	}
}

func sendCMD(conn net.Conn, remote string, cmd string) string {
	arr := []string{}
	retVal := ""
	_, err := conn.Write([]byte("Asnd " + remote + "," + cmd))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	line = strings.TrimRight(line, " \t\r\n")
	if err != nil {
		fmt.Println(err)
		return "error"
	} else {
		arr = strings.Split(line, " ")
		for index, r := range arr[2:] {
			retVal = retVal + r
			if len(arr[2:]) > (index + 1) {
				retVal = retVal + " "
			}
		}
		return retVal
	}
}
