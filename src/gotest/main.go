package main

import (
	"fmt"
	"gotest/persegi"
)

func main() {
	luas, err := persegi.HitungLuasPersegi(10)

	fmt.Println(luas, err)
}

// t.run buat penamaan testing yg dilakukan
