package main

import "fmt"

func main() {
	data := fmt.Sprintf(`{"idusuario"}:{"%v"}`, "This")
	fmt.Print(data)
}
