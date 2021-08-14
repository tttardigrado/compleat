# Operators

## (Minus) -

- C equivalent: `--(*p);`

- The value of the current cell is **decremented by 1**

- it wraps at **0** back to **255**

```
[0] ---> [255]
[1] ---> [0]
[2] ---> [1]
[3] ---> [2]
[4] ---> [3]
[5] ---> [4]
[6] ---> [5]
[7] ---> [6]
... ---> ...
[255] ---> [254]
```

## (Plus) +

- C equivalent: `++(*p);`

- The value of the current cell is **incremented by 1**

- it wraps at **255** back to **0**

```
[0] ---> [1]
[1] ---> [2]
[2] ---> [3]
[3] ---> [4]
[4] ---> [5]
[5] ---> [6]
[6] ---> [7]
[7] ---> [8]
... ---> ...
[255] ---> [0]
```

## (Left) <

- C equivalent: `--p;`

- The cursor is **moved back 1 cell**

- If the cursor is **0** it outputs an **error**

```
[0] [1] [2] [3] [4] [5] ---> [0] [1] [2] [3] [4] [5]
         ^                        ^

[0] [1] [2] [3] [4] [5] ---> [0] [1] [2] [3] [4] [5]
             ^                        ^

[0] [1] [2] [3] [4] [5] --->  ERROR
 ^
```

## (Right) >

- C equivalent: `++p;`

- The cursor is **moved up 1 cell**

- If the cursor is **at the edge** a new cell with **null value (0)** is added

```
[0] [1] [2] [3] [4] [5] ---> [0] [1] [2] [3] [4] [5]
 ^                                ^

[0] [1] [2] [3] [4] [5] ---> [0] [1] [2] [3] [4] [5]
         ^                                ^

[0] [1] [2] [3] [4] [5] ---> [0] [1] [2] [3] [4] [5] [0]
                     ^                                ^
```

## (Begin) [

- C equivalent: `while (*p) {`

- while the value of the current cell is **not 0** the operations inside of the loop ( i.e. between [ ] ) will run

- if the value **is 0**, the loop will be skipped/end

## (End) ]

- C equivalent: `}`

- while the value of the current cell is **not 0** the operations inside of the loop ( i.e. between [ ] ) will run, so the cursor will move to the **correspondent [**

- if the value **is 0**, the program will just continue

## (Print) .

- C equivalent: `putchar(*p);`

- the value of the current cell is **printed**

```
...
[10] ---> '\n'
...
[65] ---> 'A'
...
[70] ---> 'F'
...
[80] ---> 'P'
...
```

## (Get) ,

- C equivalent: `*p = getchar();`

- the value of the current cell **set to a value depending on the character** that was provided

```
...
'\n' ---> [10]
...
'A' ---> [65]
...
'F' ---> [70]
...
'P' ---> [80]
...
```
