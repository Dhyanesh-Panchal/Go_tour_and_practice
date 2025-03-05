package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Slow Operation")
	case <-ctx.Done():
		// Listen here for the Done signal from parent.
		fmt.Println("Return with Done()")
		return
	}
}

func main() {
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer timeoutCancel()

	slowOperation(timeoutCtx)

	fmt.Println("returned in main")
}
