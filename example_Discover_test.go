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
	"strings"

	"guthub.com/boseji/go-tppi"
)

func (b *myBool) Recover(Type, Tag, Data string) error {
	if Type != myBoolTypeSignature {
		return fmt.Errorf("invalid bool type")
	}
	b.Tag = Tag
	if Data == myBoolDataTrue {
		b.bool = true
	} else if Data == myBoolDataFalse {
		b.bool = false
	} else {
		return fmt.Errorf("invalid bool data")
	}
	return nil
}

func (m *myStruct) Recover(Type, Tag, Data string) (err error) {
	if Type != myStructTypeSig {
		return fmt.Errorf("wrong type supplied")
	}
	_ = Tag
	sa := strings.Split(Data, " ")
	if len(sa) != 3 {
		return fmt.Errorf("malformed data")
	}

	_, err = fmt.Sscanf(Data, myStructFormat, &m.TimeStamp, &m.Value, &m.Topic)
	return
}

func ExampleDiscover() {
	b := myBool{}
	srcBool := "B~bit%7E5~1"
	err := tppi.Discover(srcBool, b.Recover)
	if err != nil {
		fmt.Printf("failed to Discover - %v\n", err)
		return
	}
	fmt.Printf("%q Transforms back %#v\n", srcBool, b)

	m := myStruct{}
	srcStruct := "MSTR~myStruct~134000000 12.345000 \"sensor/12\""
	err = tppi.Discover(srcStruct, m.Recover)
	if err != nil {
		fmt.Printf("failed to Discover - %v\n", err)
		return
	}
	fmt.Printf("%q Transforms back\n", srcStruct)
	fmt.Printf("%#v", m)

	// Output:
	// "B~bit%7E5~1" Transforms back tppi_test.myBool{bool:true, Tag:"bit~5"}
	// "MSTR~myStruct~134000000 12.345000 \"sensor/12\"" Transforms back
	// tppi_test.myStruct{TimeStamp:134000000, Value:12.345, Topic:"sensor/12"}
}
