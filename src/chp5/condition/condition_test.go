package condition_test

import "testing"

func TestMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("It is not between 0 and 3")
		}

	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unkonw")
		}

	}
}

func TestLoop(t *testing.T) {
	for i := 0; i < 10000; i++ {
		t.Log(i)
	}
}
