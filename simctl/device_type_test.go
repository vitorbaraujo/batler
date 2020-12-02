package simctl_test

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/vitorbaraujo/batler/simctl"
)

func TestParseDeviceTypes(t *testing.T) {
	tests := []struct{
		name string
		output []byte
		want []*simctl.DeviceType
	}{
		{
			name: "noDevices",
			output: []byte(`
				{
					"devicetypes": []
				}
			`),
			want: []*simctl.DeviceType{},
		},
		{
			name: "someDevices",
			output: []byte(`
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
			`),
			want: []*simctl.DeviceType{
				{
					Name: "iPhone 4s",
					Identifier: "com.apple.CoreSimulator.SimDeviceType.iPhone-4s",
					ProductFamily: "iPhone",
				},
				{
					Name: "iPhone 5",
					Identifier: "com.apple.CoreSimulator.SimDeviceType.iPhone-5",
					ProductFamily: "iPhone",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			got, err := simctl.ParseDeviceTypesOutput(test.output)
			if err != nil {
				t.Errorf("ParseDeviceTypesOutput returned err %v", err)
			}

			if diff := pretty.CycleTracker.Compare(got, test.want); diff != "" {
			    t.Errorf("post- diff: (-got +want)\n%v", diff)
			}
		})
	}
}