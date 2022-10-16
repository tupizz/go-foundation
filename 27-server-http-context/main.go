package main

import (
	"log"
	"net/http"
	"time"
)

// temos 5 s para processar a req e vamos retornar para o usuario que a req foi processada
// se passar dos 5 segundos deve-se cancelar a requisicao e vamos enviar para o client que a requisicao foi cancelada
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// use cases:
// processamentos pesados, calculos demorados no servidor, se o cliente cancelar não precisamos
// continuar com o que estamos fazendo
func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request iniciada")
	defer log.Println("request finalizada")

	select {
	// channel que envia dados depois de 5 segundos
	case <-time.After(time.Second * 5):
		log.Println("request processada com sucesso")
		w.Write([]byte("request processada com sucesso"))

	// se o ctx do http request é finalizado de alguma forma
	case <-ctx.Done():
		log.Println("request cancelada pelo cliente")
	}
}
