# This CLI Pokedex was created while following boot.dev's guided project

#### Usage
1. Run either using `go run .` in the root of the project or do `go build` and run the executable.
2. Once running, you have the following commands available:
  - `help` list the available commands with a short usage hint and description.
  - `map` goes through the list of areas provided by the [pokeapi](https://pokeapi.co/) in batches of 20.
  - `mapb` goes back a batch of 20, to previous areas.
  - `explore <area-name>` will tell you what pokemon you can find in that area.
  - `catch <name-of-pokemon>` will try to catch the named pokemon, not guaranteed, you may need to try multiple times.
  - `inspect <name-of-pokemon>` will tell you more details about a pokemon you previously caught.
  - `pokedex` will list all pokemon you previously caught.
  - `exit` ends the program.
