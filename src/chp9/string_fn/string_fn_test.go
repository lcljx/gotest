package strng_fn

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	p := strconv.Atoi("10")
	t.Log(p + 5)
	//类型严格不能使用+进行转换
	// p := "o"
	// t.Log(p + 1)
}
