# Minecraft_Villager
[![Go Report Card](https://goreportcard.com/badge/github.com/wasd424/Minecraft_Villager)](https://goreportcard.com/report/github.com/wasd424/Minecraft_Villager) [![CodeFactor](https://www.codefactor.io/repository/github/wasd424/minecraft_villager/badge)](https://www.codefactor.io/repository/github/wasd424/minecraft_villager) [![Build Status](https://travis-ci.org/wasd424/Minecraft_Villager.svg?branch=master)](https://travis-ci.org/wasd424/Minecraft_Villager)
## Get your fellow Minecraft villagers coding.
A new esoteric programming language for Minecraft villagers. We all know that Minecraft Villagers only speak the word "_hmmm_." Thats why in this language all commands are versions of the word "hmm."

### This language is capable of:
It is extremely capable if you have the knowledge of a Minecraft villager or if you have bovine blood (Just look at the source and then compare it to COW's interpreter, you'll see).

#### Available commands (Instruction code, Name, Description):
1. hmm		  - Interpreter will search backwards from here to find the corresponding "hmmmmmmm" command and resume execution from there.
2. hmmm		  - Moves current memory position forward.
3. hmmmm	  - Moves current memory position backwards.
4. hmmmmm	  - Executes the value in current memory block as if it was a normal instruction code.
5. hmmmmmm	  - Reads single byte from input if current memory block is 0. Otherwise, prints current memory block value as a character.
6. hmmmmmmm	  - Decreases value of current memory block by 1.
7. hmmmmmmmm	  - Increases value of current memory block by 1.
8. hmmmmmmmmm	  - If current memory block is 0 skip next command and resume after next "hmm". Otherwise, continue with execution.
9. hmmmmmmmmmm	  - Set current memory block value to 0.
10. hmmmmmmmmmmm   - If no value in current register, copy current memory block value into register. If there is a value in current register copy value into current memory block.
11. hmmmmmmmmmmm  - Print value of current memory block as an integer.
12. hmmmmmmmmmmmm - Read interger and put into current memory block.

###### Notice:
The instruction code is actually the number above minus 1. For example, the instruction code for "hmm" is 0.

### License
All code is licensed under the GPL v3 license. See the license file for more info.

### Contribute
All contributions are welcome. Just submit a pull request and I will probably accept it.
