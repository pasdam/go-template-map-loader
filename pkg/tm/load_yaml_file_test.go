package tm

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadYamlFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr error
	}{
		{
			name: "Should load valid file",
			args: args{
				path: filepath.Join("testdata", "valid.yml"),
			},
			want: map[string]interface{}{
				"level1": map[string]interface{}{
					"level2": map[string]interface{}{
						"key1":  "value1",
						"key2":  "value2",
						"01234": "value3",
					},
				},
				"leaf1": "leaf value",
			},
		},
		{
			name: "Should return error if the file does not exist",
			args: args{
				path: filepath.Join("testdata", "not-existing.yml"),
			},
			wantErr: errors.New("open testdata/not-existing.yml: no such file or directory"),
		},
		{
			name: "Should return error if the file is invalid",
			args: args{
				path: filepath.Join("testdata", "invalid.yml"),
			},
			wantErr: errors.New("failed to parse yaml: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `invalid...` into map[string]interface {}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadYamlFile(tt.args.path)

			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.wantErr.Error(), err.Error())
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
