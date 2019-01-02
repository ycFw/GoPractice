package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var buf []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		buf = append(buf, rand.Intn(100))
	}
	fmt.Printf("排序前: %v\n", buf)

	// BubbleSort(buf)
	// SelectionSort(buf)
	// InsertSort(buf)
	MergeSort(buf)

	fmt.Printf("排序后: %v\n", buf)
}

//冒泡排序
func BubbleSort(buf []int) ([]int, int) {
	a := len(buf)
	times := 0
	//flag标识是否有需要交换位置，若flag==false，表示不需要交换，break
	flag := false
	for i := 0; i < a-1; i++ {
		for j := 0; j < a-1-i; j++ {
			if buf[j] > buf[j+1] {
				times++
				buf[j], buf[j+1] = buf[j+1], buf[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Printf("排序次数: %d\n", times)
	return buf, times
}

//选择排序
func SelectionSort(buf []int) ([]int, int) {
	a := len(buf)
	times := 0
	for i := 0; i < a-1; i++ {
		min := i
		for j := i; j < a; j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != i {
			buf[i], buf[min] = buf[min], buf[i]
		}
	}
	fmt.Printf("排序次数: %d\n", times)
	return buf, times
}

//插入排序
func MergeSort(buf []int) []int {
	tmp := make([]int, len(buf))
	merge_sort(buf, 0, len(buf)-1, tmp)
	return buf
}

func merge_sort(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		merge_sort(a, first, middle, tmp)       //左半部分排好序
		merge_sort(a, middle+1, last, tmp)      //右半部分排好序
		mergeArray(a, first, middle, last, tmp) //合并左右部分
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	// fmt.Printf("mergeArray a: %v, first: %v, middle: %v, end: %v, tmp: %v\n",
	//     a, first, middle, end, tmp)
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
	// fmt.Printf("sort: buf: %v\n", a)
}

//希尔排序
//快速排序
//堆排序
