// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"
)

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//go:nosplit
func sysAlloc(n uintptr, sysStat *sysMemStat) unsafe.Pointer {
	v, err := mmap(nil, n, _PROT_READ|_PROT_WRITE, _MAP_ANON|_MAP_PRIVATE, -1, 0)
	if err != 0 {
		return nil
	}
	sysStat.add(int64(n))
	return v
}

func sysUnused(v unsafe.Pointer, n uintptr) {
}

func sysUsed(v unsafe.Pointer, n uintptr) {
}

func sysHugePage(v unsafe.Pointer, n uintptr) {
}

// Don't split the stack as this function may be invoked without a valid G,
// which prevents us from allocating more stack.
//go:nosplit
func sysFree(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	sysStat.add(-int64(n))
	munmap(v, n)
}

func sysFault(v unsafe.Pointer, n uintptr) {
}

func sysReserve(v unsafe.Pointer, n uintptr) unsafe.Pointer {
	if v != nil {
		return nil
	}
	p, err := mmap(nil, n, _PROT_READ|_PROT_WRITE, _MAP_ANON|_MAP_PRIVATE, -1, 0)
	if err != 0 {
		return nil
	}
	return p
}

const _ENOMEM = 12

func sysMap(v unsafe.Pointer, n uintptr, sysStat *sysMemStat) {
	sysStat.add(int64(n))
}
