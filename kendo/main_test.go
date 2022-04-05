// Copyright 2022 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"testing"
)

func TestSt(t *testing.T) {
	var s1 []int= make([]int, 4, 6)
	fmt.Println(len(s1), cap(s1))
}
