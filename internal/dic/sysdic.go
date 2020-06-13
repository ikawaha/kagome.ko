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
	"sync"

	"github.com/ikawaha/kagome.ko/internal/dic/data"
)

const (
	// KoDicPath represents the internal KO dictionary path.
	KoDicPath = "dic"
)

var (
	sysDicKOFull     *Dic
	initSysDicKOFull sync.Once

	sysDicKOSimple     *Dic
	initSysDicKOSimple sync.Once
)

// SysDic returns the kagome system dictionary.
func SysDic() *Dic {
	return SysDicKO()
}

// SysDicSimple returns the kagome system dictionary without contents.
func SysDicSimple() *Dic {
	return SysDicKOSimple()
}

// SysDicKO returns the KO system dictionary.
func SysDicKO() *Dic {
	initSysDicKOFull.Do(func() {
		sysDicKOFull = loadInternalSysDicFull(KoDicPath)
	})
	return sysDicKOFull
}

// SysDicKOSimple returns the KO system dictionary without contents.
func SysDicKOSimple() *Dic {
	initSysDicKOSimple.Do(func() {
		sysDicKOSimple = loadInternalSysDicSimple(KoDicPath)
	})
	return sysDicKOSimple
}

func loadInternalSysDicFull(path string) (d *Dic) {
	return loadInternalSysDic(path, true)
}

func loadInternalSysDicSimple(path string) (d *Dic) {
	return loadInternalSysDic(path, false)
}

func loadInternalSysDic(path string, full bool) (d *Dic) {
	d = new(Dic)
	var (
		buf []byte
		err error
	)
	// morph.dic
	if buf, err = data.Asset(path + "/morph.dic"); err != nil {
		panic(err)
	}
	if err = d.loadMorphDicPart(bytes.NewBuffer(buf)); err != nil {
		panic(err)
	}
	// pos.dic
	if buf, err = data.Asset(path + "/pos.dic"); err != nil {
		panic(err)
	}
	if err = d.loadPOSDicPart(bytes.NewBuffer(buf)); err != nil {
		panic(err)
	}
	if full {
		// content.dic
		if buf, err = data.Asset(path + "/content.dic"); err != nil {
			panic(err)
		}
		d.Contents = NewContents(buf)
	}
	// index.dic
	if buf, err = data.Asset(path + "/index.dic"); err != nil {
		panic(err)
	}
	if err = d.loadIndexDicPart(bytes.NewBuffer(buf)); err != nil {
		panic(err)
	}
	// connection.dic
	if buf, err = data.Asset(path + "/connection.dic"); err != nil {
		panic(err)
	}
	if err = d.loadConnectionDicPart(bytes.NewBuffer(buf)); err != nil {
		panic(err)
	}
	// chardef.dic
	if buf, err = data.Asset(path + "/chardef.dic"); err != nil {
		panic(err)
	}
	if err = d.loadCharDefDicPart(bytes.NewBuffer(buf)); err != nil {
		panic(err)
	}
	// unk.dic
	if buf, err = data.Asset(path + "/unk.dic"); err != nil {
		panic(err)
	}
	if err = d.loadUnkDicPart(bytes.NewBuffer(buf)); err != nil {
		panic(err)
	}
	return
}
