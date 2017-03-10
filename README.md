# Conway's Game of Life

## Go terminal implementation

Start the application:

```
go run main.go
```

The app starts with a random board in static mode.

### Static Mode controls:

- __Mouse click:__ Toggles a cell
- __Delete:__ Clears the board
- __Enter:__ Moves forward 1 generation
- __Space bar:__ Starts play mode

### Play Mode controls:

- __Up arrow:__ Decreases time between generations by a factor of .9, to a minimum of 10ms
- __Down arrow:__ Increases time between generations by a factor of 1.1, to a maximum of 1s
- __Space bar:__ Return to static mode
