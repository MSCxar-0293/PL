package main

import (
	"fmt"
)

func main() {
	var topic string
	title := titleinit(topic)
	body := bodyinit(topic)
	fmt.Printf("欢迎来到女生自用低端slscq！\n请输入主题：\n")
	fmt.Scanf("%s", &topic)
	fmt.Printf("《", title, "》\n\n", body)
}
