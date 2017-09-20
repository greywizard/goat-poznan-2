package example

// New trows data race error :)
func New() {
	m := map[string]int{}
	go func() {
		for i := 0; i < 100; i++ {
			m["test"] = 1
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			x := m["test"]
			_ = x
		}
	}()
}
