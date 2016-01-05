package secret

// Handshake decodes the given binary code into
// the appropriate handhsake.
func Handshake(i int) (result []string) {
	if i <= 0 {
		return
	}

	if i&1 == 1 {
		result = append(result, "wink")
	}

	if i&2 == 2 {
		result = append(result, "double blink")
	}

	if i&4 == 4 {
		result = append(result, "close your eyes")
	}

	if i&8 == 8 {
		result = append(result, "jump")
	}

	if i&16 == 16 {
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}

	return result
}

// go test -bench . -test.benchtime 10s -test.benchmem
// PASS
// BenchmarkHandshake-8	50000000	       339 ns/op	     112 B/op	       3 allocs/op
// ok  	~/exercism/go/secret-handshake	17.333s
