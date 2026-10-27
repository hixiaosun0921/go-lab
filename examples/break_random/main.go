package main
import (
	"fmt"
	"math/rand"
)
func main(){
	count := 0
	for {
	n := rand.Intn(100) + 1
	fmt.Println(n)
	count ++
	if n==99{
		break
	}
	}
	fmt.Println("次数为",count)
}