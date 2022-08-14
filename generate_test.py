#!/usr/bin/env python3


from safecast import ALL_INT_BITS, to_camel_case


print('''
package safecast_test

import (
	_ "fmt"
	_ "math"
	"testing"

	"github.com/chen3feng/safecast"
)

func expectFalse(t * testing.T, value bool) {
    if value {
        t.Fail()
    }
}

func expectTrue(t * testing.T, value bool) {
    if !value {
        t.Fail()
    }
}
''')

for from_bit in ALL_INT_BITS:
    for to_bit in ALL_INT_BITS:
        to_type = f'uint{to_bit}'
        from_type = f'int{from_bit}'
        print(f'''
    func TestTo_{from_type}_to_{to_type}(t * testing.T) {{
        var i {from_type} = -1
        _, ok := safecast.To[{to_type}](i)
        expectFalse(t, ok)
        i = 1
        _, ok = safecast.To[{to_type}](i)
        expectTrue(t, ok)
    }}''')

for from_bit in ALL_INT_BITS:
    for to_bit in ALL_INT_BITS:
        to_type = f'int{to_bit}'
        from_type = f'int{from_bit}'
        print(f'''
    func TestTo_{from_type}_to_{to_type}(t * testing.T) {{
        var i {from_type} = -1
        _, ok := safecast.To[{to_type}](i)
        expectTrue(t, ok)
        i = 1
        _, ok = safecast.To[{to_type}](i)
        expectTrue(t, ok)
    }}''')

for from_bit in ALL_INT_BITS:
    for to_bit in ALL_INT_BITS:
        to_type = f'int{to_bit}'
        from_type = f'uint{from_bit}'
        print(f'''
    func TestTo_{from_type}_to_{to_type}(t * testing.T) {{
        var i {from_type} = 1
        _, ok := safecast.To[{to_type}](i)
        expectTrue(t, ok)
    }}''')
