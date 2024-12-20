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


# Trees

# BST
class BSTNode:
    def __init__(self, val=None):
        self.left = None
        self.right = None
        self.val = val
        
    def exists(self, val):
        existing = self.inorder([])
        if val in existing:
            return True
        return False    
        
    def preorder(self, visited):
        if self.val is None:
            return visited
        if self.val not in visited:
            visited.append(self.val)
        if self.left and (self.left not in visited):
            visited = self.left.preorder(visited)
        if self.right and (self.right not in visited):
            visited = self.right.preorder(visited)
        return visited
      
    def inorder(self, visited):
        if self.val is None:
            return visited
        if self.left and (self.left not in visited):
            visited = self.left.inorder(visited)
        if self.val not in visited:
            visited.append(self.val)
        if self.right and (self.right not in visited):
            visited = self.right.inorder(visited)
        return visited  
      
    def postorder(self, visited):
        if self.val is None:
            return visited
        if self.left and (self.left not in visited):
            visited = self.left.postorder(visited)
        if self.right and (self.right not in visited):
            visited = self.right.postorder(visited)
        if self.val not in visited:
            visited.append(self.val)
        return visited

    def insert(self, val):
        if self.val == None:
            self.val = val
            return
        if self.val == val:
            return
            
        if val < self.val:
            if self.left == None:
                self.left = BSTNode(val)
                return
            self.left.insert(val)
            return
            
        if val > self.val:
            if self.right == None:
                self.right = BSTNode(val)
                return
            self.right.insert(val)
            return
          
    def delete(self, val):
        if self.val == None:
            return None
        if val < self.val:
            if self.left:
                self.left = self.left.delete(val)
            return self
        if val > self.val:
            if self.right:
                self.right = self.right.delete(val)
            return self
            
        if val == self.val:
            if self.left == None:
                return self.right
            if self.right == None:
                return self.left
            min_larger_node = self.right.get_min()
            self.val = min_larger_node.val
            self.right = self.right.delete(min_larger_node.val)
            return self
          
    def get_min(self):
        if self.left:
            return self.left.get_min()
        return self

    def get_max(self):
        if self.right:
            return self.right.get_max()
        return self

# RBTree

class RBNode:
    def __init__(self, val):
        self.red = False
        self.parent = None
        self.val = val
        self.left = None
        self.right = None


class RBTree:
    def __init__(self):
        self.nil = RBNode(None)
        self.nil.red = False
        self.nil.left = None
        self.nil.right = None
        self.root = self.nil

    def insert(self, val): #old insert
        new_node = RBNode(val)
        new_parent = None
        new_node.left = self.nil
        new_node.right = self.nil
        new_node.red = True
        current = self.root
        while current != self.nil:
            new_parent = current 
            if new_node.val < current.val:
                current = current.left
            elif new_node.val > current.val:
                current = current.right
            else:
                return
        new_node.parent = new_parent
        if new_parent is None:
            self.root = new_node
        else:
            if new_node.val < new_parent.val:
                new_node.parent.left = new_node
            if new_node.val > new_parent.val:
                new_node.parent.right = new_node
                
        self.fix_insert(new_node)
                
    def fix_insert(self, new_node): #new fix_insert
        while ((new_node != self.root) and (new_node.parent.red == True)):
            if new_node.parent == new_node.parent.parent.right:
                uncle = new_node.parent.parent.left
                if uncle.red:
                    uncle.red = False
                    new_node.parent.red = False
                    new_node.parent.parent.red = True
                    new_node = new_node.parent.parent
                else:
                    if new_node == new_node.parent.left:
                        new_node = new_node.parent
                        self.rotate_right(new_node)
                    new_node.parent.red = False
                    new_node.parent.parent.red = True
                    self.rotate_left(new_node.parent.parent)
            
            elif new_node.parent == new_node.parent.parent.left:
                uncle = new_node.parent.parent.right
                
                if uncle.red:
                    uncle.red = False
                    new_node.parent.red = False
                    new_node.parent.parent.red = True
                    new_node = new_node.parent.parent
                else:
                    if new_node == new_node.parent.right:
                        new_node = new_node.parent
                        self.rotate_left(new_node)
                    new_node.parent.red = False
                    new_node.parent.parent.red = True
                    self.rotate_right(new_node.parent.parent)
                    
        self.root.red = False
                

    def exists(self, val):
        curr = self.root
        while curr != self.nil and val != curr.val:
            if val < curr.val:
                curr = curr.left
            else:
                curr = curr.right
        return curr

                
    def rotate_left(self, x):
        if x == self.nil or x.right == self.nil:
            return 
        y = x.right
        x.right = y.left
        
        if y.left != self.nil:
            y.left.parent = x
        y.parent = x.parent
        
        if x.parent is None:
            self.root = y
        elif x == x.parent.left:
            x.parent.left = y
        elif x == x.parent.right:
            x.parent.right = y
        y.left = x
        x.parent = y

    def rotate_right(self, x):
        if x == self.nil or x.left == self.nil:
            return 
        y = x.left
        x.left = y.right
        if y.right != self.nil:
            y.right.parent = x
        y.parent = x.parent
        
        if x.parent is None:
            self.root = y
        elif x == x.parent.right:
            x.parent.right = y
        elif x == x.parent.left:
            x.parent.left = y
        y.right = x
        x.parent = y
        
        
#### errors
def handle_get_player_record(player_id):
    try:
        return get_player_record(player_id)
    except IndexError:
        return("index is too high")
    except Exception as e:
        return e


# Don't edit below this line


def get_player_record(player_id):
    if player_id < 0:
        raise Exception("negative ids not allowed")
    players = [
        {"name": "Slayer", "level": 128},
        {"name": "Dorgoth", "level": 300},
        {"name": "Saruman", "level": 4000},
    ]
    return players[player_id]
        
        