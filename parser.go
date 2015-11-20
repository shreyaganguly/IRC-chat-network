package main

import(
  "strings"
  "net"
  "log"
)
var c Client
var cd ChannelDetails
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
func parseForJoin(args []string) []string{
  var tokens []string
  if len(args) == 1 {
    tokens = strings.Split(args[0],",")
  }
  if len(tokens) == 0 {
  for i:=0; i< len(args) ; i++ {
   _, ok := ChannelMap[args[i]]
   if(ok == false){
     cd.Topic = ":There is no topic"
     ChannelMap[args[i]] = cd
   }
 }
 return args
 } else {
   for i:=0; i< len(tokens) ; i++ {
    _, ok := ChannelMap[tokens[i]]
    if(ok == false){
      cd.Topic = ":There is no topic"
      ChannelMap[tokens[i]] = cd
    }
  }
 }
 return tokens
}
func parser(s string,con net.Conn)(key string,value []string){
tokens := strings.Split(s, " ")
keyword := tokens[0]
var channelnames,arguments,arg_array,channels []string
if tokens[0] == "JOIN" {
  arg_array = strings.Split(s, tokens[0]+" ")
  channelnames = arg_array[1:]
} else {
arguments = tokens[1:]
}
switch keyword {
case "USER":
  parseForUser(arguments,con)
case "NICK":
  flag := parseForNick(arguments)
  if flag == false{
    log.Println("Nick name already exists")
    }

case "JOIN":
  channels = parseForJoin(channelnames)
 return keyword,channels
}
return keyword,arguments
}
