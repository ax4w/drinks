# Drinks

A simple website that displays a random drink with instructions and ingredients.

It features a very basic UI with small animations.

The data is from [TheCocktailDB](https://www.thecocktaildb.com/api.php)

You can try it out [here](https://drinks.ax4w.me) :)

## How to run

### Docker
1. Build the `Dockerfile`
2. Run the Image and map a port to `8080`
3. Visit `localhost:<port>`

### Manually
1. Install [templ](https://github.com/a-h/templ)
2. Run `templ generate`
3. Run `go run .`
4. Visit `localhost:8080`