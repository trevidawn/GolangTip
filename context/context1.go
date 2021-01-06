package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second * 3)

	var executeWithCtx = func(ctx context.Context, f func() error) error {
		done := make(chan error)
		go func() {
			done <- f()
		}()

		select {
		case err := <-done:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	for range ticker.C {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

		if err := executeWithCtx(ctx, Print); err != nil {
			log.Println(err)
		}
	}

}

func Print() error {
	fmt.Printf("tick %s\n", time.Now())
	time.Sleep(time.Second * 5)
	return nil
}
