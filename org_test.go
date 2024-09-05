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
	"reflect"
	"strings"
	"testing"
)

func Test_Specify(t *testing.T) {
	type args struct {
		typeSig string
		tag     string
		sfn     func() string
	}
	tests := []struct {
		name  string
		args  args
		wantS string
	}{
		{
			name: "String Type",
			args: args{
				typeSig: "S",
				tag:     "",
				sfn: func() string {
					return "Test String"
				},
			},
			wantS: "S~Test String",
		},
		{
			name: "String with tag",
			args: args{
				typeSig: "S",
				tag:     "value",
				sfn: func() string {
					return "Test String"
				},
			},
			wantS: "S~value~Test String",
		},
		{
			name: "String with no function",
			args: args{
				typeSig: "S",
				tag:     "",
				sfn:     nil,
			},
			wantS: "",
		},
		{
			name: "String without type and Tag",
			args: args{
				typeSig: "",
				tag:     "Testing",
				sfn: func() string {
					return "Test String"
				},
			},
			wantS: "UN~Testing~Test String",
		},
		{
			name: "String with special Characters",
			args: args{
				typeSig: "S",
				tag:     "",
				sfn: func() string {
					return "Test String+ with| several~Special Characters"
				},
			},
			wantS: "S~Test String%2B with%7C several%7ESpecial Characters",
		},
		{
			name: "Int Type",
			args: args{
				typeSig: "I",
				tag:     "",
				sfn: func() string {
					return fmt.Sprintf("%x", 0x3508)
				},
			},
			wantS: "I~3508",
		},
		{
			name: "Error Type",
			args: args{
				typeSig: "E",
				tag:     "",
				sfn: func() string {
					return fmt.Errorf("custom Error").Error()
				},
			},
			wantS: "E~custom Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Specify(tt.args.typeSig, tt.args.tag, tt.args.sfn)
			if strings.Compare(got, tt.wantS) != 0 {
				t.Errorf("error in values got %v want %v", got, tt.wantS)
			}
		})
	}
}

func Test_Discover(t *testing.T) {
	tests := []struct {
		name           string
		arg            string
		wantSa         []string
		wantErr        bool
		wantInnerError bool
	}{
		{
			name:   "String Type",
			arg:    "S~Test String",
			wantSa: []string{"S", "", "Test String"},
		},
		{
			name:   "String with tag",
			arg:    "S~value~Test String",
			wantSa: []string{"S", "value", "Test String"},
		},
		{
			name:   "String with tag with special chars",
			arg:    "S~value%7C4~Test String",
			wantSa: []string{"S", "value|4", "Test String"},
		},
		{
			name:    "Negative empty string",
			arg:     "",
			wantErr: true,
		},
		{
			name:           "Negative Inner Error string",
			arg:            "S~value~Test String",
			wantErr:        true,
			wantInnerError: true,
		},
		{
			name:    "Negative corrupt string",
			arg:     "S~val|ue~Test Str+ing",
			wantErr: true,
		},
		{
			name:    "Negative corrupt string2",
			arg:     "S value Test String",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa := make([]string, 0, 3)
			absorb := func(s1, s2, s3 string) error {
				sa = append(sa, s1, s2, s3)
				if tt.wantInnerError {
					return fmt.Errorf("inner error")
				}
				return nil
			}
			err := Discover(tt.arg, absorb)
			if (err != nil) != tt.wantErr {
				t.Errorf("error in Discover() - %v want error %v",
					err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(tt.wantSa, sa) {
				t.Errorf("error in values got %v want %v", sa, tt.wantSa)
			}
		})
	}
}
