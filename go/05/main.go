package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// 向客户端写入这些数据，以便客户端可以填写文本并提交
var indexHTML = `<html>
<head>
	<meta http-equiv="Content-type" content="text/html; charset=utf-8">
	<title>测试</title>
</head>
<body>
	<form action="/page" method="post">
		月基本工资：<br>
		<input name="monthpay" type="text"><br>
		个人养老保险税率：<input name="EN" type="text">  企业养老保险税率：<input name="EEN" type="text"><br>
		个人医疗保险税率：<input name="ME" type="text">  企业医疗保险税率：<input name="EME" type="text"><br>
		个人工伤保险税率：<input name="EM" type="text">  企业工伤保险税率：<input name="EEM" type="text"><br>
		个人生育保险税率：<input name="MA" type="text">  企业生育保险税率：<input name="EMA" type="text"><br>
		个人失业保险税率：<input name="UN" type="text">  企业失业保险税率：<input name="EUN" type="text"><br>
		个人公积金税率：&nbsp;&nbsp;&nbsp;<input name="AF" type="text">  企业公积金税率： &nbsp;&nbsp;&nbsp;<input name="AF" type="text"><br>
		<input type="submit" value="提交">
	</form>
</body>
</html>`

// 用于将页面重定向到主页
var redirectHTML = `<html>
<head>
	<meta http-equiv="Content-type" content="text/html; charset=utf-8">
	<meta http-equiv="Refresh" content="0; url={{.}}">
</head>
<body></body>
</html>`

// 处理主页请求
func index(w http.ResponseWriter, r *http.Request) {
	// 向客户端写入我们准备好的页面
	fmt.Fprintf(w, indexHTML)
}

// 处理客户端提交的数据
func page(w http.ResponseWriter, r *http.Request) {
	//月基本工资
	var monthpay float64
	//月数
	var monthnum float64 = 12
	//养老保险
	//var EN float64
	var ratio_EN, Eratio_EN float64
	//医疗保险
	//var ME float64
	var ratio_ME, Eratio_ME float64
	//失业保险
	//var UN float64
	var ratio_UN, Eratio_UN float64
	//工伤保险
	//var EM float64
	var ratio_EM, Eratio_EM float64
	//生育保险
	//var MA float64
	var ratio_MA, Eratio_MA float64
	//公积金
	//var AF float64
	var ratio_AF, Eratio_AF float64
	//per个人所得税pertax个人所得税率，su速扣计算数
	var per,pertax,su float64

	// 我们规定必须通过 POST 提交数据
	if r.Method == "POST" {
		// 解析客户端请求的信息
		if err := r.ParseForm(); err != nil {
			log.Println(err)
		}
		// 获取客户端输入的内容
		ENS := r.Form.Get("EN")
		EMS := r.Form.Get("EM")
		UNS := r.Form.Get("UN")
		MAS := r.Form.Get("MA")
		AFS := r.Form.Get("AF")
		MES := r.Form.Get("ME")
		monthpayS := r.Form.Get("monthpay")
		EENS := r.Form.Get("EEN")
		EEMS := r.Form.Get("EEM")
		EUNS := r.Form.Get("EUN")
		EMAS := r.Form.Get("EMA")
		EAFS := r.Form.Get("EAF")
		EMES := r.Form.Get("EME")
		//userText := r.Form.Get("usertext")
		ratio_EN, _ = strconv.ParseFloat(ENS, 64)
		ratio_EM, _ = strconv.ParseFloat(EMS, 64)
		ratio_UN, _ = strconv.ParseFloat(UNS, 64)
		ratio_MA, _ = strconv.ParseFloat(MAS, 64)
		ratio_AF, _ = strconv.ParseFloat(AFS, 64)
		ratio_ME, _ = strconv.ParseFloat(MES, 64)
		monthpay, _ = strconv.ParseFloat(monthpayS, 64)
		Eratio_EN, _ = strconv.ParseFloat(EENS, 64)
		Eratio_EM, _ = strconv.ParseFloat(EEMS, 64)
		Eratio_UN, _ = strconv.ParseFloat(EUNS, 64)
		Eratio_MA, _ = strconv.ParseFloat(EMAS, 64)
		Eratio_AF, _ = strconv.ParseFloat(EAFS, 64)
		Eratio_ME, _ = strconv.ParseFloat(EMES, 64)

		switch{
			case monthpay<=5000:pertax=0;su=0
			case monthpay>5000&&monthpay<=8000:pertax=0.03;su=0
			case monthpay>8000&&monthpay<=17000:pertax=0.1;su=210
			case monthpay>17000&&monthpay<=30000:pertax=0.2;su=1410
			case monthpay>30000&&monthpay<=40000:pertax=0.25;su=2660
			case monthpay>40000&&monthpay<=60000:pertax=0.3;su=4410
			case monthpay>60000&&monthpay<=85000:pertax=0.35;su=7160
		    case monthpay>85000:pertax=0.45;su=15160
		}
		per=(monthpay-5000)*pertax-su
		//年基本工资
		YearBasePay := yearbasepay(monthpay, monthnum)
		fmt.Println(YearBasePay)
		//养老税后
		After_tax_endowment := tax_endowment(monthpay, ratio_EN/100)
		tax_Eendowment := tax_endowment(monthpay, Eratio_EN/100)
		//生育税后
		After_tax_maternity := tax_maternity(monthpay, ratio_MA/100)
		tax_Ematernity := tax_maternity(monthpay, Eratio_MA/100)
		//失业税后
		After_tax_unemployment := tax_unemployment(monthpay, ratio_UN/100)
		tax_Eunemployment := tax_unemployment(monthpay, Eratio_UN/100)
		//医疗税后
		After_tax_medical := tax_medical(monthpay, ratio_ME/100)
		tax_Emedical := tax_medical(monthpay, Eratio_ME/100)
		//工伤税后
		After_tax_employment_injury := tax_employment_injury(monthpay, ratio_EM/100)
		tax_Eemployment_injury := tax_employment_injury(monthpay, Eratio_EM/100)
		//公积金税后
		After_tax_accumulation_fund := tax_accumulation_fund(monthpay, ratio_AF/100)
		tax_Eaccumulation_fund := tax_accumulation_fund(monthpay, Eratio_AF/100)
		//税后
		After_taxmonth := After_tax_month(monthpay, After_tax_accumulation_fund, After_tax_endowment, After_tax_maternity, After_tax_unemployment, After_tax_medical, After_tax_employment_injury,per)
		//After_tax_year:=YearBasePay-(After_tax_accumulation_fund+After_tax_endowment+After_tax_maternity+After_tax_unemployment+After_tax_medical+After_tax_employment_injury)

		E_tax:=Etax(tax_Eaccumulation_fund, tax_Eendowment, tax_Ematernity, tax_Eunemployment, tax_Emedical, tax_Eemployment_injury)
		fmt.Println(After_taxmonth)
		// 将内容反馈给客户端
		fmt.Fprintf(w, "月基本工资 %f，你输入的内容是：%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,%f,月税后工资：%f,企业缴纳金额：%f",
			monthpay, ratio_EN, ratio_ME, Eratio_ME, ratio_EM, Eratio_EM, ratio_MA, Eratio_MA, ratio_UN, Eratio_UN, ratio_AF, Eratio_AF, After_taxmonth,E_tax)
	} else {
		// 如果不是通过 POST 提交的数据，则将页面重定向到主页
		fmt.Fprintf(w, redirectHTML)
	}
}

