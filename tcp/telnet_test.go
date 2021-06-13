package tcp

import (
	"fmt"
	"testing"
)

func TestTelnet(t *testing.T) {
	ok := Telnet("127.0.0.1", 65528)
	if !ok {
		fmt.Println("不通")
	} else {
		fmt.Println("通")
	}

	ok = TelnetHost("127.0.0.1:65528")
	if !ok {
		fmt.Println("不通")
	} else {
		fmt.Println("通")
	}
}
