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

import "github.com/ikawaha/kagome.ko/internal/dic"

// Dic represents a dictionary.
type Dic struct {
	dic *dic.Dic
}

// SysDic returns the system dictionary (KO dictionary).
func SysDic() Dic {
	return Dic{dic.SysDic()}
}

// SysDicKO returns the KO dictionary as the system dictionary.
func SysDicKO() Dic {
	return Dic{dic.SysDicKO()}
}

// SysDicKOSimple returns the simple KO dictionary as the system dictionary (w/o contents).
func SysDicKOSimple() Dic {
	return Dic{dic.SysDicKOSimple()}
}

// NewDic loads a dictionary from a file.
func NewDic(path string) (Dic, error) {
	d, err := dic.Load(path)
	return Dic{d}, err
}
