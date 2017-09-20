package benchmark

import "testing"

func testA() {
	func() {}()
	return
}

func testB() {
	defer func() {}()
	return
}

func BenchmarkA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testA()
	}
}

func BenchmarkB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testB()
	}
}

func TestSomething(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		return
	}
}
