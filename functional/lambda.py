def categorize_file(filename):
    # ?
    return get_category(filename[filename.rfind(".") :])

def get_category(input):
  return input