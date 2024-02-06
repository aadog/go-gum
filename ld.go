package gum

/*
#cgo LDFLAGS: -lfrida-gum -ldl
#cgo CFLAGS: -I/usr/local/include/ -w
#cgo CXXFLAGS: -std=c++11
#cgo darwin LDFLAGS: -lbsm -framework Foundation -framework AppKit -lresolv -lpthread
#cgo android LDFLAGS: -llog
#cgo android CFLAGS: -DANDROID -Wno-error=incompatible-function-pointer-types
#cgo linux,!android LDFLAGS: -lpthread
#cgo linux CFLAGS: -pthread
#include <frida-gum.h>
*/
import "C"
