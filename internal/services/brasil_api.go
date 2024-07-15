package services

import (
	"cep_finder/internal/dtos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchCepFromBrasilApiService(cep string, ch chan interface{}) {
	resp, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var cepData dtos.BrasilApiCepResponse
	err = json.Unmarshal(body, &cepData)
	if err != nil {
		return
	}

	ch <- cepData
}
