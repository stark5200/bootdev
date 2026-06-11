def count_words(text):
  words = text.split()
  return len(words)

def count_unique_chars(text):
  chars = set(text.lower())
  char_count = {}
  for char in chars:
    char_count[char] = text.lower().count(char)
  return char_count