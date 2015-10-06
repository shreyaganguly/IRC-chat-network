package main

import(
  "log"
  "bufio"
  "net"
)

func HandleUSER(con net.Conn, args []string) {
	log.Println("the arguments passed after USER are", args)
}

func HandleNICK(con net.Conn, args []string) {
	log.Println("the arguments passed after NICK are", args)
}

type commandHandler func(con net.Conn, args []string)

var cmdMap = map[string]commandHandler{
	"USER": HandleUSER,
	"NICK": HandleNICK,
}

func ConnectionHandler(con net.Conn) {

	log.Println("Request coming from", con.RemoteAddr())
	commands := bufio.NewScanner(con)
	for commands.Scan() {
		keyword,arguments := parser(commands.Text())
		handler, ok := cmdMap[keyword]
		if ok == false {
			log.Println("Sorry command not recongnized")
		} else {
			handler(con, arguments)
		}
	}
}
