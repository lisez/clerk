package version

import "fmt"

// Version ...
type Version struct {
	Number   float32
	Security int
	Suffix   string
}

func (v *Version) String() string {
	return fmt.Sprintf("%g.%d%s", v.Number, v.Security, v.Suffix)
}
