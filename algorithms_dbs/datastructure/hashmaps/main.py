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
        self.resize()

        index = self.key_to_index(key)
        self.hashmap[index] = (key, value)
        
    def resize(self):
        if len(self.hashmap) == 0:
            self.hashmap.append(None)
        if self.current_load() < .05:
            return
        old_hashmap = self.hashmap
        self.hashmap = [None for i in range(10*len(old_hashmap))]
        for item in old_hashmap:
            if item != None:
                self.insert(item[0], item[1])
            

    def current_load(self):
        current_size = 0
        for i in self.hashmap:
            if i != None:
                current_size += 1
        if len(self.hashmap) == 0:
            return 1
        return current_size/len(self.hashmap)
  
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
