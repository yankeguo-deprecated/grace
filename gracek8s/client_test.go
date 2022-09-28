package gracek8s

import (
	"os"
	"strconv"
	"strings"
)

func ShouldSkip() bool {
	ok, _ := strconv.ParseBool(strings.TrimSpace(os.Getenv("SHOULD_TEST_GRACEK8S")))
	return !ok
}
