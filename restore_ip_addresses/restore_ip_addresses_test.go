package restore_ip_addresses

import (
	"fmt"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	var text string
	var res []string

	text = "25525511135"
	res = restoreIpAddresses(text)

	fmt.Print("255.255.11.135,255.255.111.35 ", res)

}
