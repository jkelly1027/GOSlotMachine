GO Slot Machine:
1) Welcome User and get their name.
2) Ask user for their bet, need to know users balance first. Can't bet more than they have.

SLOT MACHINE
1) Made up of columns
2) Each column will consist of symbols.
3) Each symbol has a particular value.
4) Aim is to get lines of these symbols across the columns.
5) If a line of matching symbols is achieved the bet is multiplied by the value of the matching symbol.

SYMBOLS (The rarer a symbol, the less it occurs in a reel):
A: 4
B: 7
C: 12
D: 20

MULTIPLIER (How much your bet is multiplied by if you get a complete line of a symbol)
A: 20
B: 10
C: 5
D: 2

MACHINE
1) Generate array that contains symbols.
2) Randomly pick an index within that array.
3) Display chosen symbol in column
4) Repeat for 3 columns

Divided the functions between main/spin/utils to experiment with "go mod init..."