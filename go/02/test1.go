package main

import "fmt"

func main() {
	fmt.Printf("%c[47;36m%s%c[0m\n",0x1B,"呐呐呐",0x1B)
}
