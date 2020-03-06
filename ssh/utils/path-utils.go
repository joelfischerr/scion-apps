// Copyright 2020 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"os/user"
	"path"
	"strings"
)

// ParsePath performs tilde expansion, no-op if the path doesn't begin with a tilde.
func ParsePath(pth string) string {
	home := "/"

	usr, err := user.Current()
	if err == nil {
		home = usr.HomeDir
	}

	if pth == "~" {
		return path.Join(home, pth[1:])
	} else if strings.HasPrefix(pth, "~/") {
		return path.Join(home, pth[2:])
	} else {
		return pth
	}
}
