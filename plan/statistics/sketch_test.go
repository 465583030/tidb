// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package statistics

import (
	. "github.com/pingcap/check"
)

func (s *testStatisticsSuite) TestSketch(c *C) {
	maxSize := 1000
	sampleSketch, ndv, err := buildSketch(s.samples, maxSize)
	c.Check(err, IsNil)
	c.Check(ndv, Equals, int64(6616))

	rcSketch, ndv, err := buildSketch(s.rc.(*recordSet).data, maxSize)
	c.Check(err, IsNil)
	c.Check(ndv, Equals, int64(74112))

	pkSketch, ndv, err := buildSketch(s.pk.(*recordSet).data, maxSize)
	c.Check(err, IsNil)
	c.Check(ndv, Equals, int64(99840))

	var sketches []*Sketch
	sketches = append(sketches, sampleSketch)
	sketches = append(sketches, pkSketch)
	sketches = append(sketches, rcSketch)
	_, ndv = mergeSketches(sketches, maxSize)
	c.Check(ndv, Equals, int64(99840))
}