from main import *

run_cases = [
    (
        [
            ("New York", "London"),
            ("New York", "Cairo"),
            ("New York", "Tokyo"),
            ("London", "Dubai"),
        ],
        "New York",
        ["New York", "Cairo", "London", "Tokyo", "Dubai"],
    ),
    (
        [
            ("New York", "London"),
            ("New York", "Cairo"),
            ("New York", "Tokyo"),
            ("London", "Dubai"),
            ("Cairo", "Kyiv"),
            ("Cairo", "Madrid"),
            ("London", "Madrid"),
            ("Buenos Aires", "New York"),
            ("Tokyo", "Buenos Aires"),
            ("Kyiv", "San Francisco"),
        ],
        "New York",
        [
            "New York",
            "Buenos Aires",
            "Cairo",
            "London",
            "Tokyo",
            "Kyiv",
            "Madrid",
            "Dubai",
            "San Francisco",
        ],
    ),
]
submit_cases = run_cases + [
    (
        [
            ("Salt Lake City", "Sacramento"),
            ("Boise", "Chicago"),
            ("Seattle", "Chicago"),
            ("Boise", "Salt Lake City"),
        ],
        "Salt Lake City",
        ["Salt Lake City", "Boise", "Sacramento", "Chicago", "Seattle"],
    ),
]


def test(edges_to_add, starting_at, expected_visited):
    print("=================================")
    graph = Graph()
    for edge in edges_to_add:
        graph.add_edge(edge[0], edge[1])
        print(f"Added edge: {edge}")
    print("---------------------------------")
    try:
        bfs = graph.breadth_first_search(starting_at)
        for i, v in enumerate(bfs):
            print(f"Visiting: {v}")
            print(f"Expecting: {expected_visited[i]} \n")

        if bfs == expected_visited:
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
