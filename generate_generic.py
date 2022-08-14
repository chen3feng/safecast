ALL_INT_BITS = ('', '8', '16', '32', '64')


def ToCamelCase(type):
    return type[0].upper() + type[1:]


def generate_header():
    print('package safecast\n')
    print("""
type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Number interface {
	Integer | Float
}"""
          )


def generate_to_function():
    """Generate int conversion functions."""

    print("""
func To[T Number, F Number](value F)(to T, ok bool) {
    ok = true
    switch t := any(to).(type) {""")
    for to_type in ('int', 'uint'):
        for to_bits in ALL_INT_BITS:
            print(f'\tcase {to_type}{to_bits}:')
            print(f'\t\tt, ok = to{ToCamelCase(to_type)}{to_bits}(value)')
            print(f'\t\tto = T(t)')
    for to_bits in (32, 64):
        print(f'\tcase float{to_bits}:')
        print(f'\t\tt, ok = toFloat{to_bits}(value)')
        print(f'\t\tto = T(t)')

    print('\t}\n\treturn to, ok')
    print('}\n')

    for to_type in ('int', 'uint'):
        for to_bits in ALL_INT_BITS:
            generate_to_type(to_type, to_bits)
    for to_bits in (32, 64):
        generate_to_type('float', to_bits)


def generate_to_type(to_type, to_bits):
    full_to_type = f'{to_type}{to_bits}'
    print(f'func to{ToCamelCase(to_type)}{to_bits}[F Number](value F) ({full_to_type}, bool) {{')
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
            f'\tcase {full_from_type}: return {ToCamelCase(full_from_type)}To{ToCamelCase(full_to_type)}(f)')


def main():
    generate_header()
    generate_to_function()


main()
