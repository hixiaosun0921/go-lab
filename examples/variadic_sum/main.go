package main
import(
	"fmt"
)
func sum(n1 int,args... int) int {
	sum := n1
	for i := 0 ;i < len(args) ;i++{
		sum += args[i];
	}
	return sum
}


func main(){
	var sum1 int
	sum1 = sum(1,4,8,3)
	fmt.Println(sum1)
}