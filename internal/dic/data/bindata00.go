package data

import (
	"os"
	"time"
)

var _dicIpaChardefDic = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xdc\x71\x4f\xda\x4c\x18\x00\xf0\xbb\x2a\xd2\x57\x49\xde\xcf\x74\x63\x4c\x19\xc8\x0c\xea\x1f\x66\x9f\x67\x26\xf3\x3b\x9b\x2e\x28\xba\x0d\x65\x80\xa5\x1e\x95\xdf\x2f\x69\x42\xaf\xc7\xf5\xb9\x7b\xae\x07\xc9\x25\xed\x55\x3f\x8a\x58\x54\xb7\x21\xf6\x42\xf8\x5e\xdd\x86\x93\xee\xe7\xc1\x97\x74\x3d\xbe\xea\x5c\x5e\xa4\xfe\xa0\x33\x4a\x93\xaf\xc3\xa3\xcb\x9b\xf3\x4f\xdf\xc6\xdd\xc9\xf5\xf9\x60\x3a\xec\x77\xd2\xf8\xe2\x2c\x95\x67\xc3\x69\x3a\x4d\x93\x54\x8e\xd2\x55\x1a\xa5\x49\xea\x3d\x54\x7e\xaa\x74\x3a\x1d\x0c\x46\x65\xff\x66\x3a\x1c\x8f\x87\xfd\xfb\x18\x8e\x8e\xc3\x7d\x0c\xbf\xc5\x18\xc3\x9f\xe7\x4f\xc5\x07\x7f\x3b\x7c\x36\x2f\xe8\x2c\xb5\xce\xf5\x97\x37\x5c\x74\xf0\x6f\xcb\x9b\xa7\x41\xab\xf3\xc6\x52\xff\xed\x88\x59\x2c\xc7\x7b\xeb\x31\x17\xaf\x97\x02\x00\x00\x00\x00\x00\xb4\x57\xee\x8d\xc4\x7c\x72\x8f\xfc\x6e\x58\xb1\xbd\xbe\x75\x87\x1b\x7a\xef\xf8\xea\xda\xb4\x7f\x6d\xef\x2f\xed\xf6\xd1\xd6\x97\x4d\xed\x7c\x80\x34\x49\xfa\xc9\xaa\xe9\xf5\xb7\xed\xeb\x33\xcd\xca\x3d\xff\xd8\x6f\x45\x06\xb9\xef\xdf\xbc\xf5\xc7\xbf\xf6\xfa\x71\xd4\xac\xd0\x6d\x56\x03\x53\x7a\xa7\x2d\xf6\x7f\x3b\xbf\x22\xad\xb5\xf7\xff\x7f\x72\xaf\x54\x00\x00\x00\x00\x00\x00\x14\x1b\xed\xef\xae\x52\x16\x45\x59\xcc\x8e\xb2\x81\x30\x37\x6c\xb3\x7c\xdb\xd7\xa0\x96\xfa\xf3\x6d\xad\x16\x9a\x79\xc6\xc8\x4b\x56\x61\x17\x79\x2e\x01\x00\x00\x00\x00\x00\x00\xe0\x5d\xd8\xa2\x07\x00\x00\x00\x00\x00\x00\x00\x00\x00\x3e\xaa\xed\xbd\xef\x11\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xa0\x69\xb9\xdf\xec\xdc\x2e\xb5\xc6\x2b\x77\xaa\x81\x9a\x0e\x6a\xc8\x1d\xfb\xab\x16\x62\x3c\x7c\x36\x8f\xb9\xb3\xd4\x63\x85\x55\xd7\x1f\x74\xdf\x6c\xcb\xdd\x7b\x91\x83\x5e\xf5\xb3\x88\x45\x75\x17\x62\x11\xc2\xff\xd5\x5d\x38\x09\x21\xc4\x18\x67\x47\x7c\x38\x9f\x7f\x9e\xf9\x15\x00\x00\xff\xff\x33\x0b\x28\x9c\xa0\x00\x01\x00"

func dicIpaChardefDicBytes() ([]byte, error) {
	return bindataRead(
		_dicIpaChardefDic,
		"dic/ipa/chardef.dic",
	)
}

func dicIpaChardefDic() (*asset, error) {
	bytes, err := dicIpaChardefDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/ipa/chardef.dic", size: 65696, mode: os.FileMode(420), modTime: time.Unix(1559270655, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
