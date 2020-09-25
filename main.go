package main

// #include <string.h>
// #include <stdbool.h>
// #include <mysql.h>
// #cgo CFLAGS: -O3 -I/usr/include/mysql -fno-omit-frame-pointer
import "C"
import (
	"os"
	"unsafe"
)

// main function is needed even for generating shared object files
func main() {}

func msg(message *C.char, s string) {
	m := C.CString(s)
	defer C.free(unsafe.Pointer(m))

	C.strcpy(message, m)
}

//export Setenv_init
func Setenv_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 2 {
		msg(message, "`setenv`() requires 2 parameters: the environment variable name, and the value to be stored")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	argsTypes[1] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export Setenv
func Setenv(initid *C.UDF_INIT, args *C.UDF_ARGS, isNull *C.char, isError *C.char) C.longlong {
	c := 2
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]
	argsLengths := (*[1 << 30]uint64)(unsafe.Pointer(args.lengths))[:c:c]

	if argsArgs[0] == nil ||
		argsArgs[1] == nil {
		*isNull = 1
		return 0
	}

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))
	}

	os.Setenv(a[0], a[1])

	*isNull = 1
	return 0
}

//export Getenv_init
func Getenv_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 1 {
		msg(message, "`getenv` requires 1 parameter: the environment variable name")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export Getenv
func Getenv(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char, length *uint64, isNull *C.char, message *C.char) *C.char {
	c := 1
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]
	argsLengths := (*[1 << 30]uint64)(unsafe.Pointer(args.lengths))[:c:c]

	*length = 0
	*isNull = 1
	if argsArgs[0] == nil {
		return nil
	}

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))
	}

	s := os.Getenv(a[0])

	*length = uint64(len(s))
	*isNull = 0
	return C.CString(s)
}

//export Unsetenv_init
func Unsetenv_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.bool {
	if args.arg_count != 1 {
		msg(message, "`unsetenv`() requires 1 parameter: the environment variable name")
		return C.bool(true)
	}

	argsTypes := (*[2]uint32)(unsafe.Pointer(args.arg_type))

	argsTypes[0] = C.STRING_RESULT
	initid.maybe_null = 1

	return C.bool(false)
}

//export Unsetenv
func Unsetenv(initid *C.UDF_INIT, args *C.UDF_ARGS, isNull *C.char, isError *C.char) C.longlong {
	c := 1
	argsArgs := (*[1 << 30]*C.char)(unsafe.Pointer(args.args))[:c:c]
	argsLengths := (*[1 << 30]uint64)(unsafe.Pointer(args.lengths))[:c:c]

	if argsArgs[0] == nil {
		*isNull = 1
		return 0
	}

	a := make([]string, c, c)
	for i, argsArg := range argsArgs {
		a[i] = C.GoStringN(argsArg, C.int(argsLengths[i]))
	}

	os.Unsetenv(a[0])

	*isNull = 1
	return 0
}
