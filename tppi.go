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

package tppi

import (
	"fmt"
	"strings"
)

// Assemble encloses the TPPI contents into the TPPI packet with safeguards.
func Assemble(sa ...string) (s string) {
	for _, i := range sa {
		// Filter out data with TPPI packet safeguard
		i = strings.ReplaceAll(i, "|", "\\x7C")
		i = strings.ReplaceAll(i, "+", "\\x2B")
		//i = strings.ReplaceAll(i, "~", "\\x7E")
		s += i + "|"
	}
	// Remove the last Pipe Symbol if content is provided
	s, _ = strings.CutSuffix(s, "|")
	s = "~|" + s + "|~"
	return
}

// Disassemble restores collection of TPPI Contents from the enclosed
// TPPI packet removing the necessary safeguards.
func Disassemble(s string) (sa []string) {
	t, _ := strings.CutPrefix(s, "~|")
	t, _ = strings.CutSuffix(t, "|~")
	ta := strings.Split(t, "|")
	sa = make([]string, 0, len(ta))
	for _, i := range ta {
		i = strings.ReplaceAll(i, "\\x7C", "|")
		i = strings.ReplaceAll(i, "\\x2B", "+")
		//i = strings.ReplaceAll(i, "\\x7E", "~")
		sa = append(sa, i)
	}
	return
}

// ValidPacket verifies if the packet indeed conforms to TPPI protocol or not.
// It also validates the TPPI contents to follow the required guidelines.
func ValidPacket(s string) (err error) {
	var sp []string
	// Multi Protocol Strings
	if strings.Contains(s, "+") {
		sp = strings.Split(s, "+")
	} else {
		sp = append(sp, s)
	}
	for _, p := range sp {
		if !strings.HasPrefix(p, "~|") {
			err = fmt.Errorf("not valid packet as missing start indicator")
			return
		}
		if !strings.HasSuffix(p, "|~") {
			err = fmt.Errorf("not valid packet as missing end indicator")
		}
	}
	return
}

// PacketJoin helps to join multiple TPPI packets together.
// This can optionally be run after the Assemble function to bring together
// Multiple TPPI packets together.
func PacketJoin(sa ...string) string {
	return strings.Join(sa, "+")
}

// SplitPacket allows to get back multiple TPPI packets.
// This needs to be run before running the Disassemble function.
func SplitPacket(s string) []string {
	return strings.Split(s, "+")
}
