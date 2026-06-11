from stats import count_words, count_unique_chars

book_path = "books/frankenstein.txt"    
    
def get_book_text(book_path):
  with open(book_path) as f:
    file_contents = f.read()
  return file_contents

def main():
  book_text = get_book_text(book_path)
  word_count = count_words(book_text)
  print(f"Found {word_count} total words") # Print the word count
  unique_char_count = count_unique_chars(book_text)
  print(unique_char_count) # Print the unique character count

main()