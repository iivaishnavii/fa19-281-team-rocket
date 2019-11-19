// Copyright (c) 2017 David R. Jenni. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.9

package main

import "go/types"

func isAlias(obj *types.TypeName) bool { return obj.IsAlias() }
