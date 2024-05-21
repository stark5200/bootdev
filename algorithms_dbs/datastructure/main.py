# Fantasy Quest

def count_potions(inventory):
    count = 0
    for item in inventory:
        if item == "Healing potion":
            count += 1
    return count

class Stack:
    def __init__(self):
        self.arrows = []

    def push(self, arrow):
        self.arrows.append(arrow)

    def pop(self):
        if len(self.arrows) == 0:
            return None
        current = self.arrows[-1]
        del self.arrows[-1]
        return current

    def peek(self):
        if len(self.arrows) == 0:
            return None
        return self.arrows[-1]

    def size(self):
        return len(self.arrows)
