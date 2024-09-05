//
//     ॐ भूर्भुवः स्वः
//     तत्स॑वि॒तुर्वरे॑ण्यं॒
//    भर्गो॑ दे॒वस्य॑ धीमहि।
//   धियो॒ यो नः॑ प्रचो॒दया॑त्॥
//
//
// बोसजी के द्वारा रचित टिप्पी अधिलेखन प्रकृया।
// ================================
//
// एक सरल संचार सहायक और संलग्न तंत्र।
//
// ~~~~~~~~~~~~~~~~~~~~~~~
// एक रचनात्मक भारतीय उत्पाद।
// ~~~~~~~~~~~~~~~~~~~~~~~
//
//
// Sources
// --------
// https://github.com/boseji/go-tppi
//
//
// License
// --------
//
//   SPDX: Apache-2.0
//
//   Copyright (C) 2024 Abhijit Bose (aka. Boseji). All rights reserved.
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//
//   SPDX short identifier: Apache-2.0

package tppi_test

import (
	"fmt"

	"github.com/boseji/go-tppi"
)

func ExampleSpecify() {
	b := struct {
		bool
		Tag string
	}{true, "bit~5"}
	s := tppi.Specify("B", b.Tag, func() string {
		if b.bool {
			return "1"
		}
		return "0"
	})
	fmt.Printf("%#v transforms into %q\n", b, s)

	f := 12.3456
	s2 := tppi.Specify("F", "", func() string {
		return fmt.Sprintf("%2.5f", f)
	})
	fmt.Printf("%#v transforms into %q\n", f, s2)

	// Output:
	// struct { bool; Tag string }{bool:true, Tag:"bit~5"} transforms into "B~bit%7E5~1"
	// 12.3456 transforms into "F~12.34560"
}
