package pkg

import "fmt"
import "unicode"
// BEGIN (write your solution here)

func IsASCII(s string) bool {
    fmt.Println("String:", s)
    rs := []rune(s)
    for i := 0; i < len(rs); i++ {
        if rs[i] > unicode.MaxASCII { 
            return false
        }
    }
	return true
}

// END
