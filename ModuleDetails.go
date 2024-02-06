package gum

/*
#include <frida-gum.h>
*/
import "C"
import (
	"github.com/aadog/go-ffi"
	"unsafe"
)

type ModuleDetails struct {
	Name string
	Base ffi.NativePointer
	Size int64
	Path string
}

func (m *ModuleDetails) FindExportByName(SymbolName string) ffi.NativePointer {
	cModule_name := CString(m.Name)
	cSymbolName := CString(SymbolName)
	return ffi.Ptr(uintptr(C.gum_module_find_export_by_name(cModule_name, cSymbolName)))
}
func (m *ModuleDetails) FindSymbolByName(SymbolName string) ffi.NativePointer {
	cModule_name := CString(m.Name)
	cSymbolName := C.CString(SymbolName)
	return ffi.Ptr(uintptr(C.gum_module_find_symbol_by_name(cModule_name, cSymbolName)))
}

func (m *ModuleDetails) EnumerateSymbols() []*SymbolDetails {
	details := make([]*SymbolDetails, 0)
	var cb = ffi.NewNativeCallback(func(ptr ffi.NativePointer, ptr2 ffi.NativePointer) int8 {
		detail := SymbolDetailsWithNativePointer(ptr)
		details = append(details, detail)
		return 1
	}, ffi.Tint8, []ffi.ArgTypeName{ffi.TPointer, ffi.TPointer}).MakeCall().MustGet()
	C.gum_module_enumerate_symbols(CString(m.Name), (*[0]byte)(unsafe.Pointer(cb)), (C.gpointer)(unsafe.Pointer(uintptr(0))))
	return details
}
func (m *ModuleDetails) EnumerateExports() []*ExportDetails {
	details := make([]*ExportDetails, 0)
	var cb = ffi.NewNativeCallback(func(ptr ffi.NativePointer, ptr2 ffi.NativePointer) int {
		details = append(details, ExportDetailsWithPtr(ptr.Ptr()))
		return 1
	}, ffi.Tint, []ffi.ArgTypeName{ffi.TPointer, ffi.TPointer}).MakeCall().MustGet()
	C.gum_module_enumerate_exports(CString(m.Name), (*[0]byte)(unsafe.Pointer(cb)), (C.gpointer)(unsafe.Pointer(uintptr(0))))
	return details
}

func ModuleDetailsWithPtr(ptr unsafe.Pointer) *ModuleDetails {
	details := (*C.GumModuleDetails)(ptr)
	return &ModuleDetails{
		Name: GoString(details.name),
		Base: ffi.Ptr(uintptr(details._range.base_address)),
		Size: int64(details._range.size),
		Path: GoString(details.path),
	}
}
