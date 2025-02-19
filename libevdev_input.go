package golibevdev

// #cgo pkg-config: libevdev
// #include <string.h>
// #include <stdlib.h>
// #include <libevdev/libevdev.h>
// #include <libevdev/libevdev-uinput.h>
import "C"
import (
	"fmt"
	"unsafe"
)

type UInputDev struct {
	cStruct *C.struct_libevdev_uinput
	dev     *InputDev
}

func NewVirtualKeyboard(name string) (uid *UInputDev, err error) {
	dev := C.libevdev_new()
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.libevdev_set_name(dev, n)

	input := &InputDev{cStruct: dev}
	defer func() {
		if err != nil {
			input.Close()
		}
	}()

	if err = input.EnableEventType(EvKey); err != nil {
		return
	}

	for k := KeyReserved; k <= KeyMax; k++ {
		if err = input.EnableEventCode(EvKey, k); err != nil {
			return
		}
	}

	var uidev *C.struct_libevdev_uinput
	rc := C.libevdev_uinput_create_from_device(dev, C.LIBEVDEV_UINPUT_OPEN_MANAGED, &uidev)
	if rc != 0 {
		C.libevdev_free(dev)
		return nil, fmt.Errorf("failed to create uinput device: %v", strError(rc))
	}

	return &UInputDev{
		cStruct: uidev,
		dev:     input,
	}, nil
}

func NewVirtualMouse(name string) (uid *UInputDev, err error) {
	dev := C.libevdev_new()
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.libevdev_set_name(dev, n)

	input := &InputDev{cStruct: dev}
	defer func() {
		if err != nil {
			input.Close()
		}
	}()

	// Enable necessary event types for mouse
	if err = input.EnableEventType(EvKey); err != nil {
		return
	}
	if err = input.EnableEventType(EvRel); err != nil {
		return
	}

	// Enable mouse buttons
	mouseButtons := []EventCode{
		BtnLeft,
		BtnRight,
		BtnMiddle,
	}
	for _, btn := range mouseButtons {
		if err = input.EnableEventCode(EvKey, btn); err != nil {
			return
		}
	}

	// Enable relative axes
	relAxes := []EventCode{
		RelX,
		RelY,
		RelWheel,
	}
	for _, axis := range relAxes {
		if err = input.EnableEventCode(EvRel, axis); err != nil {
			return
		}
	}

	var uidev *C.struct_libevdev_uinput
	rc := C.libevdev_uinput_create_from_device(dev, C.LIBEVDEV_UINPUT_OPEN_MANAGED, &uidev)
	if rc != 0 {
		C.libevdev_free(dev)
		return nil, fmt.Errorf("failed to create uinput device: %v", strError(rc))
	}

	return &UInputDev{
		cStruct: uidev,
		dev:     input,
	}, nil
}

func NewVirtualTouchpad(name string) (uid *UInputDev, err error) {
	dev := C.libevdev_new()
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.libevdev_set_name(dev, n)

	input := &InputDev{cStruct: dev}
	defer func() {
		if err != nil {
			input.Close()
		}
	}()

	// Enable necessary event types for touchpad
	if err = input.EnableEventType(EvKey); err != nil {
		return
	}
	if err = input.EnableEventType(EvAbs); err != nil {
		return
	}

	// Enable touchpad buttons
	touchpadButtons := []EventCode{
		BtnTouch,
		BtnToolFinger,
		BtnLeft,
		BtnRight,
	}
	for _, btn := range touchpadButtons {
		if err = input.EnableEventCode(EvKey, btn); err != nil {
			return
		}
	}

	// Enable absolute axes with proper ranges
	absAxes := []struct {
		code       EventCode
		min, max   int32
		resolution int32
	}{
		{AbsX, 0, 3000, 30},
		{AbsY, 0, 2000, 30},
		{AbsPressure, 0, 255, 0},
	}
	for _, axis := range absAxes {
		if err = input.EnableEventCode(EvAbs, axis.code); err != nil {
			return
		}
		C.libevdev_set_abs_minimum(dev, C.uint(axis.code.Value()), C.int(axis.min))
		C.libevdev_set_abs_maximum(dev, C.uint(axis.code.Value()), C.int(axis.max))
		if axis.resolution > 0 {
			C.libevdev_set_abs_resolution(dev, C.uint(axis.code.Value()), C.int(axis.resolution))
		}
	}

	var uidev *C.struct_libevdev_uinput
	rc := C.libevdev_uinput_create_from_device(dev, C.LIBEVDEV_UINPUT_OPEN_MANAGED, &uidev)
	if rc != 0 {
		C.libevdev_free(dev)
		return nil, fmt.Errorf("failed to create uinput device: %v", strError(rc))
	}

	return &UInputDev{
		cStruct: uidev,
		dev:     input,
	}, nil
}

