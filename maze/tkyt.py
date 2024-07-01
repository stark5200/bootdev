import tkinter as tk

window = tk.Tk()

window.geometry("500x500")
window.title("My yt Gui")

#place objects with pack, grid or place
label = tk.Label(window, text="Hello World!", font=("Arial", 18))
label.pack(padx=20, pady=20)

textbox = tk.Text(window, height=3, font=("Arial", 16))
textbox.pack(padx=10)

my_entry = tk.Entry(window) # entry is Text box with height 1
my_entry.pack() 

button = tk.Button(window, text="Click Me!", font=("Arial", 18))
button.pack(padx=10, pady=10)

buttonframe = tk.Frame(window)
buttonframe.columnconfigure(0, weight=1)
buttonframe.columnconfigure(1, weight=1)
buttonframe.columnconfigure(2, weight=1)

btn1 = tk.Button(buttonframe, text="1", font=("Arial", 18)) # first param is master usually root or window, in this case the button frame
btn1.grid(row=0, column=0, sticky="news") # news = tk.N+tk.E+tk.W+tk.S north east west south

btn2 = tk.Button(buttonframe, text="2", font=("Arial", 18))
btn2.grid(row=0, column=1, sticky="news") 

btn3 = tk.Button(buttonframe, text="3", font=("Arial", 18))
btn3.grid(row=0, column=2, sticky="news") 

btn4 = tk.Button(buttonframe, text="4", font=("Arial", 18))
btn4.grid(row=1, column=0, sticky="news") 

btn5 = tk.Button(buttonframe, text="5", font=("Arial", 18))
btn5.grid(row=1, column=1, sticky="news") 

btn6 = tk.Button(buttonframe, text="6", font=("Arial", 18))
btn6.grid(row=1, column=2, sticky="news") 

buttonframe.pack(fill="x")

anotherbtn = tk.Button(window, text="Test")
anotherbtn.place(x=200, y=200, height=100, width=100)

window.mainloop()

####################

class  MyGUI:
  def __init__(self):
    
    self.root = tk.Tk()
    
    self.label = tk.Label(self.root, text="TOOOOOOOR", font=("Arial", 18))
    self.label.pack(padx=10, pady=10)
    
    self.textbox = tk.Text(self.root, height=5, font=("Arial", 16))
    self.textbox.pack(padx=10, pady=10)
    
    self.check = tk.Checkbutton(self.root, text="Show Message", font=("Arial", 16))
    self.check.pack(padx=10, pady=10)
    
    self.root.mainloop()