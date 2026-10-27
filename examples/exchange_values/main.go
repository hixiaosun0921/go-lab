package main
import (
	"fmt"
)
func exch (a int , b int) (int,int){
	return b,a
}
func main(){
	var a,b int
	fmt.Scanf("%d%d",&a,&b)
	a,b = exch(a,b)
	fmt.Printf("交换后a=%d,b=%d",a,b)
}