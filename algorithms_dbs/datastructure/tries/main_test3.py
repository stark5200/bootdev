import json
from main import *

run_cases = [
    (
        ["bad", "baddie", "badguy", "suck"],
        "the badguy really sucks",
        ["bad", "badguy", "suck"],
    ),
    (["be", "bad", "back", "bat"], "he is back at bat", ["back", "bat"]),
]

submit_cases = run_cases + [
    (["go", "dang", "mid"], "gosh darn it dan, help me middle", ["go", "mid"]),
]


def test(words, document, expected_matches):
    print("---------------------------------")
    print("Trie:")
    trie = Trie()
    for word in words:
        trie.add(word)
    print(json.dumps(trie.root, sort_keys=True, indent=2))
    print(f"Expected matches: {sorted(expected_matches)}")
    try:
        actual = sorted(trie.find_matches(document))
        print(f"Actual matches: {actual}")
        if actual == sorted(expected_matches):
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
