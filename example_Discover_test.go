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

func ExampleDiscover() {
	b := struct {
		bool
		Tag string
	}{}
	srcBool := "B~bit%7E5~1"
	err := tppi.Discover(srcBool, func(s1, s2, s3 string) error {
		if s1 != "B" {
			return fmt.Errorf("wrong type")
		}
		b.Tag = s2
		if s3 == "1" {
			b.bool = true
		} else if s3 == "0" {
			b.bool = false
		} else {
			return fmt.Errorf("unknown data")
		}
		return nil
	})
	if err != nil {
		fmt.Printf("failed to Discover - %v\n", err)
		return
	}
	fmt.Printf("%q Transforms back %#v\n", srcBool, b)

	srcFloat := "F~12.34560"
	f := 0.00
	err = tppi.Discover(srcFloat, func(s1, s2, s3 string) error {
		_ = s1
		_ = s2
		_, err := fmt.Sscanf(s3, "%f", &f)
		return err
	})
	if err != nil {
		fmt.Printf("failed to Discover - %v\n", err)
		return
	}
	fmt.Printf("%q Transforms back %#v\n", srcFloat, f)

	// Output:
	// "B~bit%7E5~1" Transforms back struct { bool; Tag string }{bool:true, Tag:"bit~5"}
	// "F~12.34560" Transforms back 12.3456
}
