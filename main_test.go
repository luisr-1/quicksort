package main

import(
	"testing"
	"fmt"
	"math/rand"
)

func CheckOrder(arr []int) error {
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] > arr[i + 1] {
			return fmt.Errorf("O item na posicao %d contem: %d que é maior que %d na posição %d", i, arr[i], i + 1, arr[i + 1])
		}
	}
	return nil
} 

func TestQuicksort(t *testing.T) {
	arr := make([]int, 5)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(20)
	}

	Quicksort(arr, 0, len(arr) - 1)
	err := CheckOrder(arr)
	if err != nil {
		t.Error(err)
	}
}

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)

	if a != 2 && b != 1 {
		t.Errorf("Ocorreu um erro na troca, a=%d e b=%d", a, b)
	}
}