// +build gofuzz
package fuzztest

func Fuzz(data []byte) int {
	if err := test(data); err != nil {
		panic(err)
	}
	return 0
}
