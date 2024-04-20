from main import *

run_cases = [
    (
        ["The world is changed", "I feel it in the water", "I feel it in the earth"],
        2,
        "The world is changed. I feel it in the water.",
    ),
    (["Hello", "world", "!"], 2, "Hello. world."),
]

submit_cases = run_cases + [
    ([], 0, ""),
    ([], 1, ""),
    (["Courage is found in unlikely places"], 0, ""),
    (["Even the smallest person can change the course of the future"], -1, ""),
    (
        ["You fool of a bear", "Throw yourself in next time"],
        4,
        "You fool of a bear. Throw yourself in next time.",
    ),
    (
        [
            "Even the smallest bear can change the course of the future",
            "It is not our part to master all the Boots of the world",
            "But to do what is in us for the succour of those years wherein we are set",
            "Uprooting the evil in the fields that we know",
            "So that those who live after may have clean earth to till",
            "What weather they shall have is not ours to rule",
        ],
        3,
        "Even the smallest bear can change the course of the future. It is not our part to master all the Boots of the world. But to do what is in us for the succour of those years wherein we are set.",
    ),
]


def test(input_sentences, input_n, expected_output):
    print("---------------------------------")
    print(f"Inputs:")
    print(f" * sentences: {input_sentences}")
    print(f" * n: {input_n}")
    print(f"Expecting: {expected_output}")
    result = accumulate_first_sentences(input_sentences, input_n)
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
