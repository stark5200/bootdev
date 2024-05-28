import random
from main import *

run_cases = [
    (4),
    (8),
]

submit_cases = run_cases + [
    (10),
]


class Character:
    def __init__(self, gamertag):
        self.gamertag = gamertag
        character_names = [
            "Ebork",
            "Astram",
            "Elian",
            "Tarlock",
            "Grog",
            "Myra",
            "Sullivan",
            "Marlo",
            "Jax",
            "Anthony",
            "Bhurdan",
            "Thoreuth",
            "Bob",
            "Varis",
            "Nyx",
            "Luna",
            "Ash",
            "Rhogar",
            "Ember",
            "Mikel",
            "Yamil",
            "Velithria",
        ]
        self.character_name = (
            f"{character_names[gamertag%len(character_names)]}#{gamertag}"
        )

    def __eq__(self, other):
        return isinstance(other, Character) and self.gamertag == other.gamertag

    def __lt__(self, other):
        return isinstance(other, Character) and self.gamertag < other.gamertag

    def __gt__(self, other):
        return isinstance(other, Character) and self.gamertag > other.gamertag

    def __repr__(self):
        return "".join(self.character_name)


def ref_implementation(self, val):
    n = RBNode(val)
    n.parent = None
    n.left = self.nil
    n.right = self.nil
    n.red = True

    p = None
    current = self.root
    while current != self.nil:
        p = current
        if n.val < current.val:
            current = current.left
        elif n.val > current.val:
            current = current.right
        else:
            return

    n.parent = p
    if p is None:
        self.root = n
    elif n.val < p.val:
        p.left = n
    else:
        p.right = n


def inorder(self, visited):
    if self.left:
        visited = self.left.inorder(visited)
    visited.append(self.val)
    if self.right:
        visited = self.right.inorder(visited)
    return visited


def print_tree(node):
    lines = []
    format_tree_string(node.root, lines)
    return "\n".join(lines)


def format_tree_string(node, lines, level=0):
    if node.val is not None:
        format_tree_string(node.right, lines, level + 1)
        lines.append(
            " " * 4 * level
            + "> "
            + str(node.val)
            + " "
            + ("[red]" if node.red else "[black]")
        )
        format_tree_string(node.left, lines, level + 1)


setattr(RBTree, "ref_implementation", ref_implementation)
setattr(RBNode, "inorder", inorder)
setattr(RBTree, "__repr__", print_tree)


def get_characters(num):
    random.seed(1)
    characters = []
    gamertags = []
    for i in range(num * 3):
        gamertags.append(i)
    random.shuffle(gamertags)
    gamertags = gamertags[:num]
    for gamertag in gamertags:
        character = Character(gamertag)
        characters.append(character)
    return characters


def char_list_to_string(char_list):
    character_names = []
    for char in char_list:
        character_names.append(char.character_name)
    return character_names


def test(num_characters):
    characters = get_characters(num_characters)
    tree = RBTree()
    for character in characters:
        tree.ref_implementation(character)
    print("=====================================")
    print("Expecting Tree:")
    print("-------------------------------------")
    print(tree)
    print("-------------------------------------\n")
    actual_tree = RBTree()
    for character in characters:
        print(f"Inserting {character} into tree...")
        actual_tree.insert(character)
    print("\n")
    print("Actual Tree:")
    print("-------------------------------------")
    print(actual_tree)
    print("-------------------------------------")
    if actual_tree.root.inorder([]) == tree.root.inorder([]):
        print("Pass \n")
        return True
    print("Fail \n")
    return False


def main():
    passed = 0
    failed = 0
    for test_case in test_cases:
        correct = test(test_case)
        if correct:
            passed += 1
        else:
            failed += 1
    if failed == 0:
        print("============= PASS ==============")
    else:
        print("============= FAIL ==============")
    print(f"{passed} passed, {failed} failed")


test_cases = submit_cases
if "__RUN__" in globals():
    test_cases = run_cases

main()
