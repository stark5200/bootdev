import random
import os
import matplotlib.pyplot as plt

# Funktion zum Setup des Spielfelds
def setup():
    size = 52
    grid = [[0] * size for _ in range(size)]
    # Wählt eine zufällige Position für den ersten kranken Menschen (Infektionstag 1)
    x, y = random.randint(0, size-1), random.randint(0, size-1)
    grid[x][y] = 1  # Infizierter Mensch
    return grid

# Funktion, um das Spielfeld auszugeben
def output(grid, day):
    print(f"Tag {day}")
    for row in grid:
        print(" ".join(str(cell) for cell in row))

# Funktion, die das Spielfeld für jeden Tag aktualisiert
def update(grid, p):
    size = len(grid)
    new_grid = [row[:] for row in grid]  # Kopie des Spielfeldes für die Aktualisierung

    for x in range(size):
        for y in range(size):
            if grid[x][y] == 1:  # Infizierter Mensch
                # Nachbarzellen infizieren, falls sie gesund sind und mit Wahrscheinlichkeit p
                for dx in [-1, 1]:
                    for dy in [-1, 1]:
                        nx, ny = x + dx, y + dy
                        if 0 <= nx < size and 0 <= ny < size and grid[nx][ny] == 0:
                            if random.random() < p:
                                new_grid[nx][ny] = 1
                # Infizierte Person wird zu einem genesen (8) nach einem Tag
                if grid[x][y] < 8:
                    new_grid[x][y] += 1
    return new_grid

# Funktion zum Zählen der gesunden, infizierten und genesenen Personen
def count(grid):
    healthy, infected, recovered = 0, 0, 0
    for row in grid:
        for cell in row:
            if cell == 0:
                healthy += 1
            elif 1 <= cell <= 7:
                infected += 1
            elif cell == 8:
                recovered += 1
    return healthy, infected, recovered

# Funktion zur Ausführung der Simulation
def run_simulation(tEnd, p):
    grid = setup()
    healthy_count, infected_count, recovered_count = [], [], []

    for day in range(tEnd + 1):
        output(grid, day)
        healthy, infected, recovered = count(grid)
        healthy_count.append(healthy)
        infected_count.append(infected)
        recovered_count.append(recovered)
        grid = update(grid, p)

    # Erstelle eine Visualisierung der Simulationsergebnisse
    plt.plot(range(tEnd + 1), healthy_count, label="Gesund")
    plt.plot(range(tEnd + 1), infected_count, label="Infiziert")
    plt.plot(range(tEnd + 1), recovered_count, label="Genesen")
    plt.xlabel("Tage")
    plt.ylabel("Anzahl")
    plt.title("Simulation der Ausbreitung einer Krankheit")
    plt.legend()
    plt.savefig("cx_out/plot.png")  # Speichert die Grafik als PNG-Datei

# Start der Simulation
if __name__ == "__main__":
    os.makedirs("cx_out", exist_ok=True)
    p = 0.25  # Wahrscheinlichkeit, dass ein gesunder Mensch sich infiziert
    tEnd = 10  # Anzahl der Tage der Simulation
    run_simulation(tEnd, p)
