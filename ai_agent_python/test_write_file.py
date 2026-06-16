# testing write_file in functions/write_file.py

import unittest
from functions.write_file import write_file

class TestWriteFile(unittest.TestCase):

    def test_1(self) -> None:
        result = write_file("calculator", "lorem.txt", "wait, this isn't lorem ipsum")
        print(result)
        self.assertIn("28 characters written", result)
        print("test 1 done.")
        print("-----------------------------------------")
        print("")
        
    def test_2(self) -> None:
        result = write_file("calculator", "pkg/morelorem.txt", "lorem ipsum dolor sit amet")
        print(result)
        self.assertIn("26 characters written", result)
        print("test 2 done.")
        print("-----------------------------------------")
        print("")

    def test_3(self) -> None:
        result = write_file("calculator", "/tmp/temp.txt", "this should not be allowed")
        print(result)
        self.assertIn("Error:", result)
        print("test 3 done.")
        print("-----------------------------------------")
        print("")

if __name__ == "__main__":
    unittest.main()