import matplotlib.pyplot as plt

#Case Study

plt.savefig("./cx_out/medikamenten_simulation.png")

dosis_input = float(input("Dosis? "))
frequenz_input = int(float(input("Frequenz? ")))
abbaurate_input = float(input("Abbaurate? "))
zeit_input = input("Zeit? ")

zaehler = [x%(frequenz_input+1) for x in range(50)]
zeit = [x for x in range(30)]
konzentration = dosis_input * 1
dosis_counter = 0
print("Zähler | Zeit | Konzentration")
for i in range(30):
  zeit_input = i
  #for i in range (1, len(zaehler)):
  print(zaehler[i], "|", zeit[i], "|", konzentration)  
  konzentration = konzentration * (1.0-abbaurate_input)
  