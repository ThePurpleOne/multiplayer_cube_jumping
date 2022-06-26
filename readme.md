# MCP - Multiplayer Cube Jumping

# Multiplayer architecture.
For now, I'am just trying to implement **server based** multiplayer, might implement p2p host later.

## On connect
When the server is started, A client can be launched and will try to connect to the server. 
When the connection is established, the server will send:
- An ID for the client
- The the world Size.
- The first client position
The server will then create an instance of a "player" and add it to a list of player in the world.

## Input format
Each player will send formatted inputs to the server:
MESSAGE : ```"{PLAYER_ID}-{INPUT}"```
INPUT CAN BE:
- "J" for jump
- "L" for left
- "R" for right

## Output format
The server will send formatted outputs to the clients at fixed interval:
MESSAGE : ```"{PLAYER_ID}-{(WIDTH, HEIGHT)}-{COLOR}-{(INITAL_CUBE_POS.x, INITAL_CUBE_POS.y)}"```


## Server
The server will receive each player inputs, process the physic and send back the new position of every cubes.

## Client
The client will send inputs from user and show the cubes sent back by the server at the right positions.


# HOW DOES IT WORK
There is a Server and Client(s).
The server Starts and Creates an empty world.

The Server then listens for new clients.
When a client connects, the server sends initial Data to the client for him to show the world.
It then creates a player and adds it to the world.
The player has :
- an ID
- a Position
- a Velocity
- a Color
- a Size
- Jumping things

A listenner "Thread" is created on the SERVER for each client connected:
The listenner will listen for new inputs from the client and directly modify the player struct.

On the server, at fixed interval, the server will send the new position of every cubes to be drawn on each clients.
