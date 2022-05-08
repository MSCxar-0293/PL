package main

import (
  "fmt"
  "time"
)

func main(){
     fmt.Printf("欢迎使用并联电路电阻计算器v1.1~\n")
     fmt.Printf("河神：你掉的是这个总电阻呢，还是这个分电阻呢？\n")
     fmt.Printf("> 总电阻\n> 分电阻\n[hint: 请直接输入选项文字。]\n")
     var choice string
     var cal float64
     var ro float64
     var rt float64
     var ra float64

     fmt.Scanf("%s",&choice)

     switch choice {
	case "总电阻":
	fmt.Printf("请按格式输入分电阻1和2: \n[格式: 分电阻1 分电阻2]\n")
	fmt.Scanf("%f %f",&ro,&rt)
	cal = allr(ro, rt)
	fmt.Printf("总电阻为：%f\n", cal)
	fmt.Printf("计算器将于5秒后自动关闭。\n")
	time.Sleep(time.Duration(5000)*time.Millisecond)
	case "分电阻":
	fmt.Printf("请按格式输入总电阻和已知分电阻：\n[格式: 总电阻 已知分电阻]\n")
	fmt.Scanf("%f %f",&ra, &ro)
	cal = oner(ra, ro)
	fmt.Printf("你要求的分电阻为：%f\n", cal)
	fmt.Printf("计算器将于5秒后自动关闭。\n")
	time.Sleep(time.Duration(5000)*time.Millisecond)
     }
}

func allr(r1,r2 float64) float64{
     var m float64 = r1 + r2
     var s float64 = r1 * r2
     var result float64 = s/m

     return result
}

func oner(r,r1 float64) float64{
     var m float64 = r1 - r
     var s float64 = r * r1
     var result float64 = s/m

     return result
}
