package download

import (
	"fmt"
	"strconv"
)

func readInt(override string, value int) int {
	if override != "" {
		n, err := strconv.ParseInt(override, 10, 32)
		if err != nil {
			panic(fmt.Errorf("unable to parse %q as int", override)) //nolint:forbidigo
		}
		return int(n)
	}
	return value
}
