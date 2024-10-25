import matplotlib.pyplot as plt

#Case Study

plt.savefig("./cx_out/medikamenten_simulation.png")

dosis_input = float(input("Dosis? "))
frequenz_input = int(float(input("Frequenz? ")))
abbaurate_input = float(input("Abbaurate? "))
zeit_input = input("Zeit? ")

zaehler = [x%(frequenz_input+1) for x in range(50)]
konzentration = dosis_input
z = 0
print("Zähler | Zeit | Konzentration")
for i in range(50):
  print(zaehler[i], "|", z, "|", konzentration) 
  if zaehler[i] == 3:
      z = z + 0.001
      z = round(z, 3) 
      konzentration = konzentration + dosis_input
  else:
      z = z + 1 
      konzentration = konzentration * (1.0-abbaurate_input)
  