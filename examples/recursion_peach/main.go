package main
import(
	"fmt"
)
func f (a int) int {
	var num int
	if(a == 10) {
		num = 1
	}else {
		num = 2*(f(a+1) + 1)//第a天是第a+1天加一的二倍
	}
	return num
}
func main(){
	var num int = f(1)
	fmt.Println("第一天数量为",num)
}