# snake-and-ladder

📌 Snakes & Ladders: Problem Statement (Interview-Grade)
Objective
Design and implement a Snakes & Ladders game for two players, where the game continues until one player reaches the last square on the board.

📜 Requirements
1️⃣ Game Board
The board consists of N numbered squares (1 to N), where N is taken as input.


The board contains Snakes (which move the player backward) and Ladders (which move the player forward).


The number and positions of Snakes & Ladders should be configurable via input.


2️⃣ Players & Turns
The game is played between two players (Player A and Player B).


Players take turns rolling a 6-sided die.


If a player lands exactly on the last square (N), they win.


If the roll moves a player beyond the last square, their move is skipped.


3️⃣ Snakes & Ladders Mechanics
If a player lands on the head of a snake, they are moved down to the tail.


If a player lands on the base of a ladder, they are moved up to the top.


4️⃣ Dice Rolling
Players roll a 6-sided die (random number between 1 and 6) on their turn.


The game should print each turn’s roll and resulting position.


5️⃣ Game End Condition
The game ends when one player reaches the last square (N).


The winner is announced.



📌 Required APIs
1️⃣ initGame(N, snakes, ladders)
📌 Input:
N → Total number of squares on the board.


snakes → List of snake positions (head → tail mapping).


ladders → List of ladder positions (base → top mapping).


📌 Output:
Initializes the board with given snakes, ladders, and player positions.



2️⃣ rollDice(player)
📌 Input:
player → The player rolling the dice.


📌 Output:
Generates a random number between 1 and 6.


Moves the player forward accordingly.


Applies snake or ladder effects if applicable.


Prints the new position of the player.


If the player wins, print "Player X wins!" and end the game.



3️⃣ viewBoard()
📌 Output:
Prints the board state, showing:


Player positions.


Snake and ladder locations.



🎯 Example Flow
text
CopyEdit
Game initialized with board size 30.
Snakes: {14 → 7, 25 → 5}
Ladders: {3 → 22, 8 → 26}

Player A rolls a 4 → Moves to 4.
Player B rolls a 6 → Moves to 6.
Player A rolls a 2 → Moves to 6.
Player B rolls a 5 → Moves to 11.
Player A rolls a 3 → Lands on ladder! Moves to 22.
...
Player A rolls a 5 → Moves to 30. 🎉 "Player A wins!"


📌 Additional Considerations
✅ The code must be designed for extensibility (e.g., future multiplayer support, AI).
 ✅ The board should be a separate entity (not hardcoded inside game logic).
 ✅ Random dice rolls should be encapsulated in a method.
 ✅ Thread safety is required (if playing asynchronously).

