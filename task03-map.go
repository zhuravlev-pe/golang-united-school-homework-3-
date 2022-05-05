package homework

import "sort"

type keyValue struct {
	key   int
	value string
}

type mapValues []keyValue

func (mv mapValues) Len() int {
	return len(mv)
}

func (mv mapValues) Less(i, j int) bool {
	return mv[i].key < mv[j].key
}

func (mv mapValues) Swap(i, j int) {
	mv[i], mv[j] = mv[j], mv[i]
}

func sortMapValues(input map[int]string) (result []string) {
	l := len(input)
	mapVals := make([]keyValue, 0, l)
	for k, v := range input {
		mapVals = append(mapVals, keyValue{k, v})
	}
	sort.Sort(mapValues(mapVals))
	result = make([]string, l)
	for i, mv := range mapVals {
		result[i] = mv.value
	}
	return result
}
