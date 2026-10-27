package main
import(
	"fmt"
)
func f (a int) int {
	if(a == 1){
		return 3
	} else {
		return 2 * f(a - 1) + 1
	}

}
func main(){
	fmt.Println("f(1)=",f(1))
	fmt.Println("f(2)=",f(2))
	fmt.Println("f(3)=",f(3))
}