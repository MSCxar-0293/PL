package main

import (
	"fmt"
	"time"
	)

func main(){
	var age int
	fmt.Println("小朋友，你几岁啦？")
	fmt.Scanf("%d",&age)
	time.Sleep(time.Duration(800)*time.Millisecond)
	fmt.Printf("对方输入中...\n\n")
	time.Sleep(time.Duration(1200)*time.Millisecond)
	if age >= 0 && age <= 10{
		fmt.Printf("我对%d岁的小p孩不感兴趣，请不要再来了（无慈悲）\n",age);
	}else if age > 10 && age <= 18{
		fmt.Printf("%d岁了？快来让姐姐看看你发育得怎么样 ^q^ #嘶哈嘶哈嘶哈\n",age);
	}else if age > 18 && age <= 30{
		fmt.Printf("已经%d岁了呢..\n有车吗有房吗有工作吗结婚了吗有孩子了吗\n\n有没有被烦到呢？..嘻嘻...\n",age);
	}else if age > 30 && age <= 50{
		fmt.Printf("呜哇...已经%d岁了吗..\n岁月不饶人啊..\n人到中年肯定没有以前精力旺盛了...\n累了就休息一下吧。\n",age);
	}else if age > 50 && age <= 80{
		fmt.Printf("..是在下失敬了..\n重阳节我会试着给您送礼物的。\n",age);
	}else if age > 80 && age <= 120{
		fmt.Printf("祝阁下福如东海寿比南山！\n",age);
	}else if age > 120{
	        fmt.Printf("...%d岁的你，竟然还能使用电子设备么...\n",age);
}
}
