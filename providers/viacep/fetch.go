package viacep

import (
	"challengertwo/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseViacep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func Fetch(ctx context.Context, cep string, ch chan<- model.Address) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var raw ResponseViacep
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return
	}
	addr := model.Address{
		CEP:         raw.Cep,
		Logradouro:  raw.Logradouro,
		Bairro:      raw.Bairro,
		Localidade:  raw.Localidade,
		UF:          raw.Uf,
		APIProvider: "ViaCEP",
	}

	select {
	case ch <- addr:
	case <-ctx.Done():
	}
}
