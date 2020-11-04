package loop_test

import "testing"

//数组可以比较值 但是维数和数组长度必须相等 要不然不可以比较
func TestWhile(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
