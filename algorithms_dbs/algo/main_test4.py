from main import *

run_cases = [
    ([["George", "Eva", "George"], ["Diane", "George", "Eva", "Frank"]], "George", 3),
    (
        [
            ["Amy", "Bob", "Candy"],
            ["Diane", "George", "Eva", "Frank"],
            ["Diane", "George"],
            ["George", "name", "George"],
        ],
        "George",
        4,
    ),
]

submit_cases = run_cases + [
    (
        [
            ["Alex", "name", "Chloe"],
            ["Eric", "name", "Fred"],
            ["Hector", "name"],
            ["Hector", "name"],
            ["Hector", "name"],
            ["George"],
        ],
        "Hector",
        3,
    ),
    (
        [
            ["Alex", "name", "Chloe"],
            ["Eric", "name", "Fred"],
            ["Hector", "name"],
            ["Hector", "name"],
            ["Hector", "name"],
            ["George"],
        ],
        "George",
        1,
    ),
    (
        [["Alex", "name", "Chloe"], ["Eric", "name", "Fred"], ["Hector", "name"]],
        "Alex",
        1,
    ),
    ([], "George", 0),
]


def test(input1, input2, expected_output):
    print("---------------------------------")
    print(f"Inputs:")
    print(f" * list of lists: {input1}")
    print(f" * target name: {input2}")
    print(f"Expecting: {expected_output}")
    result = count_names(input1, input2)
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
