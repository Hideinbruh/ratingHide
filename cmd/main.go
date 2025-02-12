package main

import (
	"fmt"
	"github.com/Hideinbruh/initializing_project/internal/handlers/employer"
)

func main() {
	e := employer.New("Влаdick")

	fmt.Println(e.GetName())
}
