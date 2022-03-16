package genericorderedmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewMapElement(t *testing.T) {

}

func TestNewMap(t *testing.T) {}

func TestSet(t *testing.T) {
	type args struct {
		key   int
		value string
	}
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
		args             args
	}{
		{
			name:             "Set a new value",
			preDatasetKeys:   []int{2, 3},
			preDatasetValues: []string{"haru", "hime"},
			args: args{
				key:   1,
				value: "string",
			},
		},
		{
			name:             "Update the value",
			preDatasetKeys:   []int{1, 2, 3},
			preDatasetValues: []string{"daifuku", "haru", "hime"},
			args: args{
				key:   1,
				value: "string",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			m.Set(tt.args.key, tt.args.value)

			v, ok := m.mp[tt.args.key]
			if !ok {
				t.Errorf("Map.Set() did not set value")
			}

			want := &mapElement[int, string]{key: tt.args.key, value: tt.args.value}
			opt := cmp.AllowUnexported(mapElement[int, string]{})
			if diff := cmp.Diff(v.Value, want, opt); diff != "" {
				t.Errorf("Map.Set() set invalid value:\n%s", diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type want struct {
		value string
		ok    bool
	}
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
		arg              int
		want             want
	}{
		{
			name:             "Exist",
			preDatasetKeys:   []int{1, 2, 3},
			preDatasetValues: []string{"daifuku", "haru", "hime"},
			arg:              2,
			want:             want{value: "haru", ok: true},
		},
		{
			name:             "Do not exist",
			preDatasetKeys:   []int{1, 2, 3},
			preDatasetValues: []string{"daifuku", "haru", "hime"},
			arg:              100,
			want:             want{value: "", ok: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			v, ok := m.Get(tt.arg)

			got := want{value: v, ok: ok}
			opt := cmp.AllowUnexported(want{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("Map.Get() tests failed:\n%s", diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
		arg              int
		want             bool
	}{
		{
			name:             "Exist",
			preDatasetKeys:   []int{1, 2, 3},
			preDatasetValues: []string{"daifuku", "haru", "hime"},
			arg:              1,
			want:             true,
		},
		{
			name:             "Do not exist",
			preDatasetKeys:   []int{1, 2, 3},
			preDatasetValues: []string{"daifuku", "haru", "hime"},
			arg:              5,
			want:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			if got := m.Delete(tt.arg); got != tt.want {
				t.Errorf("Map.Delete() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
		want             int
	}{
		{
			name:             "Normal",
			preDatasetKeys:   []int{1, 2, 3},
			preDatasetValues: []string{"daifuku", "haru", "hime"},
			want:             3,
		},
		{
			name:             "Normal",
			preDatasetKeys:   []int{1, 2, 3, 4, 5, 6, 7},
			preDatasetValues: []string{"go", "rust", "c++", "python", "java", "dart", "lisp"},
			want:             7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			if got := m.Len(); got != tt.want {
				t.Errorf("Map.Len() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func setup(t *testing.T, keys []int, values []string) *Map[int, string] {
	if len(keys) != len(values) {
		t.Fatal("Length of the pre-dataset is not equal")
	}
	m := NewMap[int, string]()
	for i := range keys {
		ele := &mapElement[int, string]{key: keys[i], value: values[i]}
		e := m.l.PushBack(ele)
		m.mp[keys[i]] = e
	}

	return m
}
