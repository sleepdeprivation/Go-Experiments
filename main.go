/*
	Getting to know the go language
*/
package main


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

/*
	Rotate slice using fancy syntax
	This is probably dangerous in some way
*/
func rotateSliceEasy(arr []int) []int{
	return append(arr[1:], arr[0]);
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


func main() {
}
