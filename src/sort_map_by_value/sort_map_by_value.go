package sort_map_by_value

import (
	"sort"
	"domains"
)



func SortByValue(m map[string]int) []domains.Record {

	var records []domains.Record

	var vals []int

	for _, val := range m {
		vals = append(vals, val)
	}

	valSet := make(map[int]bool)

	for _, val := range vals {
		valSet[val] = true
	}

	var uniqvals []int

	for key, _ := range valSet {

		uniqvals = append(uniqvals, key)

	}
	sort.Ints(uniqvals)

	for _, kval := range uniqvals {
		for key, val := range m {

			if val == kval {

				records = append(records, domains.Record{key, val})

			}

		}

	}

	return records

}
