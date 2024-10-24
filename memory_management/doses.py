import matplotlib.pyplot as plt

#Case Study

plt.savefig("./cx_out/medikamenten_simulation.png")

dosis_input = float(input("Dosis? "))
frequenz_input = int(float(input("Frequenz? ")))
abbaurate_input = float(input("Abbaurate? "))
zeit_input = input("Zeit? ")

zaehler = [0 for x in range(50)]
zeit_input = [0 for x in range (50)]
konzentration = [0 for x in range(50)]

print("Zähler | Zeit | Konzentration")
for i in range (1, len(zaehler)):
  
  if zaehler[i] == frequenz_input:
    zaehler[i] = 0
  else:
    zaehler[i] = zaehler[i-1] + 1
  print(zaehler[i], " | ", zeit_input[i], " | ", konzentration[i]) 
  