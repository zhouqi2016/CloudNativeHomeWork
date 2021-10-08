package main

import "fmt"

func main(){
	test_slice()
}

func test_array1(){
    var arr []int
	arr1 := [3]int{1, 2, 3}
	// 可以不指定数组的元素个数
	arr2 := [...]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(arr1 == arr2)
	fmt.Println(arr1[0])
	// 数组的遍历, 使用的是如下方式，需要注意。
	for _, value:= range arr1 {
		fmt.Print(value)
	}
}

func test_array(){
	array := [3][4]int{{1, 2, 3, 4}, {6, 7, 8, 9}, {10, 11, 12, 13}}
	for i := 0; i < len(array); i++ {
		for j:=0; j< len(array[i]); j++ {
			fmt.Print(array[i][j])
		}
	}
}

func test_slice(){
	//同时声明切片的长度和容量为5
	slice1 := make([]string, 5)
	//直接给切片赋值
	slice2 := []string{"a", "b", "c"}
	slice2 = append(slice2, "d")
	fmt.Print(slice1)
    fmt.Print(slice2)
}

func test_array_add(){
    slice := make(int[],3, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)
	fmt.Println(cap(slice))
	
}
