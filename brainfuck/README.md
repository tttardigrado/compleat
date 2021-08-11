# Compleat/brainfuck
The brainfuck interpreter for the Compleat esoteric programming language

## Brainfuck to C
Every brainfuck operator has a "equivalent" in C. The following table explains their relationship

| Brainfuck | C                     |
|-----------|-----------------------|
| `-`       | `--(*ptr);`           |
| `+`       | `++(*ptr);`           |
| `<`       | `--ptr;`              |
| `>`       | `++ptr;`              |
| `[`       | `while (*ptr){`       |
| `]`       | `}`                   |
| `.`       | `putchar(*ptr);`      |
| `,`       | `(*ptr) = getchar();` |

## Notes
1. More info on Brainfuck:

    * [Wikipedia](https://en.wikipedia.org/wiki/Brainfuck)
    * [Esolangs](https://esolangs.org/wiki/Brainfuck)
    * [Code Golf](https://code.golf/brainfuck#python)
    * [Basics of Brainfuck](https://gist.github.com/roachhd/dce54bec8ba55fb17d3a)

