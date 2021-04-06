/// Stolen from: https://skia.googlesource.com/buildbot/+/eaf7deacbb41/go/human/human.go

// Package human provides human friendly display formats.
package main
import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const durationTmpl = `\s*([0-9]+)\s*([smhdw])\s*`
var durationRe = regexp.MustCompile(`^(?:` + durationTmpl + `)+$`)
var durationSubRe = regexp.MustCompile(durationTmpl)
// ParseDuration parses a human readable duration. Note that this understands
// both days and weeks, which time.ParseDuration does not support.
func ParseDuration(s string) (time.Duration, error) {
	ret := time.Duration(0)
	if !durationRe.MatchString(s) {
		return ret, fmt.Errorf("Invalid format: %s", s)
	}
	parsed := durationSubRe.FindAllStringSubmatch(s, -1)
	if len(parsed) == 0 {
		return ret, fmt.Errorf("Invalid format: %s", s)
	}
	for _, match := range parsed {
		if len(match) != 3 {
			return ret, fmt.Errorf("Invalid format: %s", s)
		}
		n, err := strconv.ParseInt(match[1], 10, 32)
		if err != nil {
			return ret, fmt.Errorf("Invalid numeric format: %s", s)
		}
		switch match[2][0] {
		case 's':
			ret += time.Duration(n) * time.Second
		case 'm':
			ret += time.Duration(n) * time.Minute
		case 'h':
			ret += time.Duration(n) * time.Hour
		case 'd':
			ret += time.Duration(n) * 24 * time.Hour
		case 'w':
			ret += time.Duration(n) * 7 * 24 * time.Hour
		}
	}
	return ret, nil
}