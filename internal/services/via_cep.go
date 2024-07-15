package services

import (
	"cep_finder/internal/dtos"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchCepFromViaCepService(cep string, ch chan interface{}) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var cepData dtos.ViaCepCepResponse
	err = json.Unmarshal(body, &cepData)
	if err != nil {
		return
	}

	ch <- cepData
}
