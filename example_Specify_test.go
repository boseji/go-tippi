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
	"time"

	"github.com/boseji/go-tppi"
)

type myBool struct {
	bool
	Tag string
}

const (
	myBoolTypeSignature = "B"
	myBoolDataTrue      = "1"
	myBoolDataFalse     = "0"
)

func (b myBool) Type() string {
	return myBoolTypeSignature
}

func (b myBool) String() string {
	if b.bool {
		return myBoolDataTrue
	}
	return myBoolDataFalse
}

type myStruct struct {
	TimeStamp time.Duration
	Value     float32
	Topic     string
}

const (
	myStructTypeSig = "MSTR"
	myStructFormat  = "%d %f %q"
)

func (m myStruct) Type() string {
	return myStructTypeSig
}

func (m myStruct) String() string {
	return fmt.Sprintf(myStructFormat, m.TimeStamp, m.Value, m.Topic)
}

func ExampleSpecify() {
	b := myBool{true, "bit~5"}
	s := tppi.Specify(b.Type(), b.Tag, b.String)
	fmt.Printf("%#v transforms into %q\n", b, s)

	m := myStruct{134 * time.Millisecond, 12.345, "sensor/12"}
	sa := tppi.Specify(m.Type(), "myStruct", m.String)
	fmt.Printf("%#v\n", m)
	fmt.Println("transforms into")
	fmt.Println(sa)

	// Output:
	// tppi_test.myBool{bool:true, Tag:"bit~5"} transforms into "B~bit%7E5~1"
	// tppi_test.myStruct{TimeStamp:134000000, Value:12.345, Topic:"sensor/12"}
	// transforms into
	// MSTR~myStruct~134000000 12.345000 "sensor/12"
}
