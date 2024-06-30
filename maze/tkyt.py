import tkinter as tk

window = tk.Tk()

window.geometry("500x500")
window.title("My yt Gui")

label = tk.Label(window, text="Hello World!", font=("Arial", 18))
label.pack(padx=20, pady=20)

textbox = tk.Text(window, height=3, font=("Arial", 16))
textbox.pack(padx=10)

my_entry = tk.Entry(window)
my_entry.pack() 

button = tk.Button(window, text="Click Me!", font=("Arial", 18))
button.pack(padx=10, pady=10)

buttonframe = tk.Frame(window)
buttonframe.columnconfigure(0, weight=1)
buttonframe.columnconfigure(1, weight=1)
buttonframe.columnconfigure(2, weight=1)

btn1 = tk.Button(buttonframe, text="1", font=("Arial", 18))
btn1.grid(row=0, column=0, sticky="news")

window.mainloop()