package genericorderedmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSet_Set(t *testing.T) {
	tests := []struct {
		name       string
		preDataset []int
		arg        int
	}{
		{
			name:       "Set a new value",
			preDataset: []int{1, 2, 3},
			arg:        5,
		},
		{
			name:       "Duplicate",
			preDataset: []int{1, 2, 3},
			arg:        3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := setupSet(t, tt.preDataset)

			st.Set(tt.arg)

			v, ok := st.mp.mp[tt.arg]
			if !ok {
				t.Errorf("Set.Set() did not set value")
			}

			want := &mapElement[int, struct{}]{key: tt.arg, value: struct{}{}}
			opt := cmp.AllowUnexported(mapElement[int, struct{}]{})
			if diff := cmp.Diff(v.Value, want, opt); diff != "" {
				t.Errorf("Set.Set() set invalid value:\n%s", diff)
			}
		})
	}
}

func setupSet(t *testing.T, values []int) *Set[int] {
	st := NewSet[int]()
	for _, value := range values {
		ele := &mapElement[int, struct{}]{key: value, value: struct{}{}}
		e := st.mp.l.PushBack(ele)
		st.mp.mp[value] = e
	}

	return st
}
