from main import *
import random

run_cases = [
    (10, 1000, "luxa", 383.9),
    (20, 2000, "luxa", 593.25),
    (30, 3000, "luxa", 932.23),
]

submit_cases = run_cases + [
    (40, 4000, "luxa", 1495.4),
    (80, 8000, "luxa", 2608.95),
    (160, 16000, "luxa", 5920.98),
]


def test(num_handles, avg_aud_size, brand_name, expected_output):
    try:
        print("---------------------------------")
        print(
            f"Checking {num_handles} influencers with average audience sizes of {avg_aud_size}..."
        )
        print(f" * brand_name: {brand_name}")
        print(f"Expecting: {expected_output}")
        all_handles = get_all_handles(num_handles, avg_aud_size)
        avg = round(get_avg_brand_followers(all_handles, brand_name), 2)
        print(f"Actual: {avg}")
        if avg == expected_output:
            print("Pass")
            return True
        print("Fail")
        return False
    except Exception as e:
        print("Fail")
        print(e)
        return False


def get_all_handles(num, audience_size):
    all_handles = []
    for i in range(num):
        m = random.randrange(
            int(audience_size - audience_size * 1.2),
            int(audience_size + audience_size * 1.2),
        )
        handles = get_instagram_handles(m)
        all_handles.append(handles)
    return all_handles


def get_instagram_handles(num):
    handles = []
    for i in range(0, num):
        m = random.randrange(0, 6)
        if m == 0:
            handles.append(f"luxaraygirl{i}")
        elif m == 1:
            handles.append(f"theprimerog{i}")
        elif m == 2:
            handles.append(f"luxafanboi{i}")
        elif m == 3:
            handles.append(f"dashlord{i}")
        elif m == 4:
            handles.append(f"saintrex{i}")
        elif m == 5:
            handles.append(f"writergurl{i}")
    return handles


def main():
    random.seed(1)
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
