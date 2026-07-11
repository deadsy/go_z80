//-----------------------------------------------------------------------------
/*

Build Information

*/
//-----------------------------------------------------------------------------

package util

import (
	"fmt"
	"runtime/debug"
)

//-----------------------------------------------------------------------------

// Get build information.
func GetBuildInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	var osString, archString, hashString string

	for _, s := range info.Settings {
		switch s.Key {
		case "vcs.revision":
			hashString = s.Value
		case "GOARCH":
			archString = s.Value
		case "GOOS":
			osString = s.Value
		}
	}
	return fmt.Sprintf("GOARCH=%s GOOS=%s hash=%s", archString, osString, hashString)
}

//-----------------------------------------------------------------------------
