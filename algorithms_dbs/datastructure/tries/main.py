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
  
    def find_matches(self, document):
        new_set = set()
        for index in range(len(document)):
            level = self.root
            for i in range(index, len(document)):
                if document[i] not in level:
                    break
                level = level[document[i]]
                if "*" in level:
                    new_set.add(document[index:i+1])
        return new_set
  
    def words_with_prefix(self, prefix):
        words = []
        level = self.root
        for letter in prefix:
            if letter in level:
                level = level[letter]
            else:
                return words
        words = self.search_level(level, prefix, words)
        return words

    def search_level(self, cur, cur_prefix, words):
        if self.end_symbol in cur:
            words.append(cur_prefix)
        current_keys = cur.keys()
        for key in sorted(current_keys):
            if key != self.end_symbol:
                self.search_level(cur[key], cur_prefix+key, words)
        return words
  
    def exists(self, word):
        current = self.root
        for c in word:
            if c in current:
                current = current[c]
            else:
                return False
        if "*" in current:
            return True
        return False
  
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