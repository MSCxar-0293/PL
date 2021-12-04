package main

import "fmt"

func main(){
     fmt.Printf("欢迎使用并联电路电阻计算器v1.0~\n")
     fmt.Printf("河神：你掉的是这个总电阻呢，还是这个分电阻呢？\n")
     fmt.Printf("> 总电阻\n> 分电阻\n[hint: 请直接输入选项文字。]\n")
     var choice string
     var cal int
     var ro int
     var rt int
     var ra int

     fmt.Scanf("%s",&choice)

     switch choice {
	case "总电阻":
	fmt.Printf("请按格式输入分电阻1和2: \n[格式: 分电阻1 分电阻2]\n")
	fmt.Scanf("%d %d",&ro,&rt)
	cal = allr(ro, rt)
	fmt.Printf("总电阻为：%d\n", cal)
	case "分电阻":
	fmt.Printf("请按格式输入总电阻和已知分电阻：\n[格式: 总电阻 已知分电阻]\n")
	fmt.Scanf("%d %d",&ra, &ro)
	cal = oner(ra, ro)
	fmt.Printf("你要求的分电阻为：%d\n", cal)
     }
}

func allr(r1,r2 int) int{
     var m int = r1 + r2
     var s int = r1 * r2
     var result int = s/m

     return result
}

func oner(r,r1 int) int{
     var m int = r1 - r
     var s int = r * r1
     var result int = s/m

     return result
}
