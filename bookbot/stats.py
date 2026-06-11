def count_words(text):
  words = text.split()
  return len(words)

def count_unique_chars(text):
  chars = set(text.lower())
  char_count = {}
  for char in chars:
    char_count[char] = text.lower().count(char)
  return char_count

def sort_on(char: tuple[str, int]):
  return char[1]

def chars_dict_to_sorted_list(char_count_dict):
  char_count_list = []
  for char, count in char_count_dict.items():
    char_count_list.append((char, count))
  char_count_list = sorted(char_count_list, key=sort_on, reverse=True)
  return char_count_list