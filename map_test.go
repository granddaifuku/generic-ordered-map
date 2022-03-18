package genericorderedmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNewMapElement(t *testing.T) {

}

func TestNewMap(t *testing.T) {}

func TestMap_Set(t *testing.T) {
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

func TestMap_Get(t *testing.T) {
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

func TestMap_Delete(t *testing.T) {
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

func TestMap_Len(t *testing.T) {
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

func TestMap_Keys(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
	}{
		{
			name:             "Normal",
			preDatasetKeys:   []int{53, 37, 47, 2357, 1259, 2},
			preDatasetValues: []string{"daifuku", "haru", "hime", "grand", "1998", "grand"},
		},
		{
			name:             "Normal",
			preDatasetKeys:   []int{},
			preDatasetValues: []string{},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			got := m.Keys()
			if diff := cmp.Diff(got, tt.preDatasetKeys); diff != "" {
				t.Errorf("Map.Keys() tests failed (-got +want):\n%s", diff)
			}
		})
	}
}

func TestMap_Values(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
	}{
		{
			name:             "Normal",
			preDatasetKeys:   []int{53, 37, 47, 2357, 1259, 2},
			preDatasetValues: []string{"daifuku", "haru", "hime", "grand", "1998", "grand"},
		},
		{
			name:             "Normal",
			preDatasetKeys:   []int{},
			preDatasetValues: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			got := m.Values()
			if diff := cmp.Diff(got, tt.preDatasetValues); diff != "" {
				t.Errorf("Map.Values() tests failed (-got +want):\n%s", diff)
			}
		})
	}
}

func TestMap_Entries(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
	}{
		{
			name:             "Normal",
			preDatasetKeys:   []int{53, 37, 47, 2357, 1259, 2},
			preDatasetValues: []string{"daifuku", "haru", "hime", "grand", "1998", "grand"},
		},
		{
			name:             "Normal",
			preDatasetKeys:   []int{},
			preDatasetValues: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			want := make([]Entry[int, string], len(tt.preDatasetKeys))
			for i := range tt.preDatasetKeys {
				want[i] = Entry[int, string]{Key: tt.preDatasetKeys[i], Value: tt.preDatasetValues[i]}
			}

			got := m.Entries()
			if diff := cmp.Diff(got, want); diff != "" {
				t.Errorf("Map.Entries() tests failed (-got +want):\n%s", diff)
			}
		})
	}
}

func TestMap_FromEntris(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
		arg              []Entry[int, string]
	}{
		{
			name:             "Normal",
			preDatasetKeys:   []int{53, 37, 47, 2357, 1259, 2},
			preDatasetValues: []string{"daifuku", "haru", "hime", "grand", "1998", "grand"},
			arg: []Entry[int, string]{
				{
					Key:   15,
					Value: "April",
				},
				{
					Key:   9,
					Value: "September",
				},
			},
		},
		{
			name:             "Duplicate",
			preDatasetKeys:   []int{53, 37, 47, 2357, 1259, 2},
			preDatasetValues: []string{"daifuku", "haru", "hime", "grand", "1998", "grand"},
			arg: []Entry[int, string]{
				{
					Key:   15,
					Value: "April",
				},
				{
					Key:   9,
					Value: "September",
				},
				{
					Key:   2,
					Value: "February",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)
			m.FromEntries(tt.arg)

			for _, entry := range tt.arg {
				ele, ok := m.mp[entry.Key]
				if !ok {
					t.Errorf("Map.FromEntries() failed to insert")
				}
				value := ele.Value.(*mapElement[int, string]).value
				if value != entry.Value {
					t.Errorf("Map.FromEntries() = %v, want = %v", value, entry.Value)
				}

			}
		})
	}
}

func TestMap_Front(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
	}{
		{
			name:             "Normal",
			preDatasetKeys:   []int{53, 37, 47, 2357, 1259, 2},
			preDatasetValues: []string{"daifuku", "haru", "hime", "grand", "1998", "grand"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			got := m.Front()
			want := &Element[int, string]{
				Key:   tt.preDatasetKeys[0],
				Value: tt.preDatasetValues[0],
			}

			opt := cmpopts.IgnoreUnexported(Element[int, string]{})
			if diff := cmp.Diff(got, want, opt); diff != "" {
				t.Errorf("Map.Front() tests failed (-got +want):\n%s", diff)
			}
		})
	}
}

func TestMap_Back(t *testing.T) {
	tests := []struct {
		name             string
		preDatasetKeys   []int
		preDatasetValues []string
	}{
		{
			name:             "Normal",
			preDatasetKeys:   []int{53, 37, 47, 2357, 1259, 2},
			preDatasetValues: []string{"daifuku", "haru", "hime", "grand", "1998", "grand"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := setup(t, tt.preDatasetKeys, tt.preDatasetValues)

			got := m.Back()
			l := len(tt.preDatasetKeys)
			want := &Element[int, string]{
				Key:   tt.preDatasetKeys[l-1],
				Value: tt.preDatasetValues[l-1],
			}

			opt := cmpopts.IgnoreUnexported(Element[int, string]{})
			if diff := cmp.Diff(got, want, opt); diff != "" {
				t.Errorf("Map.Back() tests failed (-got +want):\n%s", diff)
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
