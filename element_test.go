package genericorderedmap

import (
	"testing"
)

func TestNext(t *testing.T) {
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

			i := 0
			for el := m.Front(); el != nil; el = el.Next() {
				if el.Key != tt.preDatasetKeys[i] || el.Value != tt.preDatasetValues[i] {
					t.Errorf("Element.Next() = [%v. %v], want = [%v. %v",
						el.Key, el.Value, tt.preDatasetKeys[i], tt.preDatasetValues[i])
				}
				i++
			}
		})
	}
}

func TestPrev(t *testing.T) {
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

			i := len(tt.preDatasetKeys) - 1
			for el := m.Back(); el != nil; el = el.Prev() {
				if el.Key != tt.preDatasetKeys[i] || el.Value != tt.preDatasetValues[i] {
					t.Errorf("Element.Back() = [%v. %v], want = [%v. %v",
						el.Key, el.Value, tt.preDatasetKeys[i], tt.preDatasetValues[i])
				}
				i--
			}
		})
	}
}
