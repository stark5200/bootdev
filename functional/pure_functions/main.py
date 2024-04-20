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
