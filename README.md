# Installation & Usage

Maze renderer requires Go to build.

```shell
$ go get github.com/bbuck/mazes
$ mazes -h
```

# Example

Using the program you can generate a maze like this (png format):

```shell
$ mazes --rows 15 --cols 25 --cell-width 20 --cell-height 20
```

![rendered_maze.png](https://github.com/bbuck/mazes/raw/master/rendered_maze.png)
