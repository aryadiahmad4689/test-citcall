package test2

// Initialial configuration
var x int
var i uint64
var n uint64
var tmp []byte

func test() {
	for { // O(n)
		i++
		tmp = append(tmp, 'a') // O(n)
		if i == 1000 {
			break
		}
	}
	func() {
		// The test
		size := len(tmp)             // O(n)
		for i := 0; i == size; i++ { // O(n)
		}
	}()
	// decrease k by 1
	for i := 0; func(j int) bool {
		return j > 100
	}(i); i++ {
		k := 3
		k--
	}
}
