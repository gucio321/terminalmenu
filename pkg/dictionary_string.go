// Code generated by "stringer -type DictionaryKey -linecomment -output dictionary_string.go"; DO NOT EDIT.

package terminalmenu

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DictionaryKeyBack-0]
	_ = x[DictionaryKeyExit-1]
	_ = x[DictionaryKeyAskForNumber-2]
	_ = x[DictionaryKeyEnterPrompt-3]
}

const _DictionaryKey_name = "BackExitWhat to do?:Press ENTER to continue..."

var _DictionaryKey_index = [...]uint8{0, 4, 8, 20, 46}

func (i DictionaryKey) String() string {
	if i >= DictionaryKey(len(_DictionaryKey_index)-1) {
		return "DictionaryKey(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DictionaryKey_name[_DictionaryKey_index[i]:_DictionaryKey_index[i+1]]
}
