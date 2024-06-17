# Maze project
from graphics import Window, Point, Line, Cell, Maze
        

def main():
    win = Window(900, 600)
    
    maze = Maze(100, 50, 15, 15, 10, 10, win, 67)
    maze = Maze(350, 50, 5, 9, 20, 20, win, 7)
    maze = Maze(600, 50, 2, 5, 20, 20, win, 11)
    maze = Maze(100, 300, 4, 3, 20, 20, win, 2)
    maze = Maze(350, 300, 5, 5, 20, 20, win, 5)
    maze = Maze(600, 300, 5, 5, 20, 20, win, 21)
    # maze.create_cells()
    # maze.break_walls_r(0, 0)
    
    win.wait_for_close()
    
    
main()
        

    
  
    