import unittest

from graphics import Maze, Window

class Tests(unittest.TestCase):
    def test_maze_create_cells(self):
        num_rows = 8
        num_cols = 10
        m1 = Maze(0, 0, num_rows, num_cols, 10, 10, Window(0, 0)) 
        self.assertEqual(len(m1.cells), num_rows)
        self.assertEqual(len(m1.cells[0]), num_cols)
        
    def test_maze_destroyFL_cells(self):
        num_rows = 8
        num_cols = 10
        m2 = Maze(0, 0, num_rows, num_cols, 10, 10, Window(0, 0)) 
        self.assertEqual(m2.cells[0][0].has_left_wall, False)
        self.assertEqual(m2.cells[-1][-1].has_right_wall, False)

if __name__ == "__main__":
    unittest.main()