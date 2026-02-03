// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package minimaltodo_test

import (
	"context"
	"os"
	"testing"

	"github.com/nlsandler/minimal-todo-go"
	"github.com/nlsandler/minimal-todo-go/internal/testutil"
	"github.com/nlsandler/minimal-todo-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := minimaltodo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Prism tests are disabled")
	tag, err := client.Tags.Get(context.TODO(), "id")
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", tag.ID)
}
