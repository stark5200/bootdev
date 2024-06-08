# Maze project
from graphics import Window, Point, Line
        

def main():
    win = Window(800, 600)
    point_1 = Point(15, 20)
    point_2 = Point(500, 200)
    line_1 = Line(point_1, point_2)
    win.draw_line(line_1, "Red")
    win.wait_for_close()
    
    
main()
        

    
  
    