# Maze project
from tkinter import Tk, BOTH, Canvas, Label
class Window:
  
    def __init__(self, width, height):
        root_widget = Tk()
        root_widget.title = "Window 1"
        root_widget.geometry(str(width)+"x"+str(height))
        label = Label(root_widget, text="Maze 1!", font=("Arial", 18))
        label.pack(padx=20, pady=20)
        c_widget = Canvas()
        c_widget.pack()
        is_running = False
        root_widget.mainloop()
        
window1 = Window(400, 400)
        

    
  
    