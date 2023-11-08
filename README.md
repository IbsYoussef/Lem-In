## Hello Welcome to lem-in
### Usage Instructions to run visualiser
`go build && cd visualiser && go build && cd ..`

`./lem-in examples/example05.txt | visualiser/visualiser`

Run test file: `go test`

Run test coverage: `go test -coverprofile=coverage.out`

### Algorithm explanation

The text file contains information about the Ants and Colony.

```
10
##start
1 23 3
2 16 7
#comment
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
#another comment
4-2
2-1
7-6
7-2
7-4
6-5
```


```console
        _________________
       /                 \
  ____[5]----[3]--[1]     |
 /            |    /      |
[6]---[0]----[4]  /       |
 \   ________/|  /        |
  \ /        [2]/________/
  [7]_________/
```

`[1]` is the start room, `[0]` the is end room

### Initial Stage, find all passible paths for ants to traverse.

Find the start room. Then check for connections with start room.
In this example: room `[1]` is connected with room `[3]` and `[2]`.
So we created a copy of our map but without these connection. It's should look like this.

Then I will set up a new start room `[3]` for this copy 

Path will be `[1] -> [3]` 

```console
        _________________           
       /                 \
  ____[5]----[3]  [1]     |
 /            |           |
[6]---[0]----[4]          |
 \   ________/|           |
  \ /        [2] _________/
  [7]_________/
```

In this example start room is `[2]` for this copy

Path will be `[1] -> [2]` 

```console
        _________________
       /                 \
  ____[5]----[3]  [1]     |
 /            |           |
[6]---[0]----[4]          |
 \   ________/|           |
  \ /        [2] _________/
  [7]_________/
```

Next step I will try to do the same unttil I no longer have anymore connections to next rooms or I will reach end room `[3]`
Start room `[3]`, connected to rooms `[5]` and `[4]`

Path will be `[1] -> [3] -> [5]` 

Path will be `[1] -> [3] -> [4]` 

```console
        _________________
       /                 \
  ____[5]    [3]  [1]     |
 /                        |
[6]---[0]----[4]          |
 \   ________/|           |
  \ /        [2] _________/
  [7]_________/
```

Start room `[2]`, connected rooms `[5]`, `[4]` and `[7]`

Path will be `[1] -> [2] -> [5]` 

Path will be `[1] -> [2] -> [4]` 

Path will be `[1] -> [2] -> [7]` 

```console
            
                          
  ____[5]----[3]  [1]      
 /            |            
[6]---[0]----[4]           
 \   ________/             
  \ /        [2]           
  [7]          
```

Continue doing the same....

After these operations I will have a list of all possible paths

`[1] -> [3] -> [4] -> [0]`

`[1] -> [2] -> [4] -> [0]`

`[1] -> [3] -> [5] -> [6] -> [0]`

`[1] -> [2] -> [5] -> [6] -> [0]`

`[1] -> [2] -> [7] -> [6] -> [0]`

`[1] -> [2] -> [7] -> [4] -> [0]`

`[1] -> [3] -> [4] -> [7] -> [6] -> [0]`

`[1] -> [3] -> [5] -> [2] -> [4] -> [0]`

`[1] -> [2] -> [5] -> [3] -> [4] -> [0]`

`[1] -> [2] -> [4] -> [7] -> [6] -> [0]`

`[1] -> [3] -> [5] -> [2] -> [7] -> [6] -> [0]`

`[1] -> [3] -> [5] -> [2] -> [7] -> [4] -> [0]`

`[1] -> [3] -> [4] -> [7] -> [2] -> [5] -> [6] -> [0]`

`[1] -> [3] -> [5] -> [2] -> [4] -> [7] -> [6] -> [0]`

`[1] -> [3] -> [5] -> [6] -> [7] -> [2] -> [4] -> [0]`

`[1] -> [2] -> [5] -> [3] -> [4] -> [7] -> [6] -> [0]`

`[1] -> [2] -> [7] -> [6] -> [5] -> [3] -> [4] -> [0]`

`[1] -> [2] -> [7] -> [4] -> [3] -> [5] -> [6] -> [0]` 

In this case I have 18 possible paths for the ants to traverse 

### Find unique combiantion of all passible paths

Just imagine that first path `[1] -> [3] -> [4] -> [0]` is just `1`

So now we need to find all other possible possible short path combinations

`[1], [2], [3], [4], ... [16], [17], [18]`

`[1,2], [1,3], [1,4], ... [1,18], [2,3], [2,4], [2,5], ... [2,17], [2,18], [3,4], [3,5], ... [16,17], [16,18], [17,18]`

`[1,2,3], [1,2,4], [1,2,5], ... [16,17], [15,17,18], [16,17,18]`

`...`

`[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17], [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,18], [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,17,18] ...`

`[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18]`

### Find paths with unqiue rooms. Path combination should have only unique rooms

After this i will find 

[0] -> `[1] -> [3] -> [4] -> [0]`

[0 3] -> `[1] -> [3] -> [4] -> [0]` and `[1] -> [2] -> [5] -> [6] -> [0]`

[0 4] -> `[1] -> [3] -> [4] -> [0]` and `[1] -> [2] -> [7] -> [6] -> [0]`

[1 2] -> `[1] -> [2] -> [4] -> [0]` and `[1] -> [3] -> [5] -> [6] -> [0]`

[2 5] -> `[1] -> [3] -> [5] -> [6] -> [0]` and `[1] -> [2] -> [7] -> [4] -> [0]`

Then i am checking wich path is faster and print the results