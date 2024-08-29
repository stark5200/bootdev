# Challenge 2 chapter 1, main_test.py

def hex_to_rgb(hex_color):
    if not is_hexadecimal(hex_color) or len(hex_color) != 6:
        raise Exception("not a hex color string")
    
    return int(hex_color[:2], 16), int(hex_color[2:4], 16), int(hex_color[4:], 16)


def is_hexadecimal(hex_string):
    try:
        int(hex_string, 16)
        return True
    except Exception:
        return False
      
# Challenge 3 chapter 1, main_test1.py

def deduplicate_lists(lst1, lst2, reverse=False):
    lst3 = []
    lst3 = add_unique(lst3, lst1)
    lst3 = add_unique(lst3, lst2)
    
    return sorted(lst3, reverse = reverse)

def is_empty(lst):
    if len(lst) == 0:
        return True
    return False

def add_unique(old_lst, new_lst):
    if is_empty(new_lst):
        return old_lst
    lst = old_lst
    for i in new_lst:
        if i not in lst:
            lst.append(i)
    return lst
    

