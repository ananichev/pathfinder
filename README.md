# A* search algorithm
In computer science, A* (pronounced as "A star") is a computer algorithm that is widely used in pathfinding and graph traversal, the process of plotting an efficiently directed path between multiple points, called nodes.

## Setup
```
git clone git@github.com:ananichev/pathfinder.git && cd pathfinder
```
### Usage
Make sure you have `field` text file in the running directory.
Content of this file describes search field. It should have:

`1` - start node [required].

`2` - destination node [required].

`X` - obstacle [optional].

For example:

```
. . . . X X X . . 2 .
. . . . . . . X X . .
. . X X X X . . . . .
. . . . . . X X X X X
X X X X X . X . . . .
. . . . X . X . X X X
. X X . X . X . . . .
. . X . X . X X X . X
1 . X . . . . . . . .
```

NOTE! You can use any character as empty node(`.` in the example above)

#### Run
```
go run main.go
```

If the algorithm find right way it will print it to std out. Like so:

```
. . . . X X X . . 2 .
. . + + + + . X X + .
. + X X X X + + + . .
. . + + + . X X X X X
X X X X X + X . . . .
. + + . X + X . X X X
+ X X + X + X . . . .
+ . X + X + X X X . X
1 . X . + . . . . . .
```
Otherwise you will get the next message:

`Can not find path to Node{X: <destX>, Y: <destY>}`
