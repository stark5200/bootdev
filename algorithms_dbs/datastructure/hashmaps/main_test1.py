from main import *

run_cases = [
    (
        4,
        ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j"],
        [1, 2, 3, 0, 1, 2, 3, 0, 1, 2],
    ),
    (256, ["hello", "world"], [20, 40]),
]

submit_cases = run_cases + [
    (
        512,
        ["golang", "python", "java", "javascript", "rust", "c", "c++"],
        [120, 162, 418, 55, 462, 99, 185],
    ),
]


def test(size, keys, expected_indexes):
    print("---------------------------------")
    print(f"Inputs:")
    print(f" * Size: {size}")
    print(f" * Keys: {keys}")
    hm = HashMap(size)
    try:
        actual = []
        for i, key in enumerate(keys):
            index = hm.key_to_index(key)
            print(f"Expecting:")
            print(f"  {key} hashes to index {expected_indexes[i]}")
            print(f"Actual:")
            print(f"  {key} hashes to index {index}")
            actual.append(index)
        if actual == expected_indexes:
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
