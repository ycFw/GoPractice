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
	// MergeSort(buf)
	// HeapSort(buf)
	// QuickSort(buf)
	HillSorting(buf)

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
func HillSorting(buf []int) []int {
	times := 0
	tmp := 0
	length := len(buf)
	incre := length
	// fmt.Println("buf: ", buf)
	for {
		incre /= 2
		for k := 0; k < incre; k++ { //根据增量分为若干子序列
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					// fmt.Println("j: ", j, " data: ", buf[j], " j-incre: ", j-incre, " data: ", buf[j-incre])
					times++
					if buf[j] < buf[j-incre] {
						tmp = buf[j-incre]
						buf[j-incre] = buf[j]
						buf[j] = tmp
					} else {
						break
					}
				}
				// fmt.Println("middle: ", buf)
			}
			// fmt.Println("outer: ", buf)
		}
		// fmt.Println("outer outer: ", buf, " incre: ", incre)

		if incre == 1 {
			break
		}
	}
	// fmt.Println("after: ", buf)
	fmt.Printf("排序次数: %d\n", times)
	return buf
}

//快速排序
func QuickSort(buf []int) []int {
	quick(buf, 0, len(buf)-1)
	return buf
}

func quick(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[l] //选择第一个数为key
	for i < j {
		for i < j && a[j] > key { //从右向左找第一个小于key的值
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && a[i] < key { //从左向右找第一个大于key的值
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	//i == j
	a[i] = key
	quick(a, l, i-1)
	quick(a, i+1, r)
}

//堆排序
func HeapSort(buf []int) []int {
	temp, n := 0, len(buf)

	for i := (n - 1) / 2; i >= 0; i-- {
		MinHeapFixdown(buf, i, n)
	}

	for i := n - 1; i > 0; i-- {
		temp = buf[0]
		buf[0] = buf[i]
		buf[i] = temp
		MinHeapFixdown(buf, 0, i)
	}
	return buf
}

func MinHeapFixdown(a []int, i, n int) {
	j, temp := 2*i+1, 0
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}

		if a[i] <= a[j] {
			break
		}

		temp = a[i]
		a[i] = a[j]
		a[j] = temp

		i = j
		j = 2*i + 1
	}
}
