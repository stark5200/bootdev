# Tries or "prefix tree" 
#Tries are data structures that can be used for language processing tasks for example. At its core, a trie is often represented as a nested tree of dictionaries where each key is a character, and it maps to the next character in the word. For example, the words "hello", "help" and "hi" would be represented as:

words = {
        "h": {
          "e": {
            "l": {
              "l": {
                "o": {
                  "*": True
                }
              },
              "p": {
                "*": True
              }
            }
          },
          "i": {
            "*": True
          }
        }
      }

class Trie:
  
    def exists(self, word):
        pass
  
    def add(self, word):
        current = self.root
        for c in word:
            if c not in current:
                current[c] = {}
            current = current[c]
            
        current[self.end_symbol] = True

    # don't touch below this line

    def __init__(self):
        self.root = {}
        self.end_symbol = "*"