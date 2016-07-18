package hid

import (
	"encoding/json"
	"fmt"

	"testing"
)

func TestHid(t *testing.T) {
	for dev := range Devices() {
		str, _ := json.Marshal(dev)
		fmt.Println(string(str))
	}
}
