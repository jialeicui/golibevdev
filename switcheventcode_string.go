// Code generated by "stringer -type=SwitchEventCode"; DO NOT EDIT.

package golibevdev

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SwLid-0]
	_ = x[SwTabletMode-1]
	_ = x[SwHeadphoneInsert-2]
	_ = x[SwRfkillAll-3]
	_ = x[SwMicroPhoneInsert-4]
	_ = x[SwDock-5]
	_ = x[SwLineoutInsert-6]
	_ = x[SwJackPhysicalInsert-7]
	_ = x[SwVideoOutInsert-8]
	_ = x[SwCameraLensCover-9]
	_ = x[SwKeypadSlide-10]
	_ = x[SwFrontProximity-11]
	_ = x[SwRotateLock-12]
	_ = x[SwLineinInsert-13]
	_ = x[SwMuteDevice-14]
	_ = x[SwPenInsert-15]
	_ = x[SwMax-15]
	_ = x[SwCnt-16]
}

const _SwitchEventCode_name = "SwLidSwTabletModeSwHeadphoneInsertSwRfkillAllSwMicroPhoneInsertSwDockSwLineoutInsertSwJackPhysicalInsertSwVideoOutInsertSwCameraLensCoverSwKeypadSlideSwFrontProximitySwRotateLockSwLineinInsertSwMuteDeviceSwPenInsertSwCnt"

var _SwitchEventCode_index = [...]uint8{0, 5, 17, 34, 45, 63, 69, 84, 104, 120, 137, 150, 166, 178, 192, 204, 215, 220}

func (i SwitchEventCode) String() string {
	if i >= SwitchEventCode(len(_SwitchEventCode_index)-1) {
		return "SwitchEventCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SwitchEventCode_name[_SwitchEventCode_index[i]:_SwitchEventCode_index[i+1]]
}
