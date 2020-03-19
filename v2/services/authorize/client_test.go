package authorize

import (
	"context"
	"testing"
	"time"
)

func Test_Client(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	auth := CreateClient()
	auth.DialUsingCredentialsWithContext(ctx)
}
