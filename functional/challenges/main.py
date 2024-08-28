# Chapter 1 Challenge 2

def hex_to_rgb(hex_color):
    if is_hexadecimal(hex_color):
        return int(hex_color[:2], 16), int(hex_color[2:4], 16), int(hex_color[4:], 16)
    return Exception("not a hex color string")


def is_hexadecimal(hex_string):
    a = False
    if len(hex_string) == 6:
        a = True
    b = False
    if int(hex_string, 16):
        b = True

    try:
        a and b
        return True
    except:
        return Exception("not a hex color string")
