# MIDI Compleat
A language made by a programmer that preferes Reaper as a Text Editor and NeoVim as a DAW.

## What is it? What does it do?
**MIDI Compleat**, simply **Compleat** or **MC** for short is an exoteric programming language based on the [brainf\*ck](https://esolangs.org/wiki/Brainfuck) language, but that instead of being written in plain text, it's written using [MIDI](https://en.wikipedia.org/wiki/MIDI) notes.

## Why???
Honestly, I don't know.

I had this idea before going to bed and it just stuck with me.

I probably saw too many brainfuck and code golf videos while making music.

## Learning Objectives
I try to learn something new with every project, and with **Compleat** I tryed to learn a bit of the following:

* **Golang:** I had never used the Go programming language, but  a lot of people had told me that it was a good one. I mostly agree. There are something I miss (i.e. generics and operator overloading), but other than that the language seem good and almost perfect for the kind of programs I like to write.
* **TDD:** Test Driven Development is something that I had never heard of until starting to read [*Learn Go with Tests*](https://quii.gitbook.io/learn-go-with-tests/). I tried to aply it to this project where it seemed suited (mainly on the brainfuck package)
* **MIDI:** I did not learn a lot about MIDI, since I already knew most of what I needed.
* **Brainfuck:** It did what it promisses, it f*cked up my brain.
* **Esolangs:** This project introduced me to a lot of esoteric languages like [bf](https://esolangs.org/wiki/Brainfuck), [Piet](https://esolangs.org/wiki/Piet) and [Shakespeare](https://en.wikipedia.org/wiki/Shakespeare_Programming_Language)

## Stack

* Languages:
    * [**Golang**](https://golang.org/)
* Packages:
    * [**testify/assert**](https://github.com/stretchr/testify)
    * [**midi/reader**](https://gitlab.com/gomidi/midi/-/tree/master/reader)
    * [**midi/writer**](https://gitlab.com/gomidi/midi/-/tree/master/writer)
* Others:
    * [**Brainfuck torture test**](https://github.com/rdebath/Brainfuck)

## Folder structure
The main program is located within the root folder, but I split this project in 2 subfolders:

* **./brainfuck**: This folder handles anything that has to do with the brainf*ck language, and it technically is a standalone interpreter if you make it run as a CLI program.
* **./midi**: This folder handles everything that messes with midi.

## Contributing
I started this project less than a week after starting to learn Golang, so I probably have lots of code that does not follow best practices, or that could, and should, be written in a more elegant and performant way.

If you want to help just submit a *pull request* or an *issue* and I will see what I can do.

## Note:
1. **Compleat** is not being misspelled. I am a huge MTG (Magic: the gathering) fan and in magic Compleation is  *a systemic process with a deliberate end: the creation of a perfect system*.
