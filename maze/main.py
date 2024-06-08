# Maze project
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
        self.c_widget = Canvas()
        self.c_widget.pack()
        self.is_running = False
        self.root_widget.mainloop()
        
        self.__root = Tk()
        ...
        self.__root.protocol("WM_DELETE_WINDOW", self.close)
        
    def redraw(self):
        self.root_widget.update_idletasks()
        self.root_widget.update()
        
    def wait_for_close(self):
        self.is_running = True
        while self.is_running:
            self.redraw()
            
    def close(self):
        self.is_running = False
        

def main():
    win = Window(800, 600)
    win.wait_for_close()
    
main()
        

    
  
    