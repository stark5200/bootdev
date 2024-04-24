from main import *

run_cases = [
    (
        ["name", "age", "city"],
        ["Alice", 25, "Wonderland"],
        {"city": "Wonderland", "age": 25, "name": "Alice"},
    ),
    (
        ["name", "age", "city"],
        ["Bob", 30, "Atlantis"],
        {"city": "Atlantis", "age": 30, "name": "Bob"},
    ),
]

submit_cases = run_cases + [
    ([], [], {}),
    (
        ["key1", "key2", "key3"],
        ["value1", "value2", "value3"],
        {"key3": "value3", "key2": "value2", "key1": "value1"},
    ),
    (
        ["father", "sister", "brother"],
        ["Homer", "Lisa", "Bart"],
        {"brother": "Bart", "sister": "Lisa", "father": "Homer"},
    ),
    (
        ["country", "state"],
        ["USA", "California"],
        {"state": "California", "country": "USA"},
    ),
]


def test(keys, values, expected_output):
    print("---------------------------------")
    print(f"Inputs: {keys}, {values}")
    print(f"Expecting: {expected_output}")
    result = zipmap(keys, values)
    print(f"Actual: {result}")
    if result == expected_output:
        print("Pass")
        return True
    print("Fail")
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
