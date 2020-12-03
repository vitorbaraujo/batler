package simctl_test

import (
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/vitorbaraujo/batler/simctl"
)

func TestParseRuntimesError(t *testing.T) {
	t.Parallel()
	invalidJSON := `{"runtimes": {},}`

	_, err := simctl.ParseRuntimesOutput([]byte(invalidJSON))
	if err == nil {
		t.Errorf("ParseRuntimesOutput should have returned error")
	}

	wantErr := "parsing runtimes output"
	if !strings.Contains(err.Error(), wantErr) {
		t.Errorf("ParseRuntimesOutput returned err = %v, want %q", err, wantErr)
	}
}

func TestParseRuntimes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		output string
		want   []*simctl.Runtime
	}{
		{
			name: "noRuntimes",
			output: `
				{
					"runtimes": []
				}
			`,
			want: []*simctl.Runtime{},
		},
		{
			name: "someDevices",
			output: `
				{
					"runtimes": [
						{
							"bundlePath": "bundle_path",
							"buildversion": "16A366",
							"runtimeRoot": "runtime_root",
							"identifier": "com.apple.CoreSimulator.SimRuntime.iOS-12-0",
							"version": "12.0",
							"isAvailable": false,
							"name": "iOS 12.0"
						},
						{
							"bundlePath": "another_bundle_path",
							"buildversion": "18B79",
							"runtimeRoot": "another_runtime_root",
							"identifier": "com.apple.CoreSimulator.SimRuntime.iOS-14-2",
							"version": "14.2",
							"isAvailable": true,
							"name": "iOS 14.2"
						}
					]
				}
			`,
			want: []*simctl.Runtime{
				{
					Name:       "iOS 12.0",
					Identifier: "com.apple.CoreSimulator.SimRuntime.iOS-12-0",
				},
				{
					Name:       "iOS 14.2",
					Identifier: "com.apple.CoreSimulator.SimRuntime.iOS-14-2",
					Available:  true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := simctl.ParseRuntimesOutput([]byte(test.output))
			if err != nil {
				t.Errorf("ParseDeviceTypesOutput returned err %v", err)
			}

			if diff := pretty.CycleTracker.Compare(got, test.want); diff != "" {
				t.Errorf("post- diff: (-got +want)\n%v", diff)
			}
		})
	}
}
