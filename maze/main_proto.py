# Maze proto project
from graphics import Window, Point, Line, Cell
        

def main():
    win = Window(800, 600)
    point_1 = Point(15, 20)
    point_2 = Point(500, 200)
    line_1 = Line(point_1, point_2)
    
    cell_1 = Cell(30, 30, 50, 50, win)
    cell_1.has_bot_wall = False
    cell_1.draw()
    cell_2 = Cell(60, 30, 80, 50, win)
    cell_2.has_left_wall = False
    cell_2.draw()
    cell_3 = Cell(90, 30, 110, 50, win)
    cell_3.has_top_wall = False
    cell_3.draw()
    cell_4 = Cell(120, 30, 140, 50, win)
    cell_4.has_right_wall = False
    cell_4.draw()
    
    cell_1.draw_move(cell_2, True)
    cell_3.draw_move(cell_4)
    
    cell_5 = Cell(30, 60, 50, 80, win)
    cell_5.has_top_wall = False
    cell_5.has_right_wall = False
    cell_5.draw()
    cell_6 = Cell(60, 60, 80, 80, win)
    cell_6.has_bot_wall = False
    cell_6.has_top_wall = False
    cell_6.draw()
    cell_7 = Cell(90, 60, 110, 80, win)
    cell_7.has_right_wall = False
    cell_7.has_left_wall = False
    cell_7.has_top_wall = False
    cell_7.draw()
    cell_8 = Cell(120, 60, 140, 80, win)
    cell_8.has_bot_wall = False
    cell_8.has_right_wall = False
    cell_8.has_left_wall = False
    cell_8.has_top_wall = False
    cell_8.draw()
    
    win.draw_line(line_1, "Red")
    win.wait_for_close()
    
    
main()
        

    
  
    