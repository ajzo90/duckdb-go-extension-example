package main

/*
#include <stdlib.h>
#include <duckdb.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

//export Register
func Register(connection C.duckdb_connection) int {

	if registerType(connection, "go_defined_type") == nil {
	} else {
		panic("failed to register extension, need better error handling")
	}
	return 0
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}

func registerType(connection C.duckdb_connection, name string) error {

	typeName := C.CString(name)
	defer C.free(unsafe.Pointer(typeName))
	//
	logicalType := C.duckdb_create_logical_type(C.DUCKDB_TYPE_INTEGER)
	defer C.duckdb_destroy_logical_type(&logicalType)
	//
	C.duckdb_logical_type_set_alias(logicalType, typeName)
	//
	status := C.duckdb_register_logical_type(connection, logicalType, nil)
	//
	if status != C.DuckDBSuccess {
		return fmt.Errorf("failed to register type %s", name)
	}

	return nil

}
