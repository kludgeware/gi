// Code generated by "stringer -type=FontStretch"; DO NOT EDIT.

package gi

import (
	"fmt"
	"strconv"
)

const _FontStretch_name = "FontStrNormalFontStrUltraCondensedFontStrExtraCondensedFontStrSemiCondensedFontStrSemiExpandedFontStrExtraExpandedFontStrUltraExpandedFontStrCondensedFontStrExpandedFontStrNarrowerFontStrWiderFontStretchN"

var _FontStretch_index = [...]uint8{0, 13, 34, 55, 75, 94, 114, 134, 150, 165, 180, 192, 204}

func (i FontStretch) String() string {
	if i < 0 || i >= FontStretch(len(_FontStretch_index)-1) {
		return "FontStretch(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FontStretch_name[_FontStretch_index[i]:_FontStretch_index[i+1]]
}

func (i *FontStretch) FromString(s string) error {
	for j := 0; j < len(_FontStretch_index)-1; j++ {
		if s == _FontStretch_name[_FontStretch_index[j]:_FontStretch_index[j+1]] {
			*i = FontStretch(j)
			return nil
		}
	}
	return fmt.Errorf("String %v is not a valid option for type FontStretch", s)
}
