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
	"reflect"
	"testing"
)

var testDic = "../../_sample/ko.dic"

const (
	KoDicEntrySize = 816283 // cat mecab-ko-dic-2.1.1-20180720/*.csv|wc -l
)

func TestDicLoad(t *testing.T) {
	dic, err := Load(testDic)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	const expected = KoDicEntrySize
	if c := len(dic.Morphs); c != expected {
		t.Errorf("morphs got %v, expected %v", c, expected)
	}
	if c := len(dic.POSTable.POSs); c != expected {
		t.Errorf("POSs got %v, expected %v", c, expected)
	}
	if c := len(dic.Contents); c != expected {
		t.Errorf("contents got %v, expected %v", c, expected)
	}
}

func TestDicCharClass(t *testing.T) {
	dic, err := Load(testDic)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []string{
		"DEFAULT",      // 0
		"SPACE",        // 1
		"HANJA",        // 2
		"KANJI",        // 3
		"SYMBOL",       // 4
		"NUMERIC",      // 5
		"ALPHA",        // 6
		"HANGUL",       // 7
		"HIRAGANA",     // 8
		"KATAKANA",     // 9
		"HANJANUMERIC", // 10
		"GREEK",        // 11
		"CYRILLIC",     // 12
	}
	if !reflect.DeepEqual(dic.CharClass, expected) {
		t.Errorf("got %v, expected %v", dic.CharClass, expected)
	}
}

