# testing run_python_file in functions/run_python_file.py

import unittest
from functions.run_python_file import run_python_file

class TestRunPythonFile(unittest.TestCase):

    def test_1(self) -> None:
        result = run_python_file("calculator", "main.py")
        print(result)
        self.assertIn("STDOUT:", result)
        print("test 1 done.")
        print("-----------------------------------------")
        print("")
        
    def test_2(self) -> None:
        result = run_python_file("calculator", "main.py", ["3 + 5"])
        print(result)
        self.assertIn('"result": 8', result)
        print("test 2 done.")
        print("-----------------------------------------")
        print("")

    def test_3(self) -> None:
        result = run_python_file("calculator", "tests.py")
        print(result)
        self.assertIn("STDERR:", result)
        print("test 3 done.")
        print("-----------------------------------------")
        print("")
        
    def test_4(self) -> None:
        result = run_python_file("calculator", "../main.py")
        print(result)
        self.assertIn('Cannot execute "../main.py" as it is outside', result)
        print("test 4 done.")
        print("-----------------------------------------")
        print("")
        
    def test_5(self) -> None:
        result = run_python_file("calculator", "nonexistent.py")
        print(result)
        self.assertIn('"nonexistent.py" does not exist', result)
        print("test 5 done.")
        print("-----------------------------------------")
        print("")
        
    def test_6(self) -> None:
        result = run_python_file("calculator", "lorem.txt")
        print(result)
        self.assertIn('"lorem.txt" is not a Python file', result)
        print("test 6 done.")
        print("-----------------------------------------")
        print("")

if __name__ == "__main__":
    unittest.main()