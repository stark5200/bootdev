# Pure functions have a lot of benefits. Whenever possible, good developers try to use pure functions instead of impure functions. Remember, pure functions:
# Return the same result if given the same arguments. They are deterministic.
# Do not change the external state of the program. For example, they do not change any variables outside of their scope.
# Do not perform any I/O operations (like reading from disk, accessing the internet, or writing from the console)
# Do not call any functions that do any of the above
# These properties result in pure functions being easier to test, debug, and think about.
#


"""
In Python, these types are passed by reference:

Lists
Dictionaries
Sets
These types are passed by value:

Integers
Floats
Strings
Booleans
Tuples
Generally speaking, most collection types are passed by reference (except for tuples) and most primitive types are passed by value.
"""
# The .copy() method can be used to create a new copy of a dictionary.

import os

def convert_file_format(filename, target_format):
  valid_extensions = ["docx", "pdf", "txt", "pptx", "ppt", "md"]
  valid_conversions = {
    "docx": ["pdf", "txt", "md"],
    "pdf": ["docx", "txt", "md"],
    "txt": ["docx", "pdf", "md"],
    "pptx": ["ppt", "pdf"],
    "ppt": ["pptx", "pdf"],
    "md": ["docx", "pdf", "txt"],
  }
  current_format = filename.split(".")[-1]
  if (
      current_format in valid_extensions
      and target_format in valid_conversions[current_format]
  ):
      return filename.replace(current_format, target_format)
  return None

def add_format(default_formats, new_format):
    copy_default_formats = default_formats.copy()
    copy_default_formats[new_format] = True
    return copy_default_formats


def remove_format(default_formats, old_format):
    copy_default_formats = default_formats.copy()
    copy_default_formats[old_format] = False
    return copy_default_formats
  
def convert_case(text, target_format):
  if not text or not target_format:
      raise ValueError(f"No text or target format provided")

  if target_format == "uppercase":
      return text.upper()
  if target_format == "lowercase":
      return text.lower()
  if target_format == "titlecase":
      return text.title()
  raise ValueError(f"Unsupported format: {target_format}")

# useless no op
def square(x):
    x * x

# this function makes no sense
# it's just useless computation
square(3)


def markdown_to_text(doc_content):
  doc_content_1 = rem_hash(doc_content)
  doc_content_2 = rem_star(doc_content_1)
  return doc_content_2
  
def rem_hash(document):
  doc_lines = document.split("\n")
  new_lines = []
  for line in doc_lines:
    new_lines.append(line.lstrip("# "))
  return("\n".join(new_lines))

def rem_star(document):
  split_lines = document.split("\n")
  new_split_lines = []
  for new_line in split_lines:
    line_words = new_line.split()
    new_words = []
    for word in line_words:
      if len(word) > 1:
        new_words.append(word.strip("*"))
      else:
        new_words.append(word)
    new_split_lines.append(" ".join(new_words))
  return("\n".join(new_split_lines))


##############
# Memos

def word_count_memo(document, memos):
  if document in memos.keys():
    return memos.get(document), memos.copy()  
  doc_word_count = word_count(document)
  memos.update({document: doc_word_count})
  return doc_word_count, memos.copy()


# Don't edit below this line


def word_count(document):
    count = len(document.split())
    return count

### Recursion

def factorial_r(x):
    if x <= 1:
      return 1
    return x * factorial_r(x-1)


### Zipmap

def zipmap(keys, values):
    if (len(keys) == 0 or len(values) == 0):
      return {}
    result = zipmap(keys[1:], values[1:])
    result[keys[0]] = values[0]
    return result
  
#tree recursion

def list_files(current_node, current_path=""):
    file_paths = []
    for node, value in current_node.items():
        if value is None:
            file_paths.append(current_path + "/" + node)
        else:
            file_paths.extend(list_files(value, current_path + "/" + node))
      
    return file_paths
  
# count tree levels

def count_nested_levels(nested_documents, target_document_id, level=1):
  for id, value in nested_documents.items():
    if id == target_document_id:
      return level
    else:
      print(level)
      current_level = count_nested_levels(value, target_document_id, level+1)
      if current_level != -1:
        return current_level
  print("did not find")
  return -1

# reverse string
def reverse_string(s):
    if s == "":
      return s
    return(s[-1]+reverse_string(s[:-1]))
  
def get_logger(formatter):
  def logger_func(*args):
    print(formatter(*args))
  return logger_func


# Don't edit below this line


def ftest(first, errors, formatter):
    print("Logs:")
    logger = get_logger(formatter)
    for err in errors:
        logger(first, err)
    print("====================================")


def colon_delimit(first, second):
    return f"{first}: {second}"


def dash_delimit(first, second):
    return f"{first} - {second}"


def fTransform():
    db_errors = [
        "out of memory",
        "cpu is pegged",
        "networking issue",
        "invalid syntax",
    ]
    ftest("Doc2Doc FATAL", db_errors, colon_delimit)

    mail_errors = [
        "email too large",
        "non alphanumeric symbols found",
    ]
    ftest("Doc2Doc WARNING", mail_errors, dash_delimit)


fTransform()

def doc_format_checker_and_converter(conversion_function, valid_formats):
  def convert_formats(filename, content):
    extension = filename.split(".")[1]
    if extension in valid_formats:
      return conversion_function(content)
    else:
      raise ValueError("Invalid file format")
  return convert_formats


# Don't edit below this line


def capitalize_content(content):
    return content.upper()


def reverse_content(content):
    return content[::-1]


## Closures

def word_count_aggregator():
    count = 0
    def count_words(content):
      nonlocal count
      count += len(content.split())
      return count
    return count_words




 

    
  


