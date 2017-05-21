/* recursive function */
package main

import "fmt"

func main() {
	fn := func() {
		fmt.Println("function called")
	}
	fn()

}
