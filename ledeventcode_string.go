// Code generated by "stringer -type=LEDEventCode"; DO NOT EDIT.

package golibevdev

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LedNumLock-0]
	_ = x[LedCapsLock-1]
	_ = x[LedScrollLock-2]
	_ = x[LedCompose-3]
	_ = x[LedKana-4]
	_ = x[LedSleep-5]
	_ = x[LedSuspend-6]
	_ = x[LedMute-7]
	_ = x[LedMisc-8]
	_ = x[LedMail-9]
	_ = x[LedCharging-10]
	_ = x[LedMax-15]
	_ = x[LedCnt-16]
}

const (
	_LEDEventCode_name_0 = "LedNumLockLedCapsLockLedScrollLockLedComposeLedKanaLedSleepLedSuspendLedMuteLedMiscLedMailLedCharging"
	_LEDEventCode_name_1 = "LedMaxLedCnt"
)

var (
	_LEDEventCode_index_0 = [...]uint8{0, 10, 21, 34, 44, 51, 59, 69, 76, 83, 90, 101}
	_LEDEventCode_index_1 = [...]uint8{0, 6, 12}
)

func (i LEDEventCode) String() string {
	switch {
	case i <= 10:
		return _LEDEventCode_name_0[_LEDEventCode_index_0[i]:_LEDEventCode_index_0[i+1]]
	case 15 <= i && i <= 16:
		i -= 15
		return _LEDEventCode_name_1[_LEDEventCode_index_1[i]:_LEDEventCode_index_1[i+1]]
	default:
		return "LEDEventCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
