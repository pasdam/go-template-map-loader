package tm

import (
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	type args struct {
		maps []map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Should return empty map if input maps array is nil",
			args: args{
				maps: nil,
			},
			want: map[string]interface{}{},
		},
		{
			name: "Should return empty map if no input maps are provided",
			args: args{
				maps: []map[string]interface{}{},
			},
			want: map[string]interface{}{},
		},
		{
			name: "Should merge non overlapping maps",
			args: args{
				maps: []map[string]interface{}{
					{
						"k1": "v1",
						"k2": "v2",
					},
					{
						"k3": "v3",
					},
				},
			},
			want: map[string]interface{}{
				"k1": "v1",
				"k2": "v2",
				"k3": "v3",
			},
		},
		{
			name: "Should merge overlapping maps",
			args: args{
				maps: []map[string]interface{}{
					{
						"k1": "v1",
						"k2": "v2_first",
						"k3": map[string]interface{}{
							"k3_1": "v3_1",
							"k3_2": "v3_2_first",
						},
					},
					{
						"k2": "v2_last",
						"k3": map[string]interface{}{
							"k3_2": "v3_last",
							"k3_3": "v3_3",
						},
						"k4": "v4",
					},
				},
			},
			want: map[string]interface{}{
				"k1": "v1",
				"k2": "v2_last",
				"k3": map[string]interface{}{
					"k3_1": "v3_1",
					"k3_2": "v3_last",
					"k3_3": "v3_3",
				},
				"k4": "v4",
			},
		},
		{
			name: "Should merge maps with interface keys",
			args: args{
				maps: []map[string]interface{}{
					{
						"l1": map[interface{}]interface{}{
							1: "k1v1",
							2: "k2v1",
						},
					},
					{
						"l1": map[interface{}]interface{}{
							2: "k2v2",
							3: "k3v1",
						},
					},
				},
			},
			want: map[string]interface{}{
				"l1": map[interface{}]interface{}{
					1: "k1v1",
					2: "k2v2",
					3: "k3v1",
				},
			},
		},
		{
			name: "Should merge maps with int keys",
			args: args{
				maps: []map[string]interface{}{
					{
						"l1": map[string]interface{}{
							"l2": map[int]interface{}{
								1: "k1v1",
								2: "k2v1",
							},
						},
					},
					{
						"l1": map[string]interface{}{
							"l2": map[int]interface{}{
								2: "k2v2",
								3: "k3v1",
							},
						},
					},
				},
			},
			want: map[string]interface{}{
				"l1": map[string]interface{}{
					"l2": map[int]interface{}{
						1: "k1v1",
						2: "k2v2",
						3: "k3v1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeMaps(tt.args.maps...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
