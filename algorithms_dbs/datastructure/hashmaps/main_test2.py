from main import *

run_cases = [
    (
        256,
        [
            ("apple", 1),
            ("banana", 2),
            ("cherry", 3),
            ("mango", 4),
        ],
        ["apple", "banana", "garbage"],
        [4, 2, "sorry, key not found"],
    ),
    (
        512,
        [
            ("golang", 1),
            ("python", 2),
            ("java", 3),
            ("javascript", 4),
            ("rust", 5),
            ("c", 6),
            ("c++", 7),
        ],
        ["golang", "python", "garbage"],
        [1, 2, "sorry, key not found"],
    ),
]

submit_cases = run_cases + [
    (
        256,
        [
            ("football", 1),
            ("basketball", 2),
            ("skateboard", 3),
        ],
        ["rollerblades", "hockey puck", "skateboard"],
        ["sorry, key not found", "sorry, key not found", 3],
    ),
]


def test(size, items, keys_to_get, expected_values):
    print("---------------------------------")
    hm = HashMap(size)
    for item in items:
        key = item[0]
        val = item[1]
        hm.insert(key, val)
    print(f"Hashmap:\n{hm}")
    try:
        actual = []
        for i, key in enumerate(keys_to_get):
            try:
                val = hm.get(key)
                print(f"Get({key}) -> Expecting: {expected_values[i]} Received: {val}")
                actual.append(val)
            except Exception as e:
                print(
                    f"get({key}) -> Expecting: '{expected_values[i]}' Received: '{e}'"
                )
                actual.append(e.args[0])
        print("=====================================")
        if actual == expected_values:
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