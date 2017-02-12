package benchmarks

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

type WithInterface struct {
	something interface{}
}

func TestWithInterface(t *testing.T) {
	var x WithInterface
	x.something = "hat"
	fmt.Printf("x is %#v\n", x)
	x.something = 1337
	fmt.Printf("x is %#v\n", x)

	typ := reflect.TypeOf(x)
	structField := typ.Field(0)
	typ = structField.Type
	fmt.Printf("Align=%d\n", typ.Align())
	fmt.Printf("FieldAlign=%d\n", typ.FieldAlign())
	fmt.Printf("Size=%d\n", typ.Size())
}

func TestLookAtPointer(t *testing.T) {
	var val int = 1337
	valptr := &val

	// Store a the pointer to our int in an interface type
	var store interface{} = valptr

	// Print out where the int is stored in memory,
	// and confirm the contents is as we expect.
	fmt.Printf("int is at %p\n", valptr)
	pointer := uintptr(unsafe.Pointer(valptr))
	contents := *(*[8]byte)(unsafe.Pointer(pointer))
	fmt.Printf("contents is %x (1337 in hex is %x)\n", contents, 1337)

	// Print out where the interface{} is stored in memory,
	// and look at the content
	pointer = uintptr(unsafe.Pointer(&store))
	fmt.Printf("interface{} is at %p\n", unsafe.Pointer(pointer))
	ifcontents := *(*[16]byte)(unsafe.Pointer(pointer))
	fmt.Printf("contents is %x\n", ifcontents)

	// The interface{} looks like two pointers
	pointers := *(*[2]uintptr)(unsafe.Pointer(&store))
	fmt.Printf("1st pointer is %x\n", pointers[0])
	fmt.Printf("2nd pointer is %x\n", pointers[1])

	// The second pointer points to our integer
	// We guess the initial pointer points to type information
	typ := reflect.TypeOf(valptr)
	pointers = *(*[2]uintptr)(unsafe.Pointer(&typ))
	fmt.Printf("Type 1st pointer is %x\n", pointers[0])
	fmt.Printf("Type 2nd pointer is %x\n", pointers[1])
}

func TestLookAtImmediate(t *testing.T) {
	var val int = 1337
	var store interface{} = val

	pointer := uintptr(unsafe.Pointer(&store))
	fmt.Printf("interface{} is at %p\n", unsafe.Pointer(pointer))

	// The interface{} looks like two pointers
	pointers := *(*[2]uintptr)(unsafe.Pointer(&store))
	fmt.Printf("1st pointer is %x\n", pointers[0])
	fmt.Printf("2nd pointer is %x\n", pointers[1])

	valcontents := *(*[8]byte)(unsafe.Pointer(pointers[1]))
	fmt.Printf("value contents is %x\n", valcontents)

	// The second pointer points to our integer
	// We guess the initial pointer points to type information
	typ := reflect.TypeOf(val)
	pointers = *(*[2]uintptr)(unsafe.Pointer(&typ))
	fmt.Printf("Type 1st pointer is %x\n", pointers[0])
	fmt.Printf("Type 2nd pointer is %x\n", pointers[1])
}

type benchType interface {
	get() int
}

type typeA struct {
	a int
}

func (a *typeA) get() int {
	return a.a
}

type typeB struct {
	b int
}

func (b *typeB) get() int {
	return b.b
}

func BenchmarkInterfaceTypeAssert(b *testing.B) {
	var x [2]interface{}

	x[0] = &typeA{}
	x[1] = &typeB{}

	b.ReportAllocs()
	b.ResetTimer()

	count := 0
	for i := 0; i < b.N; i++ {
		y := x[i&1]
		if _, ok := y.(*typeA); ok {
			count++
		}
	}

	b.Logf("count is %d", count)
}

func BenchmarkInterfaceTypeAssertInterface(b *testing.B) {
	var x [2]interface{}

	x[0] = &typeA{}
	x[1] = &typeB{}

	b.ReportAllocs()
	b.ResetTimer()

	count := 0
	for i := 0; i < b.N; i++ {
		y := x[i&1]
		if _, ok := y.(benchType); ok {
			count++
		}
	}

	b.Logf("count is %d", count)
}

func BenchmarkInterfaceTypeSwitch(b *testing.B) {
	var x [2]interface{}

	x[0] = &typeA{}
	x[1] = &typeB{}

	b.ReportAllocs()
	b.ResetTimer()

	count := 0
	for i := 0; i < b.N; i++ {

		y := x[i&1]

		switch y := y.(type) {
		case *typeA:
			count += y.get()
		case *typeB:
			count += y.get()
		case benchType:
			count += y.get()
		}
	}

	b.Logf("count is %d", count)
}

func BenchmarkInterfaceCall(b *testing.B) {
	var x [2]benchType
	x[0] = &typeA{}
	x[1] = &typeB{}
	b.ReportAllocs()
	b.ResetTimer()

	total := 0
	for i := 0; i < b.N; i++ {
		total += x[i&1].get()
	}

	if total > 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkInterfaceCallComparison(b *testing.B) {
	x := &typeA{}
	b.ReportAllocs()
	b.ResetTimer()

	total := 0
	for i := 0; i < b.N; i++ {
		total += x.get()
	}

	if total > 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkInterfaceCallTypeAssertion(b *testing.B) {
	var x [2]benchType
	x[0] = &typeA{}
	x[1] = &typeB{}
	b.ReportAllocs()
	b.ResetTimer()

	total := 0
	for i := 0; i < b.N; i++ {
		switch x := x[i&1].(type) {
		case *typeA:
			total += x.get()
		case *typeB:
			total += x.get()
		}
	}

	if total > 0 {
		b.Logf("total is %d", total)
	}
}

func BenchmarkInterfaceStore(b *testing.B) {
	var store interface{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		store = 1337
	}

	_ = store
}

func BenchmarkInterfaceStorePointer(b *testing.B) {
	var store interface{}
	val := 7
	valptr := &val

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		store = valptr
	}

	_ = store
}
