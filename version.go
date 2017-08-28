package version

import "fmt"

func String() string  {
	return fmt.Sprintf("%d.%d.%d",major,minor,patch)
}

func VersionNum() int {
	return version(major,minor,patch)
}

func LessThan(maj,min,pat int) bool  {
	return (VersionNum() < version(maj,min,pat))
}

func version(v_major, v_minor, v_patch int)  int {
	return (v_major << 16 | v_minor << 8 | v_patch)
}

const (
	major = 5
	minor = 6
	patch = 1
)