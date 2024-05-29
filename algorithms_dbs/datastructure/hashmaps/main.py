# dictionaries use hashmaps
m = {}
m["Ross"] = "Bob"
m["Jenner"] = "Kylie"

class HashMap:
  
    def get(self, key):
        t = self.hashmap[self.key_to_index(key)]
        if t is None:
            raise Exception("sorry, key not found")
        return t[1]
  
    def insert(self, key, value):
        self.hashmap[self.key_to_index(key)] = (key, value)
  
    def key_to_index(self, key):
        sum = 0
        for c in key:
            sum += ord(c)
        index = sum%len(self.hashmap)
        return index

    # don't touch below this line

    def __init__(self, size):
        self.hashmap = [None for i in range(size)]

    def __repr__(self):
        buckets = []
        for v in self.hashmap:
            if v != None:
                buckets.append(v)
        return str(buckets)
