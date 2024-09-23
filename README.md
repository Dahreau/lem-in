# lem-in

## Project Overview
`lem-in` is a pathfinding algorithm project designed to find the optimal way to move ants through a network of rooms connected by tunnels. The goal is to move all ants from the start room to the end room in the least number of turns possible.

## Features
- Efficient pathfinding using Breadth-First Search (BFS) algorithm.
- Usage of Edmonds-Karp algorithm to distribute the ants in the differents paths.
- Handles multiple paths and optimizes the flow of ants.
- Visual representation of the network and ant movements.

## Installation
To install and run the project, follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/dahreau/lem-in.git
    ```
2. Navigate to the project directory:
    ```sh
    cd lem-in
    ```
3. Build the project:
    ```sh
    go build -o lem-in
    ```

## Usage
To run the program, use the following command:
```sh
./lem-in <file_name>
```
Replace `<file_name>` with the path to your file.