package fuzztest

import "fmt"

func test(data []byte) error {
	if len(data) == 8 {
		if data[0] == 0x8 && data[1] == 0xb &&
			data[2] == 0xa && data[3] == 0xd &&
			data[4] == 0xf && data[5] == 0x0 &&
			data[6] == 0x0 && data[7] == 0xd {
			return fmt.Errorf("bad data")
		}
	}
	return nil
}