func TestDicCharCategoryMap(t *testing.T) {
	const (
		DEFAULT = iota
		SPACE
		HANJA
		KANJI
		SYMBOL
		NUMERIC
		ALPHA
		HANGUL
		HIRAGANA
		KATAKANA
		HANJANUMERIC
		GREEK
		CYRILLIC
	)
	dic, err := Load(testDic)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	type category struct {
		input    int
		category byte
	}
	// Korean & Japanese character category map
	// CODE(UCS2) TO CATEGORY MAPPING
	var data = []category{
		// SPACE
		{input: 0x0020, category: SPACE}, // DO NOT REMOVE THIS LINE, 0x0020 is reserved for SPACE
		{input: 0x000D, category: SPACE}, // CR
		{input: 0x0009, category: SPACE}, // HT
		{input: 0x000B, category: SPACE}, // VT
		{input: 0x000A, category: SPACE}, // LF
		// ASCII
		{input: 0x0021, category: SYMBOL},
		{input: 0x002F, category: SYMBOL},
		{input: 0x0030, category: NUMERIC},
		{input: 0x0039, category: NUMERIC},
		{input: 0x003A, category: SYMBOL},
		{input: 0x0040, category: SYMBOL},
		{input: 0x0041, category: ALPHA},
		{input: 0x005A, category: ALPHA},
		{input: 0x005B, category: SYMBOL},
		{input: 0x0060, category: SYMBOL},
		{input: 0x0061, category: ALPHA},
		{input: 0x007A, category: ALPHA},
		{input: 0x007B, category: SYMBOL},
		{input: 0x007E, category: SYMBOL},

		// Latin
		{input: 0x00A1, category: SYMBOL}, {input: 0x00BF, category: SYMBOL}, // Latin 1
		{input: 0x00C0, category: ALPHA}, {input: 0x00FF, category: ALPHA}, //  Latin 1
		{input: 0x0100, category: ALPHA}, {input: 0x017F, category: ALPHA}, // Latin Extended A
		{input: 0x0180, category: ALPHA}, {input: 0x0236, category: ALPHA}, // Latin Extended B
		{input: 0x1E00, category: ALPHA}, {input: 0x1EF9, category: ALPHA}, // Latin Extended Additional

		// CYRILLIC
		{input: 0x0400, category: CYRILLIC}, {input: 0x04F9, category: CYRILLIC},
		{input: 0x0500, category: CYRILLIC}, {input: 0x050F, category: CYRILLIC}, // Cyrillic supplementary

		// GREEK
		{input: 0x0374, category: GREEK}, {input: 0x03FB, category: GREEK}, // Greek and Coptic

		// HANGUL
		{input: 0xAC00, category: HANGUL}, {input: 0xD7A3, category: HANGUL},
		{input: 0x1100, category: HANGUL}, {input: 0x11FF, category: HANGUL}, // Hangul Jamo
		{input: 0x3130, category: HANGUL}, {input: 0x318F, category: HANGUL}, // Hangul Compatibility Jamo

		// HIRAGANA
		{input: 0x3041, category: HIRAGANA}, {input: 0x309F, category: HIRAGANA},

		// KATAKANA
		{input: 0x30A1, category: KATAKANA}, {input: 0x30FF, category: KATAKANA},
		{input: 0x31F0, category: KATAKANA}, {input: 0x31FF, category: KATAKANA}, // Small KU .. Small RO
		{input: 0x30FC, category: KATAKANA}, // ー

		// Half KATAKANA
		{input: 0xFF66, category: KATAKANA}, {input: 0xFF9D, category: KATAKANA},
		{input: 0xFF9E, category: KATAKANA}, {input: 0xFF9F, category: KATAKANA},

		// HANJA
		{input: 0x2E80, category: HANJA},
		{input: 0x2EF3, category: HANJA}, // CJK Raidcals Supplement
		//{input: 0x3005, category: HANJA}, // IDEOGRAPHIC ITERATION MARK
		//{input: 0x3007, category: HANJA}, // IDEOGRAPHIC NUMBER ZERO
		{input: 0x3400, category: HANJA},
		{input: 0x4DB5, category: HANJA}, // CJK Unified Ideographs Extention
		//{input: 0x4E00, category: HANJA},
		{input: 0x9FA5, category: HANJA},
		{input: 0xF900, category: HANJA},
		{input: 0xFA2D, category: HANJA},
		{input: 0xFA30, category: HANJA},
		{input: 0xFA6A, category: HANJA},

		// KANJI
		{input: 0x2F00, category: KANJI},
		{input: 0x2FD5, category: KANJI}, // KANJI Radicals

		// HANJA-NUMERIC (一 二 三 四 五 六 七 八 九 十 百 千 万 億 兆)
		{input: 0x4E00, category: HANJANUMERIC},
		{input: 0x4E8C, category: HANJANUMERIC},
		{input: 0x4E09, category: HANJANUMERIC},
		{input: 0x56DB, category: HANJANUMERIC},
		{input: 0x4E94, category: HANJANUMERIC},
		{input: 0x516D, category: HANJANUMERIC},
		{input: 0x4E03, category: HANJANUMERIC},
		{input: 0x516B, category: HANJANUMERIC},
		{input: 0x4E5D, category: HANJANUMERIC},
		{input: 0x5341, category: HANJANUMERIC},
		{input: 0x767E, category: HANJANUMERIC},
		{input: 0x5343, category: HANJANUMERIC},
		{input: 0x4E07, category: HANJANUMERIC},
		{input: 0x5104, category: HANJANUMERIC},
		{input: 0x5146, category: HANJANUMERIC},

		// ZENKAKU
		{input: 0xFF10, category: NUMERIC}, {input: 0xFF19, category: NUMERIC},
		{input: 0xFF21, category: ALPHA}, {input: 0xFF3A, category: ALPHA},
		{input: 0xFF41, category: ALPHA}, {input: 0xFF5A, category: ALPHA},
		{input: 0xFF01, category: SYMBOL}, {input: 0xFF0F, category: SYMBOL},
		{input: 0xFF1A, category: SYMBOL}, {input: 0xFF1F, category: SYMBOL},
		{input: 0xFF3B, category: SYMBOL}, {input: 0xFF40, category: SYMBOL},
		{input: 0xFF5B, category: SYMBOL}, {input: 0xFF65, category: SYMBOL},
		{input: 0xFFE0, category: SYMBOL}, {input: 0xFFEF, category: SYMBOL}, // HalfWidth and Full width Form

		// OTHER SYMBOLS
		{input: 0x2000, category: SYMBOL}, {input: 0x206F, category: SYMBOL}, // General Punctuation
		{input: 0x2070, category: NUMERIC}, {input: 0x209F, category: NUMERIC}, // Superscripts and Subscripts
		{input: 0x20A0, category: SYMBOL}, {input: 0x20CF, category: SYMBOL}, // Currency Symbols
		{input: 0x20D0, category: SYMBOL}, {input: 0x20FF, category: SYMBOL}, // Combining Diaritical Marks for Symbols
		{input: 0x2100, category: SYMBOL}, {input: 0x214F, category: SYMBOL}, // Letterlike Symbols
		{input: 0x2150, category: NUMERIC}, {input: 0x218F, category: NUMERIC}, // Number forms
		{input: 0x2100, category: SYMBOL}, {input: 0x214B, category: SYMBOL}, // Letterlike Symbols
		{input: 0x2190, category: SYMBOL}, {input: 0x21FF, category: SYMBOL}, // Arrow
		{input: 0x2200, category: SYMBOL}, {input: 0x22FF, category: SYMBOL}, // Mathematical Operators
		{input: 0x2300, category: SYMBOL}, {input: 0x23FF, category: SYMBOL}, // Miscellaneuos Technical
		{input: 0x2460, category: SYMBOL}, {input: 0x24FF, category: SYMBOL}, // Enclosed NUMERICs
		{input: 0x2501, category: SYMBOL}, {input: 0x257F, category: SYMBOL}, // Box Drawing
		{input: 0x2580, category: SYMBOL}, {input: 0x259F, category: SYMBOL}, // Block Elements
		{input: 0x25A0, category: SYMBOL}, {input: 0x25FF, category: SYMBOL}, // Geometric Shapes
		{input: 0x2600, category: SYMBOL}, {input: 0x26FE, category: SYMBOL}, // Miscellaneous Symbols
		{input: 0x2700, category: SYMBOL}, {input: 0x27BF, category: SYMBOL}, // Dingbats
		{input: 0x27F0, category: SYMBOL}, {input: 0x27FF, category: SYMBOL}, // Supplemental Arrows A
		{0x27C0, SYMBOL}, {input: 0x27EF, category: SYMBOL}, // Miscellaneous Mathematical Symbols-A
		{input: 0x2800, category: SYMBOL}, {input: 0x28FF, category: SYMBOL}, // Braille Patterns
		{input: 0x2900, category: SYMBOL}, {input: 0x297F, category: SYMBOL}, // Supplemental Arrows B
		{input: 0x2B00, category: SYMBOL}, {input: 0x2BFF, category: SYMBOL}, // Miscellaneous Symbols and Arrows
		{input: 0x2A00, category: SYMBOL}, {input: 0x2AFF, category: SYMBOL}, // Supplemental Mathematical Operators
		{input: 0x3300, category: SYMBOL}, {input: 0x33FF, category: SYMBOL},
		{input: 0x3200, category: SYMBOL}, {input: 0x32FE, category: SYMBOL}, // ENclosed CJK Letters and Months
		{input: 0x3000, category: SYMBOL}, {input: 0x303F, category: SYMBOL}, // CJK Symbol and Punctuation
		{input: 0xFE30, category: SYMBOL}, {input: 0xFE4F, category: SYMBOL}, // CJK Compatibility Forms
		{input: 0xFE50, category: SYMBOL}, {input: 0xFE6B, category: SYMBOL}, // Small Form Variants
		{input: 0x3007, category: SYMBOL}, // dded 2006/3/13
	}
	for _, v := range data {
		category := dic.CharCategory[v.input]
		if category != v.category {
			t.Errorf("input %04X, got %v, expected %v", v.input, category, v.category)
		}
		category = dic.CharacterCategory(rune(v.input))
		if category != v.category {
			t.Errorf("input %04X, got %v, expected %v", v.input, category, v.category)
		}
	}
}

