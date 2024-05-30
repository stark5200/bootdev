import json
from main import *

run_cases = [
    (
        ["arc", "art", "arg"],
        {"a": {"r": {"c": {"*": True}, "g": {"*": True}, "t": {"*": True}}}},
    ),
    (
        ["be", "bad", "back", "bat"],
        {
            "b": {
                "a": {"c": {"k": {"*": True}}, "d": {"*": True}, "t": {"*": True}},
                "e": {"*": True},
            }
        },
    ),
]

submit_cases = run_cases + [
    (
        ["a", "to", "tea", "ted", "ten", "i", "in", "inn"],
        {
            "a": {"*": True},
            "i": {"*": True, "n": {"*": True, "n": {"*": True}}},
            "t": {
                "e": {"a": {"*": True}, "d": {"*": True}, "n": {"*": True}},
                "o": {"*": True},
            },
        },
    ),
]


def test(words, expected_trie):
    print("---------------------------------")
    print(f"Inputs:")
    print(f" * Words: {words}")
    print(" * Expected trie:")
    print(f"{json.dumps(expected_trie, sort_keys=True, indent=2)}")
    try:
        trie = Trie()
        for word in words:
            trie.add(word)
            print(f"Adding {word}...")
        print("Actual Trie:")
        print(json.dumps(trie.root, sort_keys=True, indent=2))
        if trie.root == expected_trie:
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
