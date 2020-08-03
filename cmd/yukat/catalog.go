// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"ru": &dictionary{index: ruIndex, data: ruData},
		"uz": &dictionary{index: uzIndex, data: uzData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"Choose language":         0,
	"Done":                    7,
	"Hello <b>%s %s.</b>\n\n": 4,
	"Language changed":        1,
	"Settings are changed":    2,
	"Unknown command":         5,
	"Wrong callback input!":   3,
	"🇬🇧 Language":             6,
}

var enIndex = []uint32{ // 9 elements
	0x00000000, 0x00000010, 0x00000021, 0x00000036,
	0x0000004c, 0x0000006c, 0x0000007c, 0x0000008e,
	0x00000093,
} // Size: 60 bytes

const enData string = "" + // Size: 147 bytes
	"\x02Choose language\x02Language changed\x02Settings are changed\x02Wrong" +
	" callback input!\x04\x00\x02\x0a\x0a\x1a\x02Hello <b>%[1]s %[2]s.</b>" +
	"\x02Unknown command\x02🇬🇧 Language\x02Done"

var ruIndex = []uint32{ // 9 elements
	0x00000000, 0x0000001a, 0x00000032, 0x00000056,
	0x00000070, 0x00000097, 0x000000bd, 0x000000cf,
	0x000000dc,
} // Size: 60 bytes

const ruData string = "" + // Size: 220 bytes
	"\x02Выберите язык\x02Язык изменен\x02Настройки изменены\x02Неверный ввод" +
	"\x04\x00\x02\x0a\x0a!\x02Привет <b>%[1]s %[2]s.</b>\x02Неизвестная коман" +
	"да\x02🇷🇺 Язык\x02Готово"

var uzIndex = []uint32{ // 9 elements
	0x00000000, 0x0000000e, 0x00000020, 0x00000038,
	0x0000004a, 0x0000006a, 0x0000007a, 0x00000087,
	0x0000008e,
} // Size: 60 bytes

const uzData string = "" + // Size: 142 bytes
	"\x02Tilni tanlang\x02Til o'zgartirildi\x02Sozlamalar ozgartirildi\x02Not" +
	"og'ri kiritish\x04\x00\x02\x0a\x0a\x1a\x02Salom <b>%[1]s %[2]s.</b>\x02N" +
	"omalum komanda\x02🇺🇿 Til\x02Tayyor"

	// Total table size 689 bytes (0KiB); checksum: DECDCB1E
