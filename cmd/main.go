package main

import (
	"challenge/internal/handler"
	"challenge/pkg"
	"fmt"
)

func main() {
	output := handler.NewStockAccountHandlerImpl().ProccessOperationsList(pkg.ReadLines())
	fmt.Println(output)
}
