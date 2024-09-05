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

// Specify helps to organize the data to create TPPI content.
// Here we need the Type Signature, an optional Tag and a function, that
// returns the string form of the Data to be enclosed in the TPPI content.
// The Type, Tag and Data is filtered with TPPI content safeguards before being
// processed into the TPPI content form. The data function has the signature
// `func() string`, this only gives one value considered as the string
// string representation of the data.
func Specify(typeSignature, tag string, fn func() string) string {
	ta := make([]string, 0, 3)
	if len(typeSignature) == 0 {
		typeSignature = "UN"
	}
	ta = append(ta, typeSignature)
	if len(tag) > 0 {
		ta = append(ta, tag)
	}
	if fn != nil {
		ta = append(ta, fn())
	} else {
		return "" // We can't do much without this function
	}
	// Filter out the Data with TPPI content safeguards
	sa := make([]string, 0, len(ta))
	for _, i := range ta {
		i = strings.ReplaceAll(i, "|", "%7C")
		i = strings.ReplaceAll(i, "+", "%2B")
		i = strings.ReplaceAll(i, "~", "%7E")
		sa = append(sa, i)
	}
	s := strings.Join(sa, "~")
	return s
}

// Discover helps to recover the data from the TPPI content.
// In order to begin the operation we need the TPPI content in string form.
// We also need a function that would be called with the discovered
// Type Signature, Tag and Data in string form. This function helps
// to recreate the data from the String form enclosed earlier in the
// TPPI content form. Here is the signature of the function
// `func(TypeSignature, Tag, Data string) error`. The function can also
// return an error indicating that the data or any of the parameters
// are invalid or damaged. The supplied Type Signature, Tag and Data are
// restored from the fileted TPPI content safeguards before calling the
// function.
func Discover(s string, fn func(string, string, string) error) (err error) {
	if len(s) == 0 || fn == nil {
		err = fmt.Errorf("wrong inputs for discover operation")
		return
	}
	if strings.ContainsAny(s, "|+") {
		err = fmt.Errorf("wrong data or damage can't discover")
		return
	}
	// Split with Separator
	sa := strings.Split(s, "~")
	// Filter
	if len(sa) > 1 {
		for i, s := range sa {
			s = strings.ReplaceAll(s, "%7C", "|")
			s = strings.ReplaceAll(s, "%2B", "+")
			s = strings.ReplaceAll(s, "%7E", "~")
			sa[i] = s
		}
	}
	// Check
	switch len(sa) {
	case 2:
		err = fn(sa[0], "", sa[1])
	case 3:
		err = fn(sa[0], sa[1], sa[2])
	default:
		err = fmt.Errorf("invalid data discovered")
	}
	if err != nil {
		err = fmt.Errorf("failed to discover due to process error - %v",
			err)
		return
	}
	return
}
