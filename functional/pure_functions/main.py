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


