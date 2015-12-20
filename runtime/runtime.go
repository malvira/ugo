package runtime

import "unsafe" // needed for go:linkname

//go:linkname _rt0_arm_linux _rt0_arm_linux
func _rt0_arm_linux() {
	main()
}

func morestack_noctxt() {
}

func rt0_go() {
	mainPC()
}

func mainPC() {
	main()
}

//go:linkname main_main main.main
func main_main()

func emptyfunc() {
}

func morestack() {
}

func main() {
	var _ unsafe.Pointer // supress unused unsafe import
	main_main()
}


// moduledata records information about the layout of the executable
// image. It is written by the linker. Any changes here must be
// matched changes to the code in cmd/internal/ld/symtab.go:symtab.
// moduledata is stored in read-only memory; none of the pointers here
// are visible to the garbage collector.
type moduledata struct {
	pclntable    []byte
	ftab         []functab
	filetab      []uint32
	findfunctab  uintptr
	minpc, maxpc uintptr

	text, etext           uintptr
	noptrdata, enoptrdata uintptr
	data, edata           uintptr
	bss, ebss             uintptr
	noptrbss, enoptrbss   uintptr
	end, gcdata, gcbss    uintptr

	typelinks []*_type

	modulename   string
	modulehashes []modulehash

	gcdatamask, gcbssmask bitvector

	next *moduledata
}

type functab struct {
	entry   uintptr
	funcoff uintptr
}

// Needs to be in sync with ../cmd/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/internal/gc/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	_unused    uint8
	align      uint8
	fieldalign uint8
	kind       uint8
	alg        *typeAlg
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata  *byte
	_string *string
	x       *uncommontype
	ptrto   *_type
	zero    *byte // ptr to the zero value for this type
}

// typeAlg is also copied/used in reflect/type.go.
// keep them in sync.
type typeAlg struct {
	// function for hashing objects of this type
	// (ptr to object, seed) -> hash
	hash func(unsafe.Pointer, uintptr) uintptr
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
}

type method struct {
	name    *string
	pkgpath *string
	mtyp    *_type
	typ     *_type
	ifn     unsafe.Pointer
	tfn     unsafe.Pointer
}

type uncommontype struct {
	name    *string
	pkgpath *string
	mhdr    []method
}

// For each shared library a module links against, the linker creates an entry in the
// moduledata.modulehashes slice containing the name of the module, the abi hash seen
// at link time and a pointer to the runtime abi hash. These are checked in
// moduledataverify1 below.
type modulehash struct {
	modulename   string
	linktimehash string
	runtimehash  *string
}

// Information from the compiler about the layout of stack frames.
type bitvector struct {
	n        int32 // # of bits
	bytedata *uint8
}
