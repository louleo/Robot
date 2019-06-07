## Build Instructions  
1. Open a terminal and navigate to the folder contains "explorer.go".
2. Input `go build` in the terminal.
3. Run `./explorer`.

## Using Instructions  
1. Input the range for the landing area as `x y` and press `Enter`. `x` and `y` must be numeric.
2. Input the Explorer landing location as `x y Direction` and press `Enter`. `x` and `y` must be numeric and `Direction` must be one of `N,S,E,W`.
3. Input the Explorer moving commands. The commands must be a string consists of `L`, `R`, `M` in any kind of order without space. Press `Enter` to indicate your commands is ending.
4. Add extra explorers by repeating step 2 and step 3. Please remind a moving commands must be given after an Explorer landing location is given.
5. Press `Enter` or input `Done` to view the results.
6. Input `Exit` to end the program immediately without any results printed out.
