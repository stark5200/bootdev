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

# matchmaking
def matchmake(queue, player):
    player_name = player[0]
    player_action = player[1]
    # match_size = queue.size()
    if player_action == 'leave':
        queue.search_and_remove(player_name)
    if player_action == 'join':
        queue.push(player_name)
        
    match_size = queue.size()
    if match_size < 4:
        return "No match found"
        
    player_1 = queue.pop()
    player_2 = queue.pop()
    
    return f"{player_1} matched {player_2}!"

# Queue
class Queue:
    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.insert(0, item)  # self.items = [item, *self.items]

    def pop(self):
        if len(self.items) == 0:
            return None
        temp = self.items[-1]
        del self.items[-1]
        return temp

    def peek(self):
        if len(self.items) == 0:
            return None
        return self.items[-1]

    def size(self):
        return len(self.items)
      
    def search_and_remove(self, item):
        if item not in self.items:
            return None
        self.items.remove(item)
        return item
      
    def __repr__(self):
        return f"[{', '.join(self.items)}]"
      
      
# Linked list

class LinkedList:
    def __init__(self):
        self.head = None
        self.tail = None

    def __iter__(self):
        node = self.head
        while node != None:
            yield node
            node = node.next

    # don't touch below this line

    def __repr__(self):
        nodes = []
        for node in self:
            nodes.append(node.val)
        return " -> ".join(nodes)

    def add_to_tail(self, node):
        if self.head is None:
            self.head = node
            self.tail = node
            return
        self.tail.set_next(node)
        self.tail = node
          
    def add_to_head(self, node):
        if self.head is None:
            self.head = node
            self.tail = node
            return
        node.set_next(self.head)
        self.head = node
        
# LLQueue

class LLQueue:
    def remove_from_head(self):
        if self.head is None:
            return None
        # make shift solution for LL with 1 item
        #if self.head == self.tail:
        #    single_item = self.head
        #    self.head = None
        #    self.tail = None
        #    return single_item  
        temp = self.head
        self.head = self.head.next
        if self.head is None:
            self.tail = None
        return temp

    # don't touch below this line

    def add_to_tail(self, node):
        if self.head is None:
            self.head = node
            self.tail = node
            return
        self.tail.next = node
        self.tail = node

    def __init__(self):
        self.tail = None
        self.head = None

    def __iter__(self):
        node = self.head
        while node is not None:
            yield node
            node = node.next

    def __repr__(self):
        nodes = []
        for node in self:
            nodes.append(node.val)
        return " <- ".join(nodes)



class Node:
    def __init__(self, val):
        self.val = val
        self.next = None

    def set_next(self, node):
        self.next = node

    def __repr__(self):
        return self.val
