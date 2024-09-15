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
func Register(connection C.duckdb_connection, info C.duckdb_extension_info) {
	if err := registerType(connection, "go_defined_type"); err != nil {
		panic(fmt.Sprintf("failed to register extension: %s, need better error handling", err.Error()))
	}
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}

func registerType(connection C.duckdb_connection, name string) error {

	typeName := C.CString(name)
	defer C.free(unsafe.Pointer(typeName))

	logicalType := C.duckdb_create_logical_type(C.DUCKDB_TYPE_INTEGER)
	defer C.duckdb_destroy_logical_type(&logicalType)

	C.duckdb_logical_type_set_alias(logicalType, typeName)

	status := C.duckdb_register_logical_type(connection, logicalType, nil)
	if status != C.DuckDBSuccess {
		return fmt.Errorf("failed to register type %s", name)
	}

	return nil

}
