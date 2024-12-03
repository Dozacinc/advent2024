package m02

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	Run()
}

func TestArrayChop(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6}
	for i := range arr {
		fmt.Printf("Arrrr: %v\n", arr)
		newArr := getArrayWithoutElement(arr, i)
		fmt.Printf("* %v\n", newArr)
	}

}
