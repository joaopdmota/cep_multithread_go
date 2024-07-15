package handlers

import (
	"cep_finder/internal/dtos"
	"cep_finder/internal/services"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

type CepHandler struct{}

func NewCepHandler() *CepHandler {
	return &CepHandler{}
}

func (c *CepHandler) FetchCep(w http.ResponseWriter, r *http.Request) {
	var request dtos.CepRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	isReqValid := request.Cep != "" && len(request.Cep) == 8
	if err != nil || !isReqValid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	defer waitGroup.Wait()
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go services.FetchCepFromViaCepService(request.Cep, ch1)
	go services.FetchCepFromBrasilApiService(request.Cep, ch2)

	select {
	case data := <-ch1:
		log.Printf("Received from ViaCep \n %+v\n", data)
		waitGroup.Done()
	case data := <-ch2:
		log.Printf("Received from BrasilApi \n %+v\n", data)
		waitGroup.Done()
	case <-time.After(time.Second * 1):
		log.Println("Timeout occurred")
		waitGroup.Done()
	}

}
