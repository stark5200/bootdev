# Maze project
from graphics import Window, Point, Line, Cell, Maze
        

def main():
    win = Window(800, 600)
    
    maze = Maze(100, 50, 5, 5, 20, 20, win, 67)
    maze = Maze(250, 50, 5, 5, 20, 20, win, 7)
    maze = Maze(400, 50, 5, 5, 20, 20, win, 11)
    maze = Maze(100, 200, 5, 5, 20, 20, win, 2)
    maze = Maze(250, 200, 5, 5, 20, 20, win, 5)
    maze = Maze(400, 200, 5, 5, 20, 20, win, 21)
    # maze.create_cells()
    # maze.break_walls_r(0, 0)
    
    win.wait_for_close()
    
    
main()
        

    
  
    