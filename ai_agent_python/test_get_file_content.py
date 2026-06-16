# testing get_files_info in functions/get_files_info.py

import unittest
from functions.get_file_content import get_file_content


class TestGetFileContent(unittest.TestCase):

    def test_1(self) -> None:
        result = get_file_content("calculator", "lorem.txt")
        print(f"lorem.txt length: {len(result)}")
        truncated = f"lorem.txt truncated: {'truncated' in result}"
        print(truncated)
        self.assertIn("truncated: True", truncated)
        print("test 1 done.")
        print("-----------------------------------------")
        print("")
        
    def test_2(self) -> None:
        result = get_file_content("calculator", "main.py")
        print(result)
        self.assertIn("def main() -> None:", result)
        print("test 2 done.")
        print("-----------------------------------------")
        print("")

    def test_3(self) -> None:
        result = get_file_content("calculator", "pkg/calculator.py")
        print(result)
        self.assertIn("def _apply_operator(self, operators: list[str], values: list[float])", result)
        print("test 3 done.")
        print("-----------------------------------------")
        print("")

    def test_4(self) -> None:
        result = get_file_content("calculator", "/bin/cat")
        print(result)
        self.assertIn("Error:", result)
        print("test 4 done.")
        print("-----------------------------------------")
        print("")
        

    def test_5(self) -> None:
        result = get_file_content("calculator", "pkg/does_not_exist.py")
        print(result)
        self.assertIn("pkg/does_not_exist.py", result)
        self.assertIn("Error:", result)
        print("test 5 done.")
        print("-----------------------------------------")
        print("")

if __name__ == "__main__":
    unittest.main()