// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a Apache-style
// license that can be found in the LICENSE file.

package goxtp

/*
#cgo darwin LDFLAGS: -fPIC -L/usr/local/ -L. -L${SRCDIR}/api/macosx/  -Wl,-rpath,${SRCDIR}/api/macosx/ -lxtpquoteapi -lxtptraderapi -lstdc++
#cgo darwin CPPFLAGS: -fPIC -I. -I/${SRCDIR}/api/include/
#cgo linux LDFLAGS: -fPIC -L. -L${SRCDIR}/api/linux -Wl,-rpath=${SRCDIR}/api/linux/  -lxtpquoteapi -lxtptraderapi -lstdc++
#cgo linux CPPFLAGS: -fPIC -I. -I/${SRCDIR}/api/include/
*/
import "C"
