# Simple Cards Project using Golang

This is a simple Go project that simulates a deck of cards. The project includes functionality to create a new deck, save it to a file, load it from a file, and perform basic operations on the deck.

## Features

- Create a new deck of cards.
- Save the deck to a file.
- Load the deck from a file.
- Unit tests to ensure the functionality works as expected.

## Project Structure

- `main.go`: Contains the main logic for the Cards project.
- `deck.go`: Contains the implementation of the `Deck` type and its associated methods.
- `deck_test.go`: Contains unit tests for the `Deck` type.

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/navneetujjain/Cards.git
   cd Cards
   ```

2. Run the program
   ```bash
   go run main.go
   ```
3. Run the tests
   ```bash
   go test
   ```
## Unit Tests
The project includes unit tests to verify the functionality of the Deck type. The tests are located in the deck_test.go file and include:

- TestNewDeck: Ensures the deck is created with the correct number of cards and the correct first and last cards.
- TestSaveLoadFile: Ensures the deck can be saved to and loaded from a file correctly.

### Example Uses
```bash
d := newDeck()
d.shuffle()
d.print()
d.saveToFile("my_deck")
loadedDeck := newDeckFromFile("my_deck")
loadedDeck.print()
```
