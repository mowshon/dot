package dot

import "strings"

// preparePath merges the path that has already been traversed with that to be traversed
// e.g. "old.path" + ["remaining", "path"] => "old.path.remaining.path"
func preparePath(previousPath string, parts []string) string {
	if previousPath != "" {
		parts = append([]string{previousPath}, parts...)
	}

	return strings.Join(parts, ".")
}
