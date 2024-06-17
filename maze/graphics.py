from tkinter import Tk, BOTH, Canvas, Label
import time, random

class Window:
  
    def __init__(self, width, height):
        self.width = width
        self.height = height
        self.root_widget = Tk()
        self.root_widget.title = "Window 1"
        self.root_widget.geometry(str(self.width)+"x"+str(self.height))
        self.label = Label(self.root_widget, text="Maze 1!", font=("Arial", 18))
        self.label.pack(padx=20, pady=20)
        self.canvas_widget = Canvas(self.root_widget, bg="White", height=self.height, width=self.width)
        self.canvas_widget.pack(fill=BOTH, expand=1)
        self.is_running = False
        #self.root_widget.mainloop()
        print("window opened...")
        self.root_widget.protocol("WM_DELETE_WINDOW", self.close)
        
    def redraw(self):
        self.root_widget.update_idletasks()
        self.root_widget.update()
        
    def wait_for_close(self):
        self.is_running = True
        while self.is_running:
            self.redraw()
        print("window closed...")
            
    def close(self):
        self.is_running = False
        
    # functions related to other classes in the graphics file like Point or Line.
    def draw_line(self, line, fill_color):
        canvas = self.canvas_widget
        line.draw(canvas, fill_color)
        

class Point:
        
    def __init__(self, x, y):
        self.x = x
        self.y = y
        
        
class Line:
        
    def __init__(self, p1, p2):
        # Takes 2 Points as input
        self.p1 = p1
        self.p2 = p2
        
    def draw(self, canvas, fill_color):
        canvas.create_line(self.p1.x, self.p1.y, self.p2.x, self.p2.y, fill=fill_color, width=2)
        
        
class Cell:
    def __init__(self, x1, y1, x2, y2, window=None):
        # Takes topleft and bottomright coordinates
        # self.cell_size = 10
        self.visited = False
        self.has_left_wall = True
        self.has_right_wall = True
        self.has_top_wall = True
        self.has_bot_wall = True
        self.x1 = x1
        self.y1 = y1
        self.x2 = x2
        self.y2 = y2
        self.window = window
        
    def draw(self):
        if self.has_left_wall:
            self.window.canvas_widget.create_line(self.x1, self.y1, self.x1, self.y2, fill="Black", width=2)
            
        if self.has_right_wall:
            self.window.canvas_widget.create_line(self.x2, self.y1, self.x2, self.y2, fill="Black", width=2)
            
        if self.has_top_wall:
            self.window.canvas_widget.create_line(self.x1, self.y1, self.x2, self.y1, fill="Black", width=2)
            
        if self.has_bot_wall:
            self.window.canvas_widget.create_line(self.x1, self.y2, self.x2, self.y2, fill="Black", width=2)
            
    def draw_move(self, to_cell, undo=False):
        color = "Red"
        if undo:
          color = "Gray"
        c1_x = (self.x1 + self.x2) / 2
        c1_y = (self.y1 + self.y2) / 2
        c2_x = (to_cell.x1 + to_cell.x2) / 2
        c2_y = (to_cell.y1 + to_cell.y2) / 2
        
        self.window.canvas_widget.create_line(c1_x, c1_y, c2_x, c2_y, fill=color, width=1)
        
        
class Maze:
    def __init__(
        self,
        x0,
        y0,
        num_rows,
        num_cols,
        cell_size_x,
        cell_size_y,
        win=None,
        seed=None,
    ):
        #self.cells = []
        self.x0 = x0
        self.y0 = y0
        self.num_rows = num_rows
        self.num_cols = num_cols
        self.cell_size_x = cell_size_x
        self.cell_size_y = cell_size_y
        if win:
            self.win = win
        if seed:
            self.r = random.seed(seed)
        print("Creating cells...")
        self.create_cells()
        print("cells created...")
        
        
    def create_cells(self):
        print("empty cells matrix...")
        self.cells = []
        for row in range(self.num_rows):
            self.cells.append([])
            for col in range(self.num_cols):
                x1 = self.x0 + row*self.cell_size_x
                y1 = self.y0 + col*self.cell_size_y
                x2 = self.x0 + (row+1)*self.cell_size_x
                y2 = self.y0 + (col+1)*self.cell_size_y
                print(f"Creating cell[{row, col}]...")
                
                new_cell = Cell(x1, y1, x2, y2, self.win) 
                self.cells[row].append(new_cell)
                
        self.break_entrance_and_exit()
        
        for row in range(self.num_rows):
            for col in range(self.num_cols):
                
                print(f"drawing cell[{row, col}]...")
                self.draw_cell(row, col)
      
    def draw_cell(self, i, j):
        
        current_cell = self.cells[i][j]
        current_cell.draw()
        
        print(f"animating cell[{i, j}]...")
        self.animate()
      
    def animate(self):
        self.win.redraw()
        print(f"sleeping ...")
        time.sleep(0.02)
          
    def break_entrance_and_exit(self):
        first_cell = self.cells[0][0]
        last_cell = self.cells[-1][-1]
        first_cell.has_left_wall = False
        last_cell.has_right_wall = False
    
    def break_walls_r(self, i, j):
        self.cells[i][j].visited = True
        v_list = []
        v_list.extend([i, j])
        

            
        

  
      
        