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
