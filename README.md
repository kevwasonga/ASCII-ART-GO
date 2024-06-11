# ASCII-ART-GO
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII

## ASCII Banner Printer
## Overview

The ASCII Banner Printer is a Go program that takes input text as a string and converts it into a banner-like format using predefined ASCII art characters from a specified text file. The program reads the banner definitions from a .txt file and prints the input strings in the corresponding ASCII art style.
Features:

    * Reads banner definitions from a .txt file.
    * Converts input text into ASCII art characters.
    * Supports multiline input through escaped newline characters (\n).

## Prerequisites

    Go 1.16 or higher

## Installation

    Clone the repository or download the code:
## Usage

    Ensure you have a banner definition file (e.g., standard.txt) in the same directory as the program. This file should contain the ASCII art definitions for characters.

    Run the program with the default standard.txt file:
    go run . <text string>
    
## Code Structure

    main.go: The main file containing the code logic.
    ascii: Package containing the core functionality.

## Error Handling

The program includes basic error handling for:

    * Invalid file names or paths.
    * Errors during file reading.
    * Unsupported characters not defined in the banner file.

## Source(s)
Ascii manual

## Other Innformation
This project is maintained by 

@KevinWasonga
@kevwasonga