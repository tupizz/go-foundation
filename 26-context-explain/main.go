package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	// se o timeout for maior que o tempo que agt leva para bookar o hotel conseguiremos booka-lo de fato
	ctx, cancel := context.WithTimeout(ctx, 4*time.Second)
	defer cancel()
	bookHotel(ctx)
}

// tentamos bookar um hotel e verificamos a condição de sucesso
func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("timeout reached, booked cancelled")
		return
	case <-time.After(3 * time.Second):
		fmt.Println("hotel booked")
		return
	}
}
