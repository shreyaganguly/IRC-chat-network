package main

import(
  "strings"
)
func parser(s string)(key string,value []string) {
  tokens := strings.Split(s, " ")
keyword := tokens[0]
arguments := tokens[1:]
return keyword, arguments

}