func main() {
	//ROUND(MAX((Q5-5000)*{0.03;0.1;0.2;0.25;0.3;0.35;0.45}-{0;210;1410;2660;4410;7160;15160},0),2)

	http.HandleFunc("/", index)              // 设置访问的路由
	http.HandleFunc("/page", page)           // 设置访问的路由
	err := http.ListenAndServe(":8090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func yearbasepay(monthbasepay float64, monthnum float64) float64 {
	YaerBasepay := monthbasepay * monthnum
	return YaerBasepay
}
func tax_endowment(monthpay, ratio_endowment_insurance float64) float64 {
	res := monthpay * ratio_endowment_insurance
	return res
}
func tax_medical(monthpay, ratio_medical_insurance float64) float64 {
	res := monthpay * ratio_medical_insurance
	return res
}
func tax_unemployment(monthpay, ratio_unemployment_insurance float64) float64 {
	res := monthpay * ratio_unemployment_insurance
	return res
}
func tax_maternity(monthpay, ratio_maternity_insurance float64) float64 {
	res := monthpay * ratio_maternity_insurance
	return res
}
func tax_accumulation_fund(monthpay, ratio_accumulation_fund float64) float64 {
	res := monthpay * ratio_accumulation_fund
	return res
}
func tax_employment_injury(monthpay, ratio_employment_injury float64) float64 {
	res := monthpay * ratio_employment_injury
	return res
}
func After_tax_month(monthpay, tax_accumulation_fund, tax_endowment, tax_maternity, tax_unemployment, tax_medical, tax_employment_injury,per float64) float64 {
	res := monthpay - (tax_accumulation_fund + tax_endowment + tax_maternity + tax_unemployment + tax_medical + tax_employment_injury)-per
	return res
}
func Etax(tax_Eaccumulation_fund, tax_Eendowment, tax_Ematernity, tax_Eunemployment, tax_Emedical, tax_Eemployment_injury float64) float64 {
	res := tax_Eaccumulation_fund + tax_Eendowment + tax_Ematernity + tax_Eunemployment + tax_Emedical + tax_Eemployment_injury
	return res
}
