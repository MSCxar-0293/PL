package main

import "fmt"

func main() {
     for true  {
     fmt.Printf("%c[47;31m%s%c[0m\n",0x1B,"■■■■■■■■",0x1B)
     fmt.Printf("%c[47;33m%s%c[0m\n",0x1B,"■■■■■■■■■",0x1B)
     fmt.Printf("%c[47;32m%s%c[0m\n",0x1B,"■■■■■■■■■■",0x1B)
     fmt.Printf("%c[47;36m%s%c[0m\n",0x1B,"■■■■■■■■■■■",0x1B)
     fmt.Printf("%c[47;34m%s%c[0m\n",0x1B,"■■■■■■■■■■",0x1B)
     fmt.Printf("%c[47;35m%s%c[0m\n",0x1B,"■■■■■■■■■",0x1B)
     fmt.Printf("%c[47;37m%s%c[0m\n",0x1B,"■■■■■■■■",0x1B)
   }
}