func TestCharCategorySize(t *testing.T) {
	dic, err := Load(testDic)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := dic.CharacterCategory(rune(len(dic.CharCategory)))
	expected := dic.CharCategory[0]
	if c != expected {
		t.Errorf("got %v, expected %v", c, expected)
	}

}

func TestDicInvokeList(t *testing.T) {
	const (
		DEFAULT = iota
		SPACE
		HANJA
		KANJI
		SYMBOL
		NUMERIC
		ALPHA
		HANGUL
		HIRAGANA
		KATAKANA
		HANJANUMERIC
		GREEK
		CYRILLIC
	)
	var data = []struct {
		class  int
		invoke bool
	}{
		//
		//   - CATEGORY_NAME: Name of category. you have to define DEFAULT class.
		//   - INVOKE: 1/0:   always invoke unknown word processing, evan when the word can be found in the lexicon
		//   - GROUP:  1/0:   make a new word by grouping the same chracter category
		//   - LENGTH: n:     1 to n length new words are added
		{class: DEFAULT, invoke: false},     // 0 1 0  # DEFAULT is a mandatory category!
		{class: SPACE, invoke: false},       // 0 1 0
		{class: HANJA, invoke: false},       // 0 0 1
		{class: KANJI, invoke: false},       // 0 0 2
		{class: SYMBOL, invoke: true},       // 1 1 0
		{class: NUMERIC, invoke: true},      // 1 1 0
		{class: ALPHA, invoke: true},        // 1 1 0
		{class: HANGUL, invoke: false},      // 0 1 2 # Korean
		{class: HIRAGANA, invoke: true},     // 1 1 0
		{class: KATAKANA, invoke: true},     // 1 1 0
		{class: HANJANUMERIC, invoke: true}, // 1 1 0
		{class: GREEK, invoke: true},        // 1 1 0
		{class: CYRILLIC, invoke: true},     // 1 1 0
	}
	dic, err := Load(testDic)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, v := range data {
		if got := dic.InvokeList[v.class]; got != v.invoke {
			t.Errorf("input %v: got %v, expected %v", v.class, got, v.invoke)
		}
	}
}

