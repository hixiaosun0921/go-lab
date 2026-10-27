package main
import(
	"fmt"
	"strconv"
	"strings"
)
func main(){
	str := "abcd"
	fmt.Println("str.lenth=",len(str))
	str1 := "abcd你"//中文字符占三个字节
	fmt.Println("str1.lenth=",len(str1))

	r := []rune(str1)//转换
	for i := 0; i < len(r); i++{
		fmt.Printf("字符=%c\n",r[i])
	}

	//字符串转整数
	n,err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误",err)
	}else{
		fmt.Println("转换结果是",n)
	}

	//整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T",str,str)

	//字符串转byte
	var bytes = []byte("hello hhh")
	fmt.Printf("bytes=%v\n",bytes)//输出ASCII码

	//byte转字符串
	str = string([]byte{97,98,99})
	fmt.Printf("str=%v\n",str)

	//10进制转其他进制,返回字符串
	str = strconv.FormatInt(123,2)//转2进制
	fmt.Printf("123对应的二进制是%v\n",str)
	str = strconv.FormatInt(123,8)//转8进制
	fmt.Printf("123对应的八进制是%v\n",str)

	//是否包含子串
	b := strings.Contains("seafoo","foo")//返回true ||false
	fmt.Printf("b = %v\n",b)

	//有几个子串
	num := strings.Count("ceheese","e")
	fmt.Printf("num=%d\n",num)

	//不区分大小写的比较(==区分大小写)
	b = strings.EqualFold("abc","ABC") 
	fmt.Printf("b = %v\n",b)

	//子串的第一个位置,没有则返回-1
	index := strings.Index("HHH_abc","abc")
	fmt.Println("位置是",index)

	//子串的最后位置
	index = strings.LastIndex("go go go hhh go","go")
	fmt.Println("位置是",index)

	//查找替换
	str = "go go hello go"
	str = strings.Replace(str,"go","hhh",2)
	fmt.Println("str是",str)

	//拆分成切片
	strArr := strings.Split("hello,wrold,hhh,abc",",")
	fmt.Printf("strArr是%v\n",strArr)

	//大小写转换
	str = "AbcD Efg"
	str = strings.ToLower(str)//换成小写
	fmt.Println("str是",str)
	str = strings.ToUpper(str)//换成大写
	fmt.Println("str是",str)

	//去除字符串左右空格
	str = strings.TrimSpace("  aoeiuv  hh  ")
	fmt.Println("str是",str)
	//去掉指定字符
	//
}