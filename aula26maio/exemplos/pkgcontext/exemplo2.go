package main

import (
	"context"
	"fmt"
)


func saudacoes(ctx context.Context) {
	fmt.Println(ctx.Value("Saudações"))
}

func saudacoesWrapper(ctx context.Context) {
	saudacoes(ctx)
}
func main()  {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "Saudações", "Olá")
	saudacoesWrapper(ctx)
}

