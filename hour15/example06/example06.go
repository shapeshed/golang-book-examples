package salt

import (
	"fmt"
	"io/ioutil"
)

/*
var Salt []byte

func init() {
	var err error
	Salt, err = ioutil.ReadFile("salt.txt")
	if err != nil {
		fmt.Print(err)
	}
}
*/

func SaltSecret(secret string) string {
	b, err := ioutil.ReadFile("salt.txt")
	if err != nil {
		fmt.Print(err)
	}
	return secret + string(b)
}

/*
func SaltSecret(secret string) string {
	return secret + string(Salt)
}*/
