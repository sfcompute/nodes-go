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

func TestVmLogsWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sfcnodes.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Vms.Logs(context.TODO(), sfcnodes.VmLogsParams{
		InstanceID:              "instance_id",
		OrderBy:                 sfcnodes.VmLogsParamsOrderBySeqnumAsc,
		BeforeRealtimeTimestamp: sfcnodes.String("before_realtime_timestamp"),
		BeforeSeqnum:            sfcnodes.Int(0),
		Limit:                   sfcnodes.Int(1),
		SinceRealtimeTimestamp:  sfcnodes.String("since_realtime_timestamp"),
		SinceSeqnum:             sfcnodes.Int(0),
	})
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVmSSH(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := sfcnodes.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Vms.SSH(context.TODO(), sfcnodes.VmSSHParams{
		VmID: "vm_id",
	})
	if err != nil {
		var apierr *sfcnodes.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
