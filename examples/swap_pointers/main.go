package main
import (
	"fmt"
)
func swap(a, b *int) {
	i := *a
	*a = *b
	*b = i 
}
func main(){
	x := 10
	y := 20
	swap(&x,&y)
	fmt.Printf("x=%d,y=%d",x,y)
}