import random
from main import *

run_cases = [
    (4),
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


def ref_impl_left(self, x):
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
    else:
        x.parent.right = y
    y.left = x
    x.parent = y


def ref_impl_right(self, x):
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
    else:
        x.parent.left = y
    y.right = x
    x.parent = y


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


def print_tree(node):
    lines = []
    format_tree_string(node.root, lines)
    return "\n".join(lines)


setattr(RBTree, "ref_impl_right", ref_impl_right)
setattr(RBTree, "ref_impl_left", ref_impl_left)
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


def compare_trees(node, reference_node):
    if node == None and reference_node == None:
        return True
    if node.val != reference_node.val:
        print(f"Fail 2: {node.val} value does not match {reference_node.val}")
        return False

    if reference_node.parent is None:
        ref_parent = None
    else:
        ref_parent = reference_node.parent.val

    if (node.parent is None) != (ref_parent is None):
        print(
            f"Fail 3: {node.val} parent ({node.parent.val}) does not match {ref_parent}"
        )
        return False
    if node.parent is not None:
        if node.parent.val != ref_parent:
            print(
                f"Fail 4: {node.val} parent ({node.parent.val}) does not match {ref_parent}"
            )
            return False

    return compare_trees(node.left, reference_node.left) and compare_trees(
        node.right, reference_node.right
    )


def test_rotate(tree, node, reference_tree, reference_node, direction):
    print(f"Rotating {direction} at {node.val}...")
    print("-------------------------------------")
    if direction == "left":
        tree.rotate_left(node)
        reference_tree.ref_impl_left(reference_node)
    else:
        tree.rotate_right(node)
        reference_tree.ref_impl_right(reference_node)
    print("Expected Tree:")
    print("-------------------------------------")
    print(reference_tree)
    print("-------------------------------------")
    print("Actual Tree:")
    print("-------------------------------------")
    print(tree)
    print("-------------------------------------")
    return compare_trees(tree.root, reference_tree.root)


def test_rotations(tree, reference_tree):
    return (
        test_rotate(tree, tree.root, reference_tree, reference_tree.root, "left")
        and test_rotate(tree, tree.root, reference_tree, reference_tree.root, "right")
        and test_rotate(
            tree, tree.root.right, reference_tree, reference_tree.root.right, "left"
        )
        and test_rotate(
            tree, tree.root.left, reference_tree, reference_tree.root.left, "right"
        )
    )


def test(num_characters):
    characters = get_characters(num_characters)
    tree = RBTree()
    reference_tree = RBTree()
    for character in characters:
        tree.insert(character)
        reference_tree.insert(character)
    print("=====================================")
    print("Starting Tree:")
    print("-------------------------------------")
    print(tree)
    print("-------------------------------------")

    if test_rotations(tree, reference_tree):
        print("Pass \n")
        return True
    print("Fail 1 \n")
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
