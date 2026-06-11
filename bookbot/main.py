import sys
from stats import count_words, count_unique_chars, chars_dict_to_sorted_list
    
def get_book_text(book_path):
  with open(book_path) as f:
    file_contents = f.read()
  return file_contents

def print_report(book_path, word_count, sorted_list):
  print("============ BOOKBOT ============")
  print(f"Analyzing book found at {book_path}...")
  print("----------- Word Count ----------")
  print(f"Found {word_count} total words")
  print("--------- Character Count -------")
  for char, count in sorted_list:
    if char.isalpha():
      print(f"{char}: {count}")
  print("============= END ===============")

  

def main():
  if len(sys.argv) != 2:
    print("Usage: python3 main.py <path_to_book>")
    sys.exit(1)
  
  book_path = sys.argv[1]
  book_text = get_book_text(book_path)
  word_count = count_words(book_text)
  unique_char_count = count_unique_chars(book_text)
  #print(unique_char_count) # Print the unique character count
  sorted_list = chars_dict_to_sorted_list(unique_char_count)
  print_report(book_path, word_count, sorted_list)

main()