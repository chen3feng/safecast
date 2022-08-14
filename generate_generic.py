#!/usr/bin/env python3

from safecast import ALL_INT_BITS, to_camel_case


def generate_header():
    print('package safecast\n')
    print("""
type numericType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64
}"""
          )


def generate_to_function():
    """Generate int conversion functions."""

    print("""
// To converts a numeric value from the FromType to the specified ToType type safely.
// result will always be same as the usual type cast (type(value)),
// but ok is false when overflow or underflow occured.
func To[ToType numericType, FromType numericType](value FromType)(result ToType, ok bool) {
    ok = true
    switch t := any(result).(type) {""")
    for to_type in ('int', 'uint'):
        for to_bits in ALL_INT_BITS:
            print(f'\tcase {to_type}{to_bits}:')
            print(f'\t\tt, ok = to{to_camel_case(to_type)}{to_bits}(value)')
            print(f'\t\tresult = ToType(t)')
    for to_bits in (32, 64):
        print(f'\tcase float{to_bits}:')
        print(f'\t\tt, ok = toFloat{to_bits}(value)')
        print(f'\t\tresult = ToType(t)')

    print('\t}\n\treturn result, ok')
    print('}\n')

    for to_type in ('int', 'uint'):
        for to_bits in ALL_INT_BITS:
            generate_to_type(to_type, to_bits)
    for to_bits in (32, 64):
        generate_to_type('float', to_bits)


def generate_to_type(to_type, to_bits):
    full_to_type = f'{to_type}{to_bits}'
    print(
        f'func to{to_camel_case(to_type)}{to_bits}[F numericType](value F) ({full_to_type}, bool) {{')
    print(f'\tswitch f := any(value).(type) {{')
    for from_type in ('int', 'uint'):
        for from_bits in ALL_INT_BITS:
            full_from_type = f'{from_type}{from_bits}'
            generate_call(full_from_type, full_to_type)

    for from_bits in (32, 64):
        full_from_type = f'float{from_bits}'
        generate_call(full_from_type, full_to_type)
    print(f'\t}}\n\treturn {full_to_type}(value), false')
    print('}\n')


def generate_call(full_from_type, full_to_type):
    if full_from_type == full_to_type:
        print(f'\tcase {full_from_type}: return f, true')
    else:
        print(
            f'\tcase {full_from_type}: return {full_from_type}To{to_camel_case(full_to_type)}(f)')


def main():
    generate_header()
    generate_to_function()


main()
