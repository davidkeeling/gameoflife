# Conway's Game of Life

## Go terminal implementation

Start the application:

```
go run main.go
```

The app starts with a random board in static mode.

Live cells have three possible colors:
- Cells with 2 or 3 neighbors (will live to next generation) are blue
- Cells with 0 or 1 neighbors (underpopulated) are cyan
- Cells with more than 3 neighbors (overcrowded) are yellow

### Static Mode controls:

- __Mouse click:__ Toggles a cell
- __Delete:__ Clears the board
- __Enter:__ Moves forward 1 generation
- __Space bar:__ Starts play mode

### Play Mode controls:

- __Up arrow:__ Decreases time between generations by a factor of .9, to a minimum of 10ms
- __Down arrow:__ Increases time between generations by a factor of 1.1, to a maximum of 1s
- __Space bar:__ Returns to static mode
