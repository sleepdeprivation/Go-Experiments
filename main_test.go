package main
import "testing"


/*
	testFibs cross-validates the fib funcs
*/
func TestFibs(t *testing.T){
	var show [3]int;
	passed := true;
	for i := 0; i < 20 && passed; i++{
		show[0],show[1],show[2] = fib(i), fibR(i), fibArr(i);
		passed = passed && show[0] == show[1] && show[1] == show[2];
		if(!passed){
      t.Error("failure on fib ", i, show[0],show[1],show[2]);
		}
	}
}

/*
	TestRotation tests slice rotation
*/
func TestRotation(t *testing.T){
		rotArr := []int {1,2,3,4,5,6,7,8};
		orgArr := []int {1,2,3,4,5,6,7,8};
    length := len(rotArr);
		rotateSlice(rotArr);
		for i := 0; i < length; i++{
			if(rotArr[i%length] != orgArr[(i+1)%length]){
				t.Error(i, rotArr[i%length], orgArr[(i+1)%length])
			}
		}
}

/*
  TestRotationEasy tests a simpler slice rotation (perhaps slower)
*/
func TestRotationEasy(t *testing.T){
  rotArr := []int {1,2,3,4,5,6,7,8};
  orgArr := []int {1,2,3,4,5,6,7,8};
  length := len(rotArr);
  rotArr = rotateSliceEasy(rotArr);
  for i := 0; i < length; i++{
    if(rotArr[i%length] != orgArr[(i+1)%length]){
      t.Error(i, rotArr[i%length], orgArr[(i+1)%length])
    }
  }
}
