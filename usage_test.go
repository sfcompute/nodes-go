// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/sfc-nodes-go"
	"github.com/stainless-sdks/sfc-nodes-go/internal/testutil"
	"github.com/stainless-sdks/sfc-nodes-go/option"
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
	nodes, err := client.Nodes.List(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", nodes)
}
