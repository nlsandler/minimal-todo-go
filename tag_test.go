// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package minimaltodo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nlsandler/minimal-todo-go"
	"github.com/nlsandler/minimal-todo-go/internal/testutil"
	"github.com/nlsandler/minimal-todo-go/option"
)

func TestTagNew(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Tags.New(context.TODO(), minimaltodo.TagNewParams{
		Label:   minimaltodo.String("Work"),
		OwnerID: "owner_id",
	})
	if err != nil {
		var apierr *minimaltodo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTagGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Tags.Get(context.TODO(), "tag_tz4a98xxat96iws9zmbrgj3a")
	if err != nil {
		var apierr *minimaltodo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTagListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Tags.List(context.TODO(), minimaltodo.TagListParams{
		Cursor: minimaltodo.String("cursor"),
		Limit:  minimaltodo.Float(100),
	})
	if err != nil {
		var apierr *minimaltodo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTagDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Tags.Delete(context.TODO(), "tag_tz4a98xxat96iws9zmbrgj3a")
	if err != nil {
		var apierr *minimaltodo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
