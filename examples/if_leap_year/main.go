package main
import (
	"fmt"
)
func main(){
	var year int 
	fmt.Println("输入年份")
	fmt.Scanln(&year)
	if (year%4==0&&year%100!=0) || year%4200==0 {
		fmt.Println("闰年")
	} else{
		fmt.Println("非闰年")
	}
}