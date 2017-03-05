package main

import "fmt"

/*
	Rotate a slice in a circular fashion
	O(n) time complexity
*/
func rotateSlice(arr []int){
	a := arr[0];
	i := 0;
	for ;i < len(arr)-1; i++{
		arr[i] = arr[i+1];
	}
	arr[i] = a;
}

func testRotation(){
		t := []int {1,2,3,4,5,6,7,8};
		rotateSlice(t);
		fmt.Println(t);
}

func rotateSliceEasy(arr []int){

}

/*
	Fibbonacci Recursive-Style
*/
func fibR(n int) int{
	if(n <= 1){
		return 1;
	}
	return fibR(n-1) + fibR(n-2);
}

/*
	Fibbonacci Iterative-Style
*/
func fib(n int) int{
	minus1, minus2,ret := 0,0,1
	for i := 0; i < n; i++{
		minus2 = minus1;
		minus1 = ret;
		ret = minus1 + minus2;
	}
	return ret;
}

/*
	Fibbonacci built on slice rotation scheme
*/
func fibArr(n int) int{
	arr := []int{0,0,1};
	for ;n > 0; n--{
		rotateSlice(arr);
		arr[2] = arr[0] + arr[1];
	}
	return arr[2];
}

/*
	Cross validate the fib funcs
*/
func testFibs() bool{
	var show [3]int;
	passed := true;
	for i := 0; i < 20 && passed; i++{
		show[0],show[1],show[2] = fib(i), fibR(i), fibArr(i);
		passed = passed && show[0] == show[1] && show[1] == show[2];
		fmt.Println(show, passed);
	}
	return passed;
}

func main() {
	//testFibs();
}
