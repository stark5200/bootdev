from main import *

run_cases = [
    (["Short sword", "Bread", "Healing potion", "Healing potion"], 2),
    (["Bread", "Short sword"], 0),
]

submit_cases = run_cases + [
    ([], 0),
    (["Healing potion"], 1),
    (["Healing potion", "Healing potion", "Bread", "Short sword", "Healing potion"], 3),
    (["Bread", "Short sword", "Chainmail Armor"], 0),
    (["Healing potion", "Short sword", "Chainmail Armor", "Healing potion"], 2),
]


def test(input1, expected_output):
    print("---------------------------------")
    print(f"Inputs: {input1}")
    print(f"Expecting: {expected_output}")
    result = count_potions(input1)
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
