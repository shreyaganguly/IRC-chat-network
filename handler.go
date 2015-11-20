package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func HandleUSER(con net.Conn, args []string) {
	log.Println("the arguments passed after USER are", args)
}

func HandleNICK(con net.Conn, args []string) {
	log.Println("the arguments passed after NICK are", args)
	present_time := (time.Now()).Truncate(time.Duration(time.Second))
	var time = present_time.String()
	server_creation_time := strings.TrimSuffix(time, "IST")
	servername, err := os.Hostname()
	if err != nil {
		log.Fatal("ERROR", err)
	}
	replymessage1 := fmt.Sprintf(":%s 001 %s %s :Welcome to this IRC server %s!%s@", servername, args[0], args[0], args[0], ClientMap[args[0]].UserName)
	replymessage2 := fmt.Sprintf(":%s 003 %s %s :This server was created %s", servername, args[0], args[0], server_creation_time)
	replymessage3 := fmt.Sprintf(":%s 002 %s %s :Your host is %s, running version 0.01dev", servername, args[0], args[0], servername)
	replymessage4 := fmt.Sprintf(":%s 004 %s %s :%s 0.01dev aAbBcCdDeEfFGhHiIjkKlLmMnNopPQrRsStUvVwWxXyYzZ0123459*@ bcdefFhiIklmnoPqstv", servername, args[0], args[0], servername)
	replymessage5 := fmt.Sprintf(":%s 375 %s :- Message of the Day"+"\r\n"+":%s 372 %s :- Do the dance see the"+"\r\n"+":%s 376 %s :- End of /MOTD command.", servername, args[0], servername, args[0], servername, args[0])
	connection_acceptance_reply := "\r\n" + replymessage1 + "\r\n" + replymessage3 + "\r\n" + replymessage2 + "\r\n" + replymessage4 + "\r\n" + replymessage5 + "\r\n"
	con.Write([]byte(connection_acceptance_reply))

}
func HandlePING(con net.Conn, args []string) {
	log.Println("the arguments passed after PING are", args)
	var replytoPING string
	if len(args) == 1 {
		replytoPING = fmt.Sprintf("PONG   :[\"%s\\r\"]", args[0])
	}
	if len(args) > 1 {
		replytoPING = fmt.Sprintf("PONG   :[\"%s\"]", args[0])
	}
	con.Write([]byte(replytoPING + "\r\n"))
}
func HandleJOIN(con net.Conn, args []string) {
	log.Println("the arguments passed after JOIN are", args)
	servername, err := os.Hostname()
	if err != nil {
		log.Fatal("ERROR", err)
	}
	var reply1, reply2, reply3, reply4 string

	for _, v := range args {
		reply1 = fmt.Sprintf(" JOIN :%s", v)
		reply2 = fmt.Sprintf(":%s 332  %s %s", servername, v, cd.Topic)
		reply3 = fmt.Sprintf(":%s 353  = %s :", servername, v)
		reply4 = fmt.Sprintf(":%s 366  %s  :End of /NAMES list.", servername, v)
		con.Write([]byte(reply1 + "\r\n" + reply2 + "\r\n" + reply3 + "\r\n" + reply4 + "\r\n"))
	}
}
func HandleTOPIC(con net.Conn, args []string) {
  log.Println("the arguments passed after TOPIC are", args)
	var reply string
	if len(args) > 1 {
		reply = fmt.Sprintf(" TOPIC %s %s", args[0], args[1])
		con.Write([]byte(reply + "\r\n"))
	} else {
		con.Write([]byte("\r"))
	}
}

type commandHandler1 func(con net.Conn, args []string)

var cmdMap = map[string]commandHandler1{
	"USER":  HandleUSER,
	"NICK":  HandleNICK,
	"PING":  HandlePING,
	"JOIN":  HandleJOIN,
	"TOPIC": HandleTOPIC,
}

func ConnectionHandler(con net.Conn) {

	log.Println("Request coming from", con.RemoteAddr())
	commands := bufio.NewScanner(con)
	for commands.Scan() {
		keyword, arguments := parser(commands.Text(), con)
		handler, ok := cmdMap[keyword]
		if ok == false {
			log.Println("Sorry command not recognized")
			log.Println(keyword)
			log.Println("****************")
			log.Println(arguments)
		} else {
			handler(con, arguments)
		}
	}

}
