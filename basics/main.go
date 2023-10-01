package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		println("Home bookin cancelled. Timeout rached")
	case <-time.After(3 * time.Second):
		println("Hotel booked.")
	}
}
