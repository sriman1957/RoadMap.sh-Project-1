# Number Guessing Game (Go)

A simple CLI-based Number Guessing Game built using Go.  
This project was developed as part of practice exercises to improve Go programming fundamentals and CLI application structure.

---

## Project Overview

The computer randomly selects a number between 1 and 100.  
The player selects a difficulty level which determines the number of chances available.  
The player must guess the correct number within the given attempts.

The game provides feedback after each guess:
- If the guess is too high
- If the guess is too low
- If the guess is correct

The game also supports replay functionality.

---

## Features

- Random number generation (1–100)
- Three difficulty levels:
  - Easy (10 chances)
  - Medium (5 chances)
  - Hard (3 chances)
- Input validation for:
  - Non-numeric values
  - Out-of-range numbers
- Attempts counter
- Remaining chances display
- Game Over handling
- Replay option
- Clean function-based architecture

---

## Project Structure

- `getRandomNumber()`  
  Generates a random number between 1 and 100.

- `selectDifficulty()`  
  Displays difficulty options and returns allowed chances.

- `playRound(chances int)`  
  Handles the main guessing loop, attempts tracking, and win/lose logic.

- `main()`  
  Controls overall program flow and replay functionality.

---

## How to Run

Make sure Go is installed on your system.

Navigate to the project folder and run:

```bash
go run NumberGuessingGame.go
```

---

## Example Gameplay

```
Welcome to the Number Guessing Game!
Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your guess: 30
Incorrect! The number is greater than 30
Remaining chances: 9
```

---

## Project URL

https://github.com/sriman1957/PRACTICE-PROJECTS/tree/main/golang/RoadMap.sh_Project_1

---

## Author

Sriman  
Go Practice Project
