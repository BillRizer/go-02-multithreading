package brasilapi

import (
	"challengertwo/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseBrasilApi struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func Fetch(ctx context.Context, cep string, ch chan<- model.Address) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var raw ResponseBrasilApi
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return
	}
	addr := model.Address{
		CEP:         raw.CEP,
		Logradouro:  raw.Street,
		Bairro:      raw.Neighborhood,
		Localidade:  raw.City,
		UF:          raw.State,
		APIProvider: "BrasilAPI",
	}

	select {
	case ch <- addr:
	case <-ctx.Done():
	}
}
