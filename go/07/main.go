package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var wish int
	var in int
	var exit string
	times := 0
	min := 1
	max := 1000
	minT := strconv.Itoa(min)
	maxT := strconv.Itoa(max)
	compare := minT + "-" + maxT
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("来！和我来一场惊心动魄的猜数字游戏吧！\n")
	wish = rand.Intn(1000) + 1
	fmt.Println("命运之轮已经抉择好了此刻的神圣！\n输入任意一个1-1000的数字开始游戏~\n")
	INPUT: fmt.Scanf("%d", &in)
	times = times + 1
	switch {
	case in == wish : 
		fmt.Println("真厉害~答对了~\n这一轮回，你一共猜了", times, "次哦~\n<按回车键退出游戏。>")
		fmt.Scanf("%s", &exit)
	case in > 1000 || in < 0 : 
		fmt.Println("...你输入的是什么鬼啊！重新来！\n即使是这一次也会被记录在猜数次数里哦？\n")
		goto INPUT
	default:
		switch {
		case in > wish && in < max: 
		max = in
		case in < wish && in > min: 
		min = in
		}
		minT := strconv.Itoa(min)
		maxT := strconv.Itoa(max)
		compare = minT + "-" + maxT
		fmt.Println("答错了！再猜一次吧~\n你已经试了", times, "次啦~\n当前范围为：", compare)
		goto INPUT
	}
}