func TestDicGroupList(t *testing.T) {
	const (
		DEFAULT = iota
		SPACE
		HANJA
		KANJI
		SYMBOL
		NUMERIC
		ALPHA
		HANGUL
		HIRAGANA
		KATAKANA
		HANJANUMERIC
		GREEK
		CYRILLIC
	)
	data := []struct {
		class  int
		invoke bool
	}{
		//
		//   - CATEGORY_NAME: Name of category. you have to define DEFAULT class.
		//   - INVOKE: 1/0:   always invoke unknown word processing, evan when the word can be found in the lexicon
		//   - GROUP:  1/0:   make a new word by grouping the same chracter category
		//   - LENGTH: n:     1 to n length new words are added
		{class: DEFAULT, invoke: true},      // 0 1 0  # DEFAULT is a mandatory category!
		{class: SPACE, invoke: true},        // 0 1 0
		{class: HANJA, invoke: false},       // 0 0 1
		{class: KANJI, invoke: false},       // 0 0 2
		{class: SYMBOL, invoke: true},       // 1 1 0
		{class: NUMERIC, invoke: true},      // 1 1 0
		{class: ALPHA, invoke: true},        // 1 1 0
		{class: HANGUL, invoke: true},       // 0 1 2 # Korean
		{class: HIRAGANA, invoke: true},     // 1 1 0
		{class: KATAKANA, invoke: true},     // 1 1 0
		{class: HANJANUMERIC, invoke: true}, // 1 1 0
		{class: GREEK, invoke: true},        // 1 1 0
		{class: CYRILLIC, invoke: true},     // 1 1 0
	}
	dic, err := Load(testDic)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, v := range data {
		if got := dic.GroupList[v.class]; got != v.invoke {
			t.Errorf("input %v: got %v, expected %v", v.class, got, v.invoke)
		}
	}
}
