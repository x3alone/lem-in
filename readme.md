<h2 align="center">Lem-in</h2>

This project involves recreating an optimized ant colony simulation in Go. The program reads a map consisting of rooms and tunnels from a file and determines the most efficient way to move a given number of ants from the start room to the end room. The challenge lies in finding all valid paths while optimizing the movement of ants to avoid traffic jams and ensure minimal turns.

## Features

- Parses a file containing rooms, tunnels, and ant information.
- Finds all valid paths between the start and end rooms.
- Simulates and outputs ant movement, adhering to the rules of the colony.
- Handles errors gracefully, with clear messages for invalid input.
- Efficiently solves large maps with thousands of rooms.

## Prerequisites

- [Go](https://go.dev/doc/install) must be installed on your machine.

## Example Input
```zsh
3
##start
0 1 0
##end
1 9 2
2 5 0
0-2
2-1
```
## Example Input
```zsh
$ go run . map.txt
3
##start
0 1 0
##end
1 9 2
2 5 0
0-2
2-1

L1-2
L1-1 L2-2
L2-1
$
```
### Clone the repository

```zsh
git clone https://learn.zone01oujda.ma/git/ahssaini/lem-in
cd lem-in
```
## Authors

- [Mahfoud EL Bachiri](https://learn.zone01oujda.ma/git/melbachi)
- [Abdelhafid Hssaini](https://learn.zone01oujda.ma/git/ahssaini)
- [Ndiasse Dieye](https://learn.zone01oujda.ma/git/ndieye)
- [Otmane Chouari](https://learn.zone01oujda.ma/git/ochouari)
