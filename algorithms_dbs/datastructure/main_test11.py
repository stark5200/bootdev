import random
from main import *

run_cases = [
    (6, 2, ["Ebork#0", "Anthony#9", "Ash#16", "Rhogar#17"]),
    (
        12,
        4,
        [
            "Elian#2",
            "Bhurdan#10",
            "Thoreuth#11",
            "Rhogar#17",
            "Ebork#22",
            "Myra#27",
            "Jax#30",
            "Thoreuth#33",
        ],
    ),
]

submit_cases = run_cases + [
    (
        24,
        6,
        [
            "Elian#2",
            "Tarlock#3",
            "Anthony#9",
            "Bhurdan#10",
            "Bob#12",
            "Ash#16",
            "Ember#18",
            "Mikel#19",
            "Ebork#22",
            "Astram#23",
            "Varis#35",
            "Rhogar#39",
            "Astram#45",
            "Marlo#51",
            "Bhurdan#54",
            "Elian#68",
            "Tarlock#69",
            "Grog#70",
        ],
    ),
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


def inorder_char_names(self, visited):
    if self.left:
        visited = self.left.inorder_char_names(visited)
    visited.append(self.val.character_name)
    if self.right:
        visited = self.right.inorder_char_names(visited)
    return visited


def print_tree(bst_node):
    lines = []
    format_tree_string(bst_node, lines)
    return "\n".join(lines)


def format_tree_string(bst_node, lines, level=0):
    if bst_node != None:
        format_tree_string(bst_node.right, lines, level + 1)
        lines.append(" " * 4 * level + "> " + str(bst_node.val))
        format_tree_string(bst_node.left, lines, level + 1)


setattr(BSTNode, "inorder_char_names", inorder_char_names)
setattr(BSTNode, "__repr__", print_tree)


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


def test(num_characters, num_to_delete, expected):
    characters = get_characters(num_characters)
    characters_copy = characters.copy()
    random.shuffle(characters_copy)
    characters_to_delete = characters_copy[:num_to_delete]
    bst = BSTNode()
    for character in characters:
        bst.insert(character)
    print("=====================================")
    print("Tree:")
    print("-------------------------------------")
    print(bst)
    print("-------------------------------------\n")
    try:
        actual_bst = BSTNode()
        for character in characters:
            actual_bst.insert(character)
        print("Deleting characters: " + str(characters_to_delete))
        for character in characters_to_delete:
            actual_bst = actual_bst.delete(character)
        print("Actual Tree:")
        print("-------------------------------------")
        print(actual_bst)
        print("-------------------------------------")
        actual = actual_bst.inorder_char_names([])
        print(f"Expecting: {expected}")
        print(f"Actual: {actual}")
        if expected == actual:
            print("Pass \n")
            return True
        print("Fail \n")
        return False
    except Exception as e:
        print(f"Error: {e}")
        return False


def main():
    passed = 0
    failed = 0
    for test_case in test_cases:
        correct = test(*test_case)
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
