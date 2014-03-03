package liquor

import (
	"sort"
)

type Licenses []*License

// Implement the sort.Interface for sorting
func (a Licenses) Len() int {
	return len(a)
}

func (a Licenses) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type ByUniqueId struct {
	Licenses
}

// Sort by the unique Id string
func (a ByUniqueId) Less(i, j int) bool {
	return a.Licenses[i].UniqueId > a.Licenses[j].UniqueId
}

func (a Licenses) Sort() {
	sort.Sort(ByUniqueId{a})
}
