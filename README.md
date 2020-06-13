Kagome Korean Tokenizer
===

kagome.ko is a Korean version of [kagome](https://github.com/ikawaha/kagome).
This package supports the [mecab-ko-dic](https://bitbucket.org/eunjeon/mecab-ko-dic/downloads/mecab-ko-dic-2.1.1-20180720.tar.gz).

# Programming example

Below is a simple go example that demonstrates how a simple text can be segmented.

See also https://github.com/ikawaha/kagome 


sample code:

```go:example
package main

import (
	"fmt"
	"strings"

	"github.com/ikawaha/kagome.ipadic/tokenizer"
)

func main() {
	t := tokenizer.New()
	tokens := t.Tokenize("초밥이 먹고 싶다.")
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
	}
}
```

output:

```text:outputs
BOS
초밥	NNG,*,T,초밥,Compound,*,*,초/NNG/*+밥/NNG/*
이	JKS,*,F,이,*,*,*,*
먹	VV,*,T,먹,*,*,*,*
고	EC,*,F,고,*,*,*,*
싶	VX,*,T,싶,*,*,*,*
다	EF,*,F,다,*,*,*,*
.	SF,*,*,*,*,*,*,*
EOS
```
# Dictionary version

This repository contains [mecab-ko-dic-2.1.1-20180720](https://bitbucket.org/eunjeon/mecab-ko-dic/downloads/mecab-ko-dic-2.1.1-20180720.tar.gz).

## Dictionary format

Information about the dictionary format and part-of-speech tags used by mecab-ko-dic id documented in [this Google Spreadsheet](https://docs.google.com/spreadsheets/d/1-9blXKjtjeKZqsf4NzHeYJCrr49-nXeRF6D80udfcwY/edit#gid=589544265), linked to from mecab-ko-dic's [repository readme](https://bitbucket.org/eunjeon/mecab-ko-dic/src/master/README.md).

Note how ko-dic has one less feature column than NAIST JDIC, and has an altogether different set of information (e.g. doesn't provide the "original form" of the word).

The tags are a slight modification of those specified by 세종 (Sejong), whatever that is. The mappings from Sejong to mecab-ko-dic's tag names are given in tab `태그 v2.0` on the above-linked spreadsheet.

The dictionary format is specified fully (in Korean) in tab `사전 형식 v2.0` of the spreadsheet. Any blank values default to `*`.

| Index | Name (Korean) | Name (English) | Notes |
| --- | --- | --- | --- |
| 0 | 품사 태그 | part-of-speech tag | See `태그 v2.0` tab on spreadsheet  |
| 1 | 의미 부류 | meaning | (too few examples for me to be sure) |
| 2 | 종성 유무 | presence or absence | `T` for true; `F` for false; else `*` |
| 3 | 읽기 | reading | usually matches surface, but may differ for foreign words e.g. Chinese character words |
| 4 | 타입 | type | One of: `Inflect` (활용); `Compound` (복합명사); or `Preanalysis` (기분석) |
| 5 | 첫번째 품사 | first part-of-speech | e.g. given a part-of-speech tag of "VV+EM+VX+EP", would return `VV` |
| 6 | 마지막 품사 | last part-of-speech | e.g. given a part-of-speech tag of "VV+EM+VX+EP", would return `EP` |
| 7 | 표현 | expression | `활용, 복합명사, 기분석이 어떻게 구성되는지 알려주는 필드` – Fields that tell how usage, compound nouns, and key analysis are organized |

License
---
Kagome is licensed under the Apache License v2.0 and uses the MeCab-Ko model. See NOTICE.txt for license details.
