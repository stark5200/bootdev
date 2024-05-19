from main import *
import time

run_cases = [
    (10, [i for i in range(200)], True),
    (-1, [i for i in range(20000)], False),
    (15, [], False),
    (0, [0], True),
    (-1, [-2, -1], True),
    (105028, [i for i in range(2000000)], True),
    (2000001, [i for i in range(2000000)], False),
]

submit_cases = run_cases + [
    (15, [], False),
    (0, [0], True),
    (-1, [-2, -1], True),
    (105028, [i for i in range(2000000)], True),
    (2000001, [i for i in range(2000000)], False),
]


def test(target, arr, expected_output):
    print("---------------------------------")
    print(f"Inputs:")
    print(f" * target: {target}")
    print(f" * arr length: {len(arr)} items")
    print(f"Expecting: {expected_output} & completed in less than 50 milliseconds")
    start = time.time()
    result = binary_search(target, arr)
    end = time.time()
    timeout = 0.05
    if (end - start) < timeout:
        print(f"binary_search completed in less than {timeout * 1000} milliseconds!")
        if result == expected_output:
            print(f"Actual: {result}")
            print("Pass")
            return True
        else:
            print(f"Actual: {result}")
            print("Fail")
            return False
    else:
        print(
            f"binary_search took too long ({(end - start) * 1000} milliseconds). Speed it up!"
        )
        print(f"Actual: {result}")
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


test_cases = run_cases
if "__RUN__" in globals():
    test_cases = run_cases

main()
