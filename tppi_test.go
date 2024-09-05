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
	"reflect"
	"strings"
	"testing"
)

func Test_Assemble(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		wantS string
	}{
		{
			name: "Basic test of Assembly",
			args: []string{
				"P1",
				"P2",
			},
			wantS: "~|P1|P2|~",
		},
		{
			name: "Assemble string containing special chars",
			args: []string{
				"P1|",
				"P~2",
			},
			wantS: "~|P1\\x7C|P~2|~",
		},
		{
			name: "Assemble string containing multiple special chars",
			args: []string{
				"Pa~rt1|",
				"P|+art~2",
			},
			wantS: "~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
		},
		{
			name:  "Blank contents",
			args:  []string{},
			wantS: "~||~",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Assemble(tt.args...)
			if strings.Compare(got, tt.wantS) != 0 {
				t.Errorf("error in values got %v want %v", got, tt.wantS)
			}
		})
	}
}

func Test_Disassemble(t *testing.T) {
	tests := []struct {
		name   string
		args   string
		wantSa []string
	}{
		{
			name: "Basic test of Assembly",
			args: "~|P1|P2|~",
			wantSa: []string{
				"P1",
				"P2",
			},
		},
		{
			name: "Assemble string containing special chars",
			args: "~|P1\\x7C|P~2|~",
			wantSa: []string{
				"P1|",
				"P~2",
			},
		},
		{
			name: "Assemble string containing multiple special chars",
			args: "~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
			wantSa: []string{
				"Pa~rt1|",
				"P|+art~2",
			},
		},
		{
			name:   "Blank contents",
			args:   "",
			wantSa: []string{""},
		},
		{
			name:   "Blank packet",
			args:   "~||~",
			wantSa: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Disassemble(tt.args)
			if !reflect.DeepEqual(got, tt.wantSa) {
				t.Errorf("error in values got %v want %v", got, tt.wantSa)
			}
		})
	}
}

func Test_ValidPacket(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{
			name: "Basic Single Packet",
			arg:  "~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
		},
		{
			name: "Dual Packets",
			arg:  "~|P1\\x7C|P~2|~+~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
		},
		{
			name:    "Error Single Packet1",
			arg:     "~Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
			wantErr: true,
		},
		{
			name:    "Error Single Packet2",
			arg:     "~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|",
			wantErr: true,
		},
		{
			name:    "Error Single Packet3",
			arg:     "~|Pa~rt1\\x7C|~P+\\x2Bart~2|~",
			wantErr: true,
		},
		{
			name:    "Error Single Packet4",
			arg:     "~|Pa~rt1\\x7C+P\\x7C\\x2Bart~2|~",
			wantErr: true,
		},
		{
			name:    "Error Single Packet5",
			arg:     "~|Pa~rt1\\x7C|~~|+~|P\\x7C\\x2Bart~2|~",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidPacket(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("error ValidPacket() got %v want error %v",
					err, tt.wantErr)
			}
		})
	}
}

func Test_PacketJoin(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		wantS string
	}{
		{
			name: "Join a single packet",
			args: []string{
				"~||~",
			},
			wantS: "~||~",
		},
		{
			name: "Join two packet",
			args: []string{
				"~||~",
				"~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
			},
			wantS: "~||~+~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PacketJoin(tt.args...)
			if strings.Compare(got, tt.wantS) != 0 {
				t.Errorf("error in values got %v want %v", got, tt.wantS)
			}
		})
	}
}

func Test_PacketSplit(t *testing.T) {
	tests := []struct {
		name   string
		arg    string
		wantSa []string
	}{
		{
			name: "Join a single packet",
			arg:  "~||~",
			wantSa: []string{
				"~||~",
			},
		},
		{
			name: "Join two packet",
			arg:  "~||~+~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
			wantSa: []string{
				"~||~",
				"~|Pa~rt1\\x7C|P\\x7C\\x2Bart~2|~",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitPacket(tt.arg)
			if !reflect.DeepEqual(got, tt.wantSa) {
				t.Errorf("error in values got %v want %v", got, tt.wantSa)
			}
		})
	}
}
