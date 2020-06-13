// Copyright 2020 ikawaha
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// 	You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dic

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMorphsSave(t *testing.T) {
	m := []Morph{
		{LeftID: 1, RightID: 1, Weight: 1},
		{LeftID: 2, RightID: 2, Weight: 2},
		{LeftID: 3, RightID: 3, Weight: 3},
	}
	var b bytes.Buffer
	n, err := MorphSlice(m).WriteTo(&b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if n != int64(b.Len()) {
		t.Errorf("got %v, expected %v", n, b.Len())
	}
}

func TestLoadMorphSlice(t *testing.T) {
	expected := []Morph{
		{LeftID: 1, RightID: 1, Weight: 1},
		{LeftID: 2, RightID: 2, Weight: 2},
		{LeftID: 3, RightID: 3, Weight: 3},
	}
	var b bytes.Buffer
	_, err := MorphSlice(expected).WriteTo(&b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	got, err := LoadMorphSlice(&b)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
