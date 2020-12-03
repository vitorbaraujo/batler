package simctl_test

import (
	"strings"
	"testing"

	"github.com/vitorbaraujo/batler/simctl"

	"github.com/kylelemons/godebug/pretty"
)

func TestParseDeviceTypesError(t *testing.T) {
	t.Parallel()
	invalidJSON := `{"devicetypes": {},}`

	_, err := simctl.ParseDeviceTypesOutput([]byte(invalidJSON))
	if err == nil {
		t.Errorf("ParseDeviceTypesOutput should have returned error")
	}

	wantErr := "parsing devicetypes output"
	if !strings.Contains(err.Error(), wantErr) {
		t.Errorf("ParseDeviceTypesOutput returned err = %v, want %q", err, wantErr)
	}
}

func TestParseDeviceTypes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		output string
		want   []*simctl.DeviceType
	}{
		{
			name: "noDevices",
			output: `
				{
					"devicetypes": []
				}
			`,
			want: []*simctl.DeviceType{},
		},
		{
			name: "someDevices",
			output: `
				{
					"devicetypes": [
						{
						  "minRuntimeVersion": 327680,
						  "bundlePath": "some_bundle_path",
						  "maxRuntimeVersion": 655359,
						  "name": "iPhone 4s",
						  "identifier": "com.apple.CoreSimulator.SimDeviceType.iPhone-4s",
						  "productFamily": "iPhone"
						},
						{
							"minRuntimeVersion": 393216,
							"bundlePath": "another_bundle_path",
							"maxRuntimeVersion": 720895,
							"name": "iPhone 5",
							"identifier": "com.apple.CoreSimulator.SimDeviceType.iPhone-5",
							"productFamily": "iPhone"
						}
					]
				}
			`,
			want: []*simctl.DeviceType{
				{
					Name:          "iPhone 4s",
					Identifier:    "com.apple.CoreSimulator.SimDeviceType.iPhone-4s",
					ProductFamily: "iPhone",
				},
				{
					Name:          "iPhone 5",
					Identifier:    "com.apple.CoreSimulator.SimDeviceType.iPhone-5",
					ProductFamily: "iPhone",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := simctl.ParseDeviceTypesOutput([]byte(test.output))
			if err != nil {
				t.Errorf("ParseDeviceTypesOutput returned err %v", err)
			}

			if diff := pretty.CycleTracker.Compare(got, test.want); diff != "" {
				t.Errorf("post- diff: (-got +want)\n%v", diff)
			}
		})
	}
}