func NewVirtualGamepad(name string) (uid *UInputDev, err error) {
	dev := C.libevdev_new()
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	C.libevdev_set_name(dev, n)

	input := &InputDev{cStruct: dev}
	defer func() {
		if err != nil {
			input.Close()
		}
	}()

	// Enable necessary event types for gamepad
	if err = input.EnableEventType(EvKey); err != nil {
		return
	}
	if err = input.EnableEventType(EvAbs); err != nil {
		return
	}

	// Enable gamepad buttons
	gamepadButtons := []EventCode{
		BtnA,
		BtnB,
		BtnX,
		BtnY,
		BtnTL,
		BtnTR,
		BtnSelect,
		BtnStart,
	}
	for _, btn := range gamepadButtons {
		if err = input.EnableEventCode(EvKey, btn); err != nil {
			return
		}
	}

	// Enable absolute axes with proper ranges for analog sticks
	absAxes := []struct {
		code       EventCode
		min, max   int32
		flat, fuzz int32
	}{
		{AbsX, -32768, 32767, 128, 16},     // Left stick X
		{AbsY, -32768, 32767, 128, 16},     // Left stick Y
		{AbsRx, -32768, 32767, 128, 16},    // Right stick X
		{AbsRy, -32768, 32767, 128, 16},    // Right stick Y
		{AbsZ, 0, 255, 0, 0},               // Left trigger
		{AbsRz, 0, 255, 0, 0},              // Right trigger
	}
	for _, axis := range absAxes {
		if err = input.EnableEventCode(EvAbs, axis.code); err != nil {
			return
		}
		C.libevdev_set_abs_minimum(dev, C.uint(axis.code.Value()), C.int(axis.min))
		C.libevdev_set_abs_maximum(dev, C.uint(axis.code.Value()), C.int(axis.max))
		C.libevdev_set_abs_flat(dev, C.uint(axis.code.Value()), C.int(axis.flat))
		C.libevdev_set_abs_fuzz(dev, C.uint(axis.code.Value()), C.int(axis.fuzz))
	}

	var uidev *C.struct_libevdev_uinput
	rc := C.libevdev_uinput_create_from_device(dev, C.LIBEVDEV_UINPUT_OPEN_MANAGED, &uidev)
	if rc != 0 {
		C.libevdev_free(dev)
		return nil, fmt.Errorf("failed to create uinput device: %v", strError(rc))
	}

	return &UInputDev{
		cStruct: uidev,
		dev:     input,
	}, nil
}

func (dev *UInputDev) Close() {
	C.libevdev_uinput_destroy(dev.cStruct)
}

func (dev *UInputDev) WriteEvent(typ EventType, code EventCode, value int32) error {
	rc := C.libevdev_uinput_write_event(dev.cStruct, C.uint(typ), C.uint(code.Value()), C.int(value))
	if rc != 0 {
		return fmt.Errorf("failed to write event: %v", strError(rc))
	}
	return nil
}

// Sync sends a sync event
func (dev *UInputDev) Sync() error {
	return dev.WriteEvent(EvSyn, SynReport, 0)
}

// GetName returns the name of the device
func (dev *UInputDev) GetName() string {
	return C.GoString(C.libevdev_get_name(dev.dev.cStruct))
}

// DeviceInfo contains basic information about an input device
type DeviceInfo struct {
	Name       string
	ID         struct{ Vendor, Product, Version, BusType uint16 }
	DriverVersion int
}

// GetDeviceInfo returns information about the device
func (dev *UInputDev) GetDeviceInfo() DeviceInfo {
	info := DeviceInfo{}
	info.Name = C.GoString(C.libevdev_get_name(dev.dev.cStruct))
	info.ID.Vendor = uint16(C.libevdev_get_id_vendor(dev.dev.cStruct))
	info.ID.Product = uint16(C.libevdev_get_id_product(dev.dev.cStruct))
	info.ID.Version = uint16(C.libevdev_get_id_version(dev.dev.cStruct))
	info.ID.BusType = uint16(C.libevdev_get_id_bustype(dev.dev.cStruct))
	info.DriverVersion = int(C.libevdev_get_driver_version(dev.dev.cStruct))
	return info
}

// ReadEvent reads the next event from the device
func (dev *UInputDev) ReadEvent(flag ReadFlag) (Event, error) {
	return dev.dev.NextEvent(flag)
}
