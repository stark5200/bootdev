from main import *

run_cases = [
    (
        """The Importance of FP
Learn how functional programming can change the way you think about code.
Benefits include immutability, simplicity, and composability.""",
        """          The Importance of FP          
****************************************
Learn how functional programming can change the way you think about code.
Benefits include immutability, simplicity, and composability.""",
    ),
]

submit_cases = run_cases + [
    (
        """Short Title
Equally short story""",
        """              Short Title               \n****************************************
Equally short story""",
    ),
    (
        """DocToDoc: A Guide
Understanding the art of document conversion.
We write cool functional code to make it happen.""",
        """           DocToDoc: A Guide            
****************************************
Understanding the art of document conversion.
We write cool functional code to make it happen.""",
    ),
]


def test(input1, expected_output):
    pass


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
