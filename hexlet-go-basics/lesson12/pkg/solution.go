package pkg

const (
	OK = iota
	CANCELLED
	UNKNOWN
)

// BEGIN (write your solution here)
func ErrorMessageToCode(msg string) int {
	switch msg {
	case "OK":
		return 0
	case "CANCELLED":
		return 1
    }
	return 2
}

// END
