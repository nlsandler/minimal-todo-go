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

func TestTaskTagList(t *testing.T) {
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
	err := client.Tasks.Tags.List(context.TODO(), 123)
	if err != nil {
		var apierr *minimaltodo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskTagAdd(t *testing.T) {
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
	_, err := client.Tasks.Tags.Add(
		context.TODO(),
		"tag_tz4a98xxat96iws9zmbrgj3a",
		minimaltodo.TaskTagAddParams{
			TodoID: "todo_nc6bzmkmd014706rfda898to",
		},
	)
	if err != nil {
		var apierr *minimaltodo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskTagRemove(t *testing.T) {
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
	_, err := client.Tasks.Tags.Remove(
		context.TODO(),
		"tag_tz4a98xxat96iws9zmbrgj3a",
		minimaltodo.TaskTagRemoveParams{
			TodoID: "todo_nc6bzmkmd014706rfda898to",
		},
	)
	if err != nil {
		var apierr *minimaltodo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
