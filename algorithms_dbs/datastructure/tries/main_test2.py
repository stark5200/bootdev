import json
from main import *

run_cases = [
    (["be", "bad", "back", "bat"], "bad", True),
    (["red", "green", "blue", "yellow"], "purple", False),
    (["boots", "tea", "ted", "ten", "i", "inn"], "boot", False),
]

submit_cases = run_cases + [
    (
        ["a", "to", "tea", "ted", "ten", "i", "in", "inn"],
        "notfound",
        False,
    ),
    (["cat", "dog", "fish", "bird"], "dog", True),
]


def test(words, word_to_check, expected_output):
    print("---------------------------------")
    trie = Trie()
    for word in words:
        trie.add(word)
    print("Trie:")
    print(json.dumps(trie.root, sort_keys=True, indent=2))
    print(f'Checking if "{word_to_check}" exists:')
    print(f"Expecting: {expected_output}")
    try:
        actual = trie.exists(word_to_check)
        print(f"Actual: {actual}")
        if actual == expected_output:
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
