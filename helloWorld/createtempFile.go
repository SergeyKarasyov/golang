// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func main() {
// 	helloWorld := "/tmp/HelloWorld1"
// 	file, err := ioutil.TempFile("", helloWorld)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer os.Remove(file.Name())
// 	if _, err := file.Write([]byte(helloWorld)); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(file.Name())
// }
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	helloWorld := "Hello, World"
	file, err := ioutil.TempFile("", "hello_world_temp")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte(helloWorld)); err != nil {
		panic(err)
	}
	fmt.Println(file.Name())
}
