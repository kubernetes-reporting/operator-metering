/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this ***REMOVED***le except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the speci***REMOVED***c language governing permissions and
limitations under the License.
*/

package clientcmd

import (
	"fmt"
	"strconv"
	"time"
)

// ParseTimeout returns a parsed duration from a string
// A duration string value must be a positive integer, optionally followed by a corresponding time unit (s|m|h).
func ParseTimeout(duration string) (time.Duration, error) {
	if i, err := strconv.ParseInt(duration, 10, 64); err == nil && i >= 0 {
		return (time.Duration(i) * time.Second), nil
	}
	if requestTimeout, err := time.ParseDuration(duration); err == nil {
		return requestTimeout, nil
	}
	return 0, fmt.Errorf("Invalid timeout value. Timeout must be a single integer in seconds, or an integer followed by a corresponding time unit (e.g. 1s | 2m | 3h)")
}
