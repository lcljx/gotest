package type_test

import "testing"

//数组可以比较值 但是维数和数组长度必须相等 要不然不可以比较
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 4, 3, 4}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d)
}
