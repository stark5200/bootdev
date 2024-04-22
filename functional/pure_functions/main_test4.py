from main import *

run_cases = [
    (
        "This number is intentionally wrong to test that previous memoization is working!",
        9000,
    ),
    (
        "Doc 1: In computing, memoization or memoisation is an optimization technique used primarily to speed up computer programs by storing the results of expensive function calls and returning the cached result when the same inputs occur again.",
        37,
    ),
    (
        "Doc 1: A very small document.",
        6,
    ),
]

submit_cases = run_cases + [
    ("", 0),
    (
        "Doc 1: In computing, memoization or memoisation is an optimization technique used primarily to speed up computer programs by storing the results of expensive function calls and returning the cached result when the same inputs occur again.",
        37,
    ),
    (
        "Doc 1: A very small document.",
        6,
    ),
]


def test(memos, document, expected):
    print("---------------------------------")
    print(f"Inputs:")
    print(f" * Document: {document}")
    print(f"Expected: {expected}")
    try:
        result, memos_copy = word_count_memo(document, memos)
    except Exception as e:
        result, memos_copy = e, {}
    print(f"Actual: {result}")
    if result == expected:
        print("Pass")
        return True, memos_copy
    print("Fail")
    return False, memos_copy


def main():
    test_cases = submit_cases
    if "__RUN__" in globals():
        test_cases = run_cases
    passed = 0
    failed = 0
    memos = {
        "This number is intentionally wrong to test that previous memoization is working!": 9000
    }
    correct, memos_copy = test(memos, *test_cases[0])
    if correct:
        passed += 1
    else:
        failed += 1
    for test_case in test_cases[1:]:
        correct, memos_copy = test(memos_copy, *test_case)
        if correct:
            passed += 1
        else:
            failed += 1

    memos[
        "This number is intentionally wrong to test that previous memoization is working!"
    ] = 1
    correct, _ = test(memos_copy, *test_cases[0])
    if correct:
        passed += 1
    else:
        failed += 1
        print("The function has not been returning a copy of the memo!")

    if failed == 0:
        print("============= PASS ==============")
    else:
        print("============= FAIL ==============")
    print(f"{passed} passed, {failed} failed")


main()
