# testing get_files_info in functions/get_files_info.py

import unittest
from functions.get_files_info import get_files_info


class TestGetFilesInfo(unittest.TestCase):
    

    def test_1(self) -> None:
        result = get_files_info("calculator", ".")
        print(result)
        self.assertEqual(result.__contains__("Success"), True)

    def test_2(self) -> None:
        result = get_files_info("calculator", "/bin")
        print(result)
        self.assertEqual(result.__contains__("Error"), True)

    def test_3(self) -> None:
        result = get_files_info("calculator", "../")
        print(result)
        self.assertEqual(result.__contains__("Error"), True)

    def test_4(self) -> None:
        result = get_files_info("calculator", "main.py")
        print(result)
        self.assertEqual(result.__contains__("Error"), True)


if __name__ == "__main__":
    unittest.main()