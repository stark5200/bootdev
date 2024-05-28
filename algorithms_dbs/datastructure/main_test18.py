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


def ref_impl_ins(self, val):
    new_node = RBNode(val)
    new_node.parent = None
    new_node.left = self.nil
    new_node.right = self.nil
    new_node.red = True

    parent = None
    current = self.root
    while current != self.nil:
        parent = current
        if new_node.val < current.val:
            current = current.left
        elif new_node.val > current.val:
            current = current.right
        else:
            # duplicate, just ignore
            return

    new_node.parent = parent
    if parent is None:
        self.root = new_node
    elif new_node.val < parent.val:
        parent.left = new_node
    else:
        parent.right = new_node

    self.ref_impl_fix(new_node)


def ref_impl_fix(self, new_node):
    while new_node != self.root and new_node.parent.red:
        if new_node.parent == new_node.parent.parent.right:
            u = new_node.parent.parent.left
            if u.red:
                u.red = False
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
        else:
            u = new_node.parent.parent.right

            if u.red:
                u.red = False
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


setattr(RBTree, "__repr__", print_tree)
setattr(RBTree, "ref_impl_insert", ref_impl_ins)
setattr(RBTree, "ref_impl_fix", ref_impl_fix)


def get_characters(num):
    characters = []
    for i in range(num):
        character = Character(i)
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
    reference_tree = RBTree()
    for character in characters:
        tree.insert(character)
        reference_tree.ref_impl_insert(character)
    print("=====================================")
    print("Expecting:")
    print("-------------------------------------")
    print(reference_tree)
    print("-------------------------------------\n")
    print("Actual:")
    print("-------------------------------------")
    print(tree)
    print("-------------------------------------\n")

    if print_tree(tree) == print_tree(reference_tree):
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
