// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package plugin

import (
	"path"
	"strings"
)

func match(name string, patterns []string) bool {
	if len(patterns) == 0 {
		return false
	}
	name = strings.ToLower(name)
	for _, pattern := range patterns {
		pattern = strings.ToLower(pattern)
		match, _ := path.Match(pattern, name)
		if match {
			return true
		}
	}
	return false
}
