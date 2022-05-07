# MCP - Multiplayer Cube Jumping

# Multiplayer architecture.
For now, I'am just trying to implement server based multiplayer, might implement p2p host later.

## On connect
When the server is started, A client can be launched and will try to connect to the server. When the connection is established, the server will send its ID. The server will then create an instance of a "player" and add it to a list. 

## Input format
Each player will send formatted inputs to the server:
MESSAGE : ```"{PLAYER_ID}-{INPUT}"```
INPUT CAN BE:
- "J" for jump
- "L" for left
- "R" for right

## Output format
The server will send formatted outputs to the clients at fixed interval:
MESSAGE : ```"{PLAYER_ID}-{(x, y)}-{COLOR}-{SIZE}"```


## Server
The server will receive each player inputs, process the physic and send back the new position of every cubes.

## Client
The client will send inputs from user and show the cubes sent back by the server at the right positions.