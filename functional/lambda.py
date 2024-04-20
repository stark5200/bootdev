type = {
    ".txt": "Text", 
    ".docx": "Document", 
    ".py": "Code"
}

def categorize_file(filename):
    get_category = lambda extension: type.get(extension, "Unknown")
    return get_category(filename[filename.rfind(".") :])

# first class function example
def square(x):
    return x * x

# Assign function to a variable
f = square

print(f(5))
# 25


# Higher order example
def square(x):
    return x * x

def my_map(func, arg_list):
    result = []
    for i in arg_list:
        result.append(func(i))
    return result

squares = my_map(square, [1, 2, 3, 4, 5])
print(squares)
# [1, 4, 9, 16, 25]

# With map, we can operate on lists without using loops and nasty stateful variables. For example:

def square(x):
    return x * x

nums = [1, 2, 3, 4, 5]
squared_nums = map(square, nums)
print(list(squared_nums))
print(squared_nums) # prints location if not stated that it it in fact a list
# [1, 4, 9, 16, 25]

# map
def change_bullet_style(document):
    return "\n".join(map(convert_line, list(document.split("\n"))))


# Don't edit below this line


def convert_line(line):
    old_bullet = "-"
    new_bullet = "*"
    if len(line) > 0 and line[0] == old_bullet:
        return new_bullet + line[1:]
    return line
