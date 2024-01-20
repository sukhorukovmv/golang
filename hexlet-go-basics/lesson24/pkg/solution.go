package pkg

import "fmt"

// BEGIN (write your solution here)
func GenerateSelfStory(name string, age int, money float64) string {
    return fmt.Sprintf("Hello! My name is %[1]s. I'm %[2]d y.o. And I also have $%.[3]2f in my wallet right now.", name, age, money)
}

// END
