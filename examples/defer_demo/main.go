package main
import(
	"fmt"
)
func sum(a int ,b int) int{
	//压入栈
	defer fmt.Println("a =",a)
	defer fmt.Println("b =",b)
	a++ //不会影响上面输出的值
	b++
	ans := a + b
	fmt.Println("ans_= ",ans)
	return ans
}
func main(){
	ans := sum(10, 20)
	fmt.Println("ans =",ans)
}