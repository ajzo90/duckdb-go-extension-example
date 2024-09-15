# Example duckdb-plugin-templates for creating plugin for duckdb in Go.

## Example
```
 GEN=ninja make
 
 ./build/release/duckdb -unsigned 
 
```

```
D load './build/release/extension/quack/quack.duckdb_extension';

D select null::go_defined_type;

```
