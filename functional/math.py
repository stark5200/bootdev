def get_median_font_size(font_sizes):
    font_len = len(font_sizes)
    if font_len == 0:
        return None
    sorted_list = sorted(font_sizes)
    if font_len % 2 == 0:
        return ( sorted_list[font_len//2 - 1] + sorted_list[font_len//2] ) / 2
    return sorted_list[(font_len-1)//2]
    
