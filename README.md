This program analyzes the difference between Object Oriented, Prodedural, and Data Oriented. 

The program represents many structures with a position and velocity, with the position being updated every second for `loops` seconds.

For `count = 1024` and `loops = 100000` Running `go test ./oop ./pro ./dod -v` gives the following times:
OOP: 0.72 s
PRO: 0.71 s
DOD: 0.41 s

A huge speedup, just by passing a list of structs into a function. Cache magic!