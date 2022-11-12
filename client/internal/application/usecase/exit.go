package usecase

import (
	"context"
	"fmt"
	"os"
)

// Exit stop the application.
func Exit(
	ctx context.Context,
) (string, error) {
	fmt.Println("Bye!")
	os.Exit(0)

	return "", nil
}
