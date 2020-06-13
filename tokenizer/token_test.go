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

package tokenizer

import (
	"fmt"
	"testing"

	"github.com/ikawaha/kagome.ko/internal/lattice"
)

func TestTokenClassString(t *testing.T) {
	data := []struct {
		inp      TokenClass
		expected string
	}{
		{DUMMY, "DUMMY"},
		{KNOWN, "KNOWN"},
		{UNKNOWN, "UNKNOWN"},
		{USER, "USER"},
	}

	for _, v := range data {
		if got := v.inp.String(); got != v.expected {
			t.Errorf("got %v, expected %v", got, v.expected)
		}
	}
}

func TestTokenString(t *testing.T) {
	tok := Token{
		ID:      123,
		Class:   TokenClass(lattice.DUMMY),
		Start:   0,
		End:     1,
		Surface: "テスト",
	}
	expected := "テスト(0, 1)DUMMY[123]"
	if got := fmt.Sprintf("%v", tok); got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
