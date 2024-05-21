# Fantasy Quest

# Calling a function pushes a call frame onto the runtime stack
# Returning from a function pops the top frame off the stack.

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

# Simulating Function call stack

def is_balanced(input_str):
    stack = BigStack()
    if input_str == None:
        print(str(stack.items))
        return None
    input = [*input_str]
    print(input)
    for i in input:
        if i == "(":
            stack.push(i)
            print(str(stack.items))
        if i == ")":
            if len(stack.items) == 0:
                print(str(stack.items))
                return False
            print(str(stack.items))
            stack.pop()
    if len(stack.items) > 0:
        print(str(stack.items))
        return False
    print(str(stack.items))
    return True


# don't modify below this line


class BigStack:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        if len(self.items) == 0:
            return None
        return self.items.pop()

    def peek(self):
        if len(self.items) == 0:
            return None
        return self.items[len(self.items) - 1]


def attack_action():
    call(shoot_arrow)
    call(calc_new_health)


def shoot_arrow():
    call(calc_damage)


def calc_damage():
    call(apply_damage)


# don't touch below this line


def calc_new_health():
    pass


def apply_damage():
    pass


class Stack:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        if len(self.items) == 0:
            return None
        return self.items.pop()

    def peek(self):
        if len(self.items) == 0:
            return None
        return self.items[len(self.items) - 1]


stack = Stack()


def call(func):
    stack.push(func.__name__)
    print("Pushing " + func.__name__)
    print("Stack: " + str(stack.items))
    print("=================================")
    func()
    stack.pop()
    print("Popping " + func.__name__)
    print("Stack: " + str(stack.items))
    print("=================================")


call(attack_action)