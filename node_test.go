// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package sfcnodes_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/sfcompute/nodes-go"
	"github.com/sfcompute/nodes-go/internal/testutil"
	"github.com/sfcompute/nodes-go/option"
)

func TestNodeNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Nodes.New(context.TODO(), sfcnodes.NodeNewParams{
		CreateNodesRequest: sfcnodes.CreateNodesRequestParam{
			DesiredCount:        1,
			MaxPricePerNodeHour: 1600,
			AnyZone:             sfcnodes.Bool(false),
			CloudInitUserData:   sfcnodes.String("aGVsbG8gd29ybGQ="),
			EndAt:               sfcnodes.Int(0),
			Forward443:          sfcnodes.Bool(false),
			ImageID:             sfcnodes.String("vmi_1234567890abcdef"),
			Names:               []string{"cuda-crunch"},
			NodeType:            sfcnodes.NodeTypeAutoreserved,
			StartAt:             sfcnodes.Int(1640995200),
			Zone:                sfcnodes.String("hayesvalley"),
		},
	})
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNodeListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Nodes.List(context.TODO(), sfcnodes.NodeListParams{
		ID:   []string{"string"},
		Name: []string{"string"},
		Type: sfcnodes.NodeTypeAutoreserved,
	})
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNodeDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	err := client.Nodes.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNodeExtend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Nodes.Extend(
		context.TODO(),
		"id",
		sfcnodes.NodeExtendParams{
			ExtendNodeRequest: sfcnodes.ExtendNodeRequestParam{
				DurationSeconds:     7200,
				MaxPricePerNodeHour: 1000,
			},
		},
	)
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNodeGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Nodes.Get(context.TODO(), "id")
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNodeRedeployWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Nodes.Redeploy(
		context.TODO(),
		"id",
		sfcnodes.NodeRedeployParams{
			CloudInitUserData: sfcnodes.String("aGVsbG8gd29ybGQ="),
			ImageID:           sfcnodes.String("vmi_1234567890abcdef"),
			OverrideEmpty:     sfcnodes.Bool(true),
		},
	)
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNodeRelease(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Nodes.Release(context.TODO(), "id")
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
