package golibevdev

// #cgo pkg-config: libevdev
// #include <libevdev/libevdev.h>
import "C"
import (
	"fmt"
	"os"
)

type InputDev struct {
	fd      *os.File
	name    string
	cStruct *C.struct_libevdev
}

func NewInputDev(path string) (*InputDev, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", path, err)
	}

	cStruct := C.libevdev_new()
	if cStruct == nil {
		return nil, fmt.Errorf("failed to create libevdev struct")
	}

	rc := C.libevdev_set_fd(cStruct, C.int(fd.Fd()))
	if rc != 0 {
		C.libevdev_free(cStruct)
		return nil, fmt.Errorf("failed to set fd: %v", strError(rc))
	}

	name := C.GoString(C.libevdev_get_name(cStruct))
	if name == "" {
		C.libevdev_free(cStruct)
		return nil, fmt.Errorf("failed to get device name")
	}

	return &InputDev{
		fd:      fd,
		name:    name,
		cStruct: cStruct,
	}, nil
}

func (dev *InputDev) Close() {
	C.libevdev_free(dev.cStruct)
	_ = dev.fd.Close()
}

func (dev *InputDev) Name() string {
	return dev.name
}

func (dev *InputDev) HasEventType(typ EventType) bool {
	return C.libevdev_has_event_type(dev.cStruct, C.uint(typ)) == 1
}

func (dev *InputDev) HasEventCode(typ EventType, code KeyEventCode) bool {
	return C.libevdev_has_event_code(dev.cStruct, C.uint(typ), C.uint(code)) == 1
}

type Event struct {
	Time  int64
	Type  EventType
	Code  EventCode
	Value int32
}

type ReadFlag int

const (
	ReadFlagSync      ReadFlag = 1
	ReadFlagNormal    ReadFlag = 2
	ReadFlagForceSync ReadFlag = 4
	ReadFlagGetRepeat ReadFlag = 8
)

func (dev *InputDev) NextEvent(flag ReadFlag) (Event, error) {
	var ev C.struct_input_event
	rc := C.libevdev_next_event(dev.cStruct, C.uint(flag), &ev)
	if rc != 0 {
		return Event{}, fmt.Errorf("failed to get next event: %v", strError(rc))
	}

	var code EventCode
	switch ev._type {
	case C.EV_KEY:
		code = KeyEventCode(ev.code)
	case C.EV_REL:
		code = RelativeEventCode(ev.code)
	case C.EV_SYN:
		code = SyncEventCode(ev.code)
	case C.EV_MSC:
		code = MiscEventCode(ev.code)
	case C.EV_SW:
		code = SwitchEventCode(ev.code)
	case C.EV_LED:
		code = LEDEventCode(ev.code)
	case C.EV_SND:
		code = SoundEventCode(ev.code)
	case C.EV_REP:
		code = RepeatEventCode(ev.code)
	case C.EV_FF:
		code = ForceEventCode(ev.code)
	case C.EV_FF_STATUS:
		code = FFStatusEventCode(ev.code)
	case C.EV_PWR:
		code = PowerEventCode(ev.code)
	default:
		code = UnknownEventCode(ev.code)
	}

	return Event{
		Time:  int64(ev.time.tv_sec),
		Type:  EventType(ev._type),
		Code:  code,
		Value: int32(ev.value),
	}, nil
}

func (dev *InputDev) Grab() error {
	rc := C.libevdev_grab(dev.cStruct, C.LIBEVDEV_GRAB)
	if rc != 0 {
		return fmt.Errorf("failed to grab device: %v", strError(rc))
	}
	return nil
}

func (dev *InputDev) EnableEventType(typ EventType) error {
	rc := C.libevdev_enable_event_type(dev.cStruct, C.uint(typ))
	if rc != 0 {
		return fmt.Errorf("failed to enable event type: %v", strError(rc))
	}
	return nil
}

func (dev *InputDev) EnableEventCode(typ EventType, code EventCode) error {
	rc := C.libevdev_enable_event_code(dev.cStruct, C.uint(typ), C.uint(code.Value()), nil)
	if rc != 0 {
		return fmt.Errorf("failed to enable event code: %v", strError(rc))
	}
	return nil
}
