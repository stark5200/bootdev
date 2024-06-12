import unittest

from graphics import Maze, Window

class Tests(unittest.TestCase):
    def test_maze_create_cells(self):
        num_rows = 8
        num_cols = 10
        m1 = Maze(0, 0, num_rows, num_cols, 10, 10, Window(0, 0)) 
        self.assertEqual(len(m1.cells), num_rows)
        self.assertEqual(len(m1.cells[0]), num_cols)

if __name__ == "__main__":
    unittest.main()