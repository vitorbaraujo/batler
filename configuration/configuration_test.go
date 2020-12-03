package configuration_test

import (
	"strings"
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"github.com/vitorbaraujo/batler/configuration"
)

func TestInvalidConfiguration(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, projectPath string
		fileContent       string
		wantErr           string
	}{
		{
			name:    "empty",
			wantErr: "configuration is not valid",
		},
		{
			name: "invalidYAML",
			fileContent: `workspace: MyWorkspace.xcworkspace
				scheme:: TestScheme`, // extra tabs
			wantErr: "could not parse yaml file",
		},
		{
			name: "no_buildDir",
			fileContent: `workspace: MyWorkspace.xcworkspace
scheme: TestScheme
xcode_version: 11.6`,
			wantErr: "missing build_dir",
		},
		{
			name: "no_scheme",
			fileContent: `workspace: MyWorkspace.xcworkspace
build_dir: Build
xcode_version: 11.6`,
			wantErr: "missing scheme",
		},
		{
			name: "no_workspace",
			fileContent: `scheme: TestScheme
build_dir: Build
xcode_version: 11.6`,
			wantErr: "missing workspace",
		},
		{
			name: "both_xcodeVersion_xcodeDeveloperDir",
			fileContent: `workspace: MyWorkspace.xcworkspace
scheme: TestScheme
build_dir: Build
xcode_version: 12.0
xcode_developer_dir: my/xcode/dir`,
			wantErr: "cannot set xcode_version and xcode_developer_dir at the same time",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := configuration.CreateConfiguration([]byte(test.fileContent), test.projectPath)

			if err == nil {
				t.Errorf("CreateConfiguration should have returned error")
			}

			if !strings.Contains(err.Error(), test.wantErr) {
				t.Errorf("CreateConfiguration returned err = %v, want %q", err, test.wantErr)
			}
		})
	}
}

func TestCreateConfiguration(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name, projectPath, fileContent string
		want                           *configuration.Configuration
	}{
		{
			name:        "OK_xcodeVersion",
			projectPath: "my/project/path",
			fileContent: `workspace: MyWorkspace.xcworkspace
scheme: TestScheme
build_dir: Build
xcode_version: 11.6`,
			want: &configuration.Configuration{
				BuildDir:  "my/project/path/Build",
				Scheme:    "TestScheme",
				XcodePath: "/Applications/Xcode-11.6.app/Contents/Developer",
				Workspace: "my/project/path/MyWorkspace.xcworkspace",
			},
		},
		{
			name:        "OK_xcodeDeveloperDir",
			projectPath: "my/project/path",
			fileContent: `workspace: MyWorkspace.xcworkspace
scheme: TestScheme
build_dir: Build
xcode_developer_dir: my/xcode/developer/dir`,
			want: &configuration.Configuration{
				BuildDir:  "my/project/path/Build",
				Scheme:    "TestScheme",
				XcodePath: "my/xcode/developer/dir",
				Workspace: "my/project/path/MyWorkspace.xcworkspace",
			},
		},
		// TODO (vitor.araujo): test default xcode path
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got, err := configuration.CreateConfiguration([]byte(test.fileContent), test.projectPath)
			if err != nil {
				t.Errorf("CreateConfiguration returned err %v", err)
			}

			if diff := pretty.CycleTracker.Compare(got, test.want); diff != "" {
				t.Errorf("post- diff: (-got +want)\n%v", diff)
			}
		})
	}
}
