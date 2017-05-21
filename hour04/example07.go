/* recursive function */
package main

import "fmt"

func anotherFunction(f func() string) string {
	return f()
}

func main() {
	fn := func() string {
		return "function called"
	}
	fmt.Println(anotherFunction(fn))

}
