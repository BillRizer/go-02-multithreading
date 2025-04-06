package main

import (
	"context"
	"fmt"
	"time"

	"challengertwo/brasilapi"
	"challengertwo/model"
	"challengertwo/viacep"
)

func main() {
	cep := "01153000"
	MAX_TIME_REQUEST := 1 * time.Second

	ch := make(chan model.Address, 1)
	ctx, cancel := context.WithTimeout(context.Background(), MAX_TIME_REQUEST)
	defer cancel()

	go brasilapi.Fetch(ctx, cep, ch)
	go viacep.Fetch(ctx, cep, ch)

	select {
	case addr := <-ch:
		fmt.Printf("Resposta obtida pelo provedor (%s):\n", addr.APIProvider)
		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nEstado: %s\n",
			addr.CEP, addr.Logradouro, addr.Bairro, addr.Localidade, addr.UF)
	case <-ctx.Done():
		fmt.Println("Nenhum provider respondeu")
	}
}
