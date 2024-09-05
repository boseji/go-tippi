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

func ExampleAssemble() {
	ba := []myBool{
		{true, "Bit~0"},
		{false, "Bit-1"},
		{false, "Bit-2"},
		{true, "Bit-3"},
	}
	sa := make([]string, 0, len(ba))
	for _, b := range ba {
		sa = append(sa, tppi.Specify(b.Type(), b.Tag, b.String))
	}
	s := tppi.Assemble(sa...)
	s = tppi.PacketJoin(s)
	fmt.Println(s)

	// Output:
	// ~|B~Bit%7E0~1|B~Bit-1~0|B~Bit-2~0|B~Bit-3~1|~
}

func ExampleDisassemble() {
	ba := make([]bool, 4)
	ta := make([]string, 4)
	srcBool := "~|B~Bit%7E0~1|B~Bit-1~0|B~Bit-2~0|B~Bit-3~1|~"
	s := tppi.SplitPacket(srcBool)
	if len(s) != 1 {
		fmt.Println("expected it to 1 packet")
		return
	}
	sa := tppi.Disassemble(s[0])
	if len(sa) != len(ba) {
		fmt.Println("Packet does not contain enough data")
		return
	}
	for i, s := range sa {
		err := tppi.Discover(s, func(s1, s2, s3 string) error {
			if s1 != "B" {
				return fmt.Errorf("invalid type Signature")
			}
			ta[i] = s2
			switch s3 {
			case "1":
				ba[i] = true
			case "0":
				ba[i] = false
			default:
				return fmt.Errorf("invalid value")
			}
			return nil
		})
		if err != nil {
			fmt.Println("unable to discover -", err)
			return
		}
	}
	fmt.Println(ba)
	fmt.Println(ta)

	// Output:
	// [true false false true]
	// [Bit~0 Bit-1 Bit-2 Bit-3]
}

func ExamplePacketJoin() {
	sa := []string{
		"~|B~Bit%7E0~1|B~Bit-1~0|B~Bit-2~0|B~Bit-3~1|~",
		"~|S~Bit Wise Data|~",
	}
	s := tppi.PacketJoin(sa...)
	fmt.Println(s)

	// Output:
	// ~|B~Bit%7E0~1|B~Bit-1~0|B~Bit-2~0|B~Bit-3~1|~+~|S~Bit Wise Data|~
}

func ExampleSplitPacket() {
	s := "~|B~Bit%7E0~1|B~Bit-1~0|B~Bit-2~0|B~Bit-3~1|~+~|S~Bit Wise Data|~"
	sa := tppi.SplitPacket(s)
	fmt.Println(sa)

	// Output:
	// [~|B~Bit%7E0~1|B~Bit-1~0|B~Bit-2~0|B~Bit-3~1|~ ~|S~Bit Wise Data|~]
}
