package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	<-sigs
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	select {
	case <-ctx.Done():
		fmt.Println("Closed")
		return
	}
}
