from tkinter import Tk, BOTH, Canvas, Label

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
    def __init__(self, x1, y1, x2, y2, window):
        # Takes topleft and bottomright coordinates
        # self.cell_size = 10
        
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
            
        

  
      
        