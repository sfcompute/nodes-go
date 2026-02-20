// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes_test

import (
	"context"
	"os"
	"testing"

	"github.com/sfcompute/nodes-go"
	"github.com/sfcompute/nodes-go/internal/testutil"
	"github.com/sfcompute/nodes-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sfcnodes.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	t.Skip("Mock server tests are disabled")
	listResponseNode, err := client.Nodes.List(context.TODO(), sfcnodes.NodeListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", listResponseNode.Data)
}
