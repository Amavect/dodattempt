This program analyzes the difference between Object Oriented, Prodedural, Inline-Procedural, and Data Oriented. 

The program represents many structures with a position and velocity, with the position being updated every second for `loops` seconds.

For `count = 1024` and `loops = 100000` Running `go test ./oop ./pro ./inl ./dod -v` gives the following times:
* OOP: 0.75 s
* PRO: 0.72 s
* INL: 0.49 s
* DOD: 0.45 s

A huge speedup, just by passing a slice of structs into a function. It's even slightly faster than inline-procedural. Cache magic!