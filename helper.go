package golibevdev

// #include <string.h>
import "C"

func strError(rc C.int) string {
	return C.GoString(C.strerror(-rc))
}
