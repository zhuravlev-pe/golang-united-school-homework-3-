package homework

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

var tiny = map[int]string{2: "a", 0: "b", 1: "c"}
var tinyResult = []string{"b", "c", "a"}
var test1000 = getTestMap(1000, "qqq")
var test100000 = getTestMap(100000, "qqq")

var blackhole []string

func getTestMap(length int, payload string) map[int]string {
	result := map[int]string{}
	for i := 0; i < length; i++ {
		result[i] = payload
	}
	return result
}

//func TestMain(m *testing.M) {
//	test1000 = getTestMap(1000, "qqq")
//	exitVal := m.Run()
//	os.Exit(exitVal)
//}

func TestSortMapValuesTiny(t *testing.T) {
	result := sortMapValues(tiny)
	diff := cmp.Diff(result, tinyResult)
	if diff != "" {
		t.Errorf("Unexpected result: %s", diff)
	}
}

func TestSortMapValuesAltTiny(t *testing.T) {
	result := sortMapValuesAlt(tiny)
	diff := cmp.Diff(result, tinyResult)
	if diff != "" {
		t.Errorf("Unexpected result: %s", diff)
	}
}

func BenchmarkSortMapValues_Tiny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackhole = sortMapValues(tiny)
	}
}

func BenchmarkSortMapValuesAlt_Tiny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackhole = sortMapValuesAlt(tiny)
	}
}

func BenchmarkSortMapValues_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackhole = sortMapValues(test1000)
	}
}

func BenchmarkSortMapValuesAlt_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackhole = sortMapValuesAlt(test1000)
	}
}

func BenchmarkSortMapValues_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackhole = sortMapValues(test100000)
	}
}

func BenchmarkSortMapValuesAlt_100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		blackhole = sortMapValuesAlt(test100000)
	}
}
