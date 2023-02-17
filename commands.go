package songci

import (
	"strings"
)

type commands []string

func (c *commands) append(command ...string) {
	*c = append(*c, command...)
}

func (c *commands) String() string {
	return strings.Join(*c, space)
}
