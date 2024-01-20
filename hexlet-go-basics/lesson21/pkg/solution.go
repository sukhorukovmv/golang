package pkg

// BEGIN (write your solution here)

func ShiftASCII(s string, shift int) string {
    shiftString := ""
	for _, ch := range s {
        shiftString += string(byte(ch) + byte(shift))
	}
	return shiftString
}

// END
