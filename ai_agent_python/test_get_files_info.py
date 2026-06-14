# testing get_files_info in functions/get_files_info.py

import unittest
from unittest import result
from functions.get_files_info import get_files_info


class TestGetFilesInfo(unittest.TestCase):
    

    def test_1(self) -> None:
        result = get_files_info("calculator", ".")
        print(result)
        self.assertIn("Result for current directory:", result)
        self.assertIn("- main.py: file_size=", result)
        self.assertIn("bytes, is_dir=False", result)
        self.assertIn("- tests.py: file_size", result)
        self.assertIn("bytes, is_dir=False", result)
        self.assertIn("- pkg: file_size", result)
        self.assertIn("bytes, is_dir=True", result)

    def test_2(self) -> None:
        result = get_files_info("calculator", "pkg")
        print(result)
        self.assertIn("Result for 'pkg' directory:", result)
        self.assertIn("- calculator.py: file_size=", result)
        self.assertIn("bytes, is_dir=False", result)
        self.assertIn("- render.py: file_size", result)
        self.assertIn("bytes, is_dir=False", result)

    def test_3(self) -> None:
        result = get_files_info("calculator", "/bin")
        print(result)
        self.assertIn("Result for '/bin' directory:", result)
        self.assertIn('Error: Cannot list "/bin" as it is outside the permitted working directory', result)
    def test_4(self) -> None:
        result = get_files_info("calculator", "../")
        print(result)
        self.assertIn("Result for '../' directory:", result)
        self.assertIn('Error: Cannot list "../" as it is outside the permitted working directory', result)

if __name__ == "__main__":
    unittest.main()