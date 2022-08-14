#!/usr/bin/env python3

from safecast import ALL_INT_BITS, to_camel_case


def generate_header():
    print('package safecast\n')
    print('import "math"\n')


def generate(from_type, from_bits, to_type, to_bits):
    full_from_type = f'{from_type}{from_bits}'
    full_to_type = f'{to_type}{to_bits}'

    def print_false_return(cond):
        print(f'\tif {cond} {{\n\t\treturn {full_to_type}(value), false\n\t}}')

    overflow = f'{full_from_type}({full_to_type}(value)) != value'

    def same_signedness():
        if from_bits and to_bits:  # intnn to intxx
            if int(from_bits) <= int(to_bits):
                return
            print_false_return(overflow)
        elif from_bits:  # u?intnn to u?int
            if int(from_bits) <= 32:  # int is at least 32 bits
                return
            print_false_return(f'math.MaxInt == math.MaxInt32 && {overflow}')
        elif to_bits:  # u?int to u?intnn
            if int(to_bits) > 32:  # int64 >= int, ok
                return
            # int to intnn
            cond = ('math.MaxInt == math.MaxInt64 && ' if int(
                to_bits) == 32 else '') + overflow
            print_false_return(cond)

    def uint_to_int():
        if from_bits and to_bits:  # uintnn to intxx
            if int(from_bits) < int(to_bits):
                return
            if int(from_bits) > int(to_bits):
                print_false_return(overflow)
            else:
                cond = f'value > math.Max{to_camel_case(full_to_type)}'
                print_false_return(cond)
        elif from_bits:  # uintnn to int
            if int(from_bits) < 32:
                return
            print_false_return(f'value > math.MaxInt{from_bits}')
        elif to_bits:  # uint_to_intnn
            if int(to_bits) > 32:
                return
            print_false_return(f'value > math.MaxInt{from_bits}')
        else:  # uint to int
            print_false_return('value > math.MaxInt')

    def int_to_uint():
        print_false_return('value < 0')
        if from_bits and to_bits:  # intnn to uintxx
            if int(from_bits) <= int(to_bits):
                return
            print_false_return(overflow)
        elif from_bits:  # intnn to uint
            if int(from_bits) <= 32:
                return
            #  int32/64 to uint
            print_false_return(
                'math.MaxInt == math.MaxInt32 && value > math.MaxUint32')
        elif to_bits:  # int to uintnn
            if int(to_bits) > 32:
                return
            print_false_return(f'value > math.MaxUint{to_bits}')
        else:  # int to uint
            # Always OK
            pass

    def generate_range_chack():
        if from_type == to_type:
            same_signedness()
        elif from_type == 'uint':
            uint_to_int()
        else:  # int to uint
            int_to_uint()

    func_name = f'{full_from_type}To{to_camel_case(full_to_type)}'
    print(f'// {func_name} converts the {full_from_type} value to {full_to_type} safely.')
    print(f'func {func_name}(value {full_from_type}) (result {full_to_type}, ok bool) {{')

    generate_range_chack()
    print(f'\treturn {full_to_type}(value), true')
    print('}\n')


def generate_int_convs():
    """Generate int conversion functions."""
    for from_type in ('int', 'uint'):
        for from_bit in ALL_INT_BITS:
            for to_type in ('int', 'uint'):
                for to_bits in ALL_INT_BITS:
                    if from_type == to_type and from_bit == to_bits:
                        continue
                    generate(from_type, from_bit, to_type, to_bits)


def generate_float_convs():
    """Generate int conversion functions."""
    for float_bit in (32, 64):
        for int_type in ('int', 'uint'):
            for int_bits in ALL_INT_BITS:
                generate_float_to_conv(float_bit, int_type, int_bits)
                generate_to_float_conv(int_type, int_bits, float_bit)

    print("""
    func float32ToFloat64(value float32) (float64, bool) {
        return float64(value), true
    }

    func float64ToFloat32(value float64) (float32, bool) {
        if value > math.MaxFloat32 || value < -math.MaxFloat32 {
            return float32(value), false
        }
        return float32(value), true
    }
    """)


def generate_float_to_conv(from_bits, to_type, to_bits):
    full_from_type = f'float{from_bits}'
    full_to_type = f'{to_type}{to_bits}'

    def print_false_return(cond):
        print(f'\tif {cond} {{\n\t\treturn {full_to_type}(value), false\n\t}}')

    def generate_range_chack():
        min = f'{full_from_type}(math.Min{to_camel_case(full_to_type)})' if to_type == 'int' else '0'
        cond = f'value < {min}'
        cond += f' || value > {full_from_type}(math.Max{to_camel_case(full_to_type)})'
        print_false_return(cond)

    func_name = f'float{from_bits}To{to_camel_case(to_type)}{to_bits}'
    print(f'// {func_name} converts the {full_from_type} value to {full_to_type} safely.')
    print(f'func {func_name}(value {full_from_type}) (result {full_to_type}, ok bool) {{')
    generate_range_chack()
    print(f'\treturn {full_to_type}(value), true')
    print('}\n')


def generate_to_float_conv(from_type, from_bits, to_bits):
    full_from_type = f'{from_type}{from_bits}'
    full_to_type = f'float{to_bits}'
    print(f'func {full_from_type}To{to_camel_case(full_to_type)}(value {full_from_type}) ({full_to_type}, bool) {{')
    print(f'\treturn {full_to_type}(value), true')
    print('}')


def main():
    generate_header()
    generate_int_convs()
    generate_float_convs()


main()
