package main

//divisão: Partition Array int[p..r] -> subArray1[p..q-1] && subArray2[q+1..r]
//ex: lista = [4,5,3,6,7,0] -> q == 3 -> subArray1 == [4, 5] && subArray[6, 7, 0]
//mas subArray1 <= q <= subArray2, entao ate essa condicao ser satisfeita o proximo passo se faz necessario

//conquista: Ordenar subArray1[p..q-1] && subArray2[q+1..r] por chamadas recursivas de QuickSort
//combinacao: Como os subarranjos depois das chamadas estao ordenados, nao precisa fazer o Merge, Array[p..r] fica totalmente ordenado

func Quicksort(a []int, low, high int) {
	if(low < high) {
		pivotLocation := Partition(a, low, high)
		Quicksort(a, low, pivotLocation - 1)
		Quicksort(a, pivotLocation + 1, high) 
	}
}

func Partition(a []int, low, high int) int {
	pivot := a[high]
	leftwall := low - 1

	for i := low; i <= high - 1; i++ {
		if(a[i] < pivot) {
			leftwall += 1
			Swap(&a[i], &a[leftwall])
		}
	}

	Swap(&a[high], &a[leftwall + 1])

	return leftwall + 1
}

func Swap(r, l *int) {
	*r, *l = *l, *r
}