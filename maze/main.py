# Maze project
from graphics import Window, Point, Line, Cell, Maze
        

def main():
    win = Window(800, 600)
    
    maze = Maze(50, 50, 5, 5, 20, 20, win)
    maze.create_cells()
    
    win.wait_for_close()
    
    
main()
        

    
  
    