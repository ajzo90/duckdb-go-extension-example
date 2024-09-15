package main

/*
#include <stdlib.h>
#include <duckdb.h>
*/
import "C"
import (
	"github.com/cespare/xxhash/v2"
	"github.com/marcboeker/go-duckdb"
	"github.com/marcboeker/go-duckdb/aggregates"
	"unsafe"
)

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}

//export Register
func Register(connection C.duckdb_connection, info C.duckdb_extension_info) {

	conn := duckdb.UpgradeConn((duckdb.Connection)(unsafe.Pointer(connection)))

	Check(duckdb.RegisterCast(conn, VarcharToTinyInt{}))

	Check(duckdb.RegisterType(conn, "vec3", "float[3]"))

	Check(duckdb.RegisterScalarUDFConn(conn, "xxhash_64", xxhash64{}))

	Check(duckdb.RegisterAggregateUDFConn[aggregates.ArraySumAggregateState](conn, "array_sum", aggregates.ArraySumAggregateFunc{}))

}

type VarcharToTinyInt struct {
}

func (v VarcharToTinyInt) Config() duckdb.CastFunctionConfig {
	return duckdb.CastFunctionConfig{
		Source: "VARCHAR",
		Target: "TINYINT",
	}
}

func (v VarcharToTinyInt) Exec(ctx *duckdb.CastExecContext) error {
	in := duckdb.GetData[duckdb.Varchar](ctx.Input())
	out := duckdb.GetData[int8](ctx.Output())

	for i := range ctx.Count() {
		out[i] = int8(len(in[0].Bytes()))
	}
	return nil
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

type xxhash64 struct {
}

func (xxh xxhash64) Config() duckdb.ScalarFunctionConfig {
	return duckdb.ScalarFunctionConfig{
		InputTypes: []string{duckdb.VARCHAR},
		ResultType: duckdb.UBIGINT,
	}
}

func (xxh xxhash64) Exec(ctx *duckdb.ExecContext) error {
	var chunkSize = ctx.ChunkSize()

	var a duckdb.Vec[duckdb.Varchar]
	_ = a.LoadCtx(ctx, 0, chunkSize)

	var out = duckdb.UDFScalarVectorResult[uint64](ctx)[:chunkSize]
	var in = a.Data[:chunkSize]

	for i, v := range in {
		out[i] = xxhash.Sum64(v.Bytes())
	}
	return nil
}
