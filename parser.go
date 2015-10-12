package main

import(
  "strings"
  "net"
  "log"
)
var c Client
func parseForUser(args []string,con net.Conn){
  c.RemoteAddress = con.RemoteAddr().String()
  c.UserName = args[0]
  c.HostIP = args[2]
  val := strings.Split(args[3], ":")
  c.RealName = val[1]

}

func parseForNick(args []string) bool {
   _, ok := ClientMap[args[0]]
   if(ok == true){
     return false
   }
   ClientMap[args[0]] = c
   return true
}
func parser(s string,con net.Conn)(key string,value []string){
tokens := strings.Split(s, " ")
keyword := tokens[0]
arguments := tokens[1:]
switch keyword {
case "USER":
  parseForUser(arguments,con)
case "NICK":
  flag := parseForNick(arguments)
  if flag == false{
    log.Println("Nick name already exists")
    }
 }
 return keyword,arguments
}
