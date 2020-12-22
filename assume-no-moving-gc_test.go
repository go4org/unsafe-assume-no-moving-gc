// Copyright 2020 Brad Fitzpatrick. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package assume_no_moving_gc

import (
	"os"
	"runtime"
	"testing"
)

func TestInit(t *testing.T) {
	if e := os.Getenv(env); e != "" {
		t.Skipf("env %s is set to %q; skipping test", env, e)
	}
	t.Logf("%s doesn't use a moving collector", runtime.Version())
}
