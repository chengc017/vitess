// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mytype

type MyType struct {
	Ptr       *int64
	PtrPtr    **int64
	PtrBytes  *[]byte
	PtrSlice  *[]int64
	PtrMap    *map[string]int64
	PtrCustom *Custom
}
