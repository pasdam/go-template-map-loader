package tm

import (
	"reflect"
	"testing"
)

func TestWithPrefix(t *testing.T) {
	type args struct {
		prefix string
		data   map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Should ignore empty prefix",
			args: args{
				prefix: "",
				data: map[string]interface{}{
					"prev0_1": "prev0_1-value",
					"prev0_2": "prev0_2-value",
				},
			},
			want: map[string]interface{}{
				"prev0_1": "prev0_1-value",
				"prev0_2": "prev0_2-value",
			},
		},
		{
			name: "Should add 1 prefix",
			args: args{
				prefix: "pref1_1",
				data: map[string]interface{}{
					"prev1_1": "prev1_1-value",
					"prev1_2": "prev1_2-value",
				},
			},
			want: map[string]interface{}{
				"pref1_1": map[string]interface{}{
					"prev1_1": "prev1_1-value",
					"prev1_2": "prev1_2-value",
				},
			},
		},
		{
			name: "Should add 3 prefixes",
			args: args{
				prefix: "pref3_1.pref3_2.pref3_3",
				data: map[string]interface{}{
					"prev3_1": "prev3_1-value",
					"prev3_2": "prev3_2-value",
				},
			},
			want: map[string]interface{}{
				"pref3_1": map[string]interface{}{
					"pref3_2": map[string]interface{}{
						"pref3_3": map[string]interface{}{
							"prev3_1": "prev3_1-value",
							"prev3_2": "prev3_2-value",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithPrefix(tt.args.prefix, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
