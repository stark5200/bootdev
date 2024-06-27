import tkinter as tk

root = tk.Tk()

root.geometry("500x500")
root.title("My yt Gui")

label = tk.Label(root, text="Hello World!", font=("Arial", 18))
label.pack(padx=20, pady=20)

textbox = tk.Text(root, height=3, font=("Arial", 16))
textbox.pack(padx=10)

my_entry = tk.Entry(root)
my_entry.pack() 

button = tk.Button(root, text="Click Me!", font=("Arial", 18))
button.pack(padx=10, pady=10)

buttonframe = tk.Frame(root)
buttonframe.columnconfigure(0, weight=1)
buttonframe.columnconfigure(1, weight=1)
buttonframe.columnconfigure(2, weight=1)

btn1 = tk.Button(buttonframe, text="1", font=("Arial", 18))

root.mainloop()