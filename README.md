# ASCII Art

A command-line utility that transforms text into ASCII art patterns.

## Description

This program converts input strings into stylized ASCII art text using predefined banner templates. It reads character patterns from banner files and outputs text in various ASCII art styles.

## Features

- Transform input text to ASCII art
- Support for multiple banner styles (standard, shadow, thinkertoy)
- Clean and efficient Go implementation

## Usage

```bash
go run . [STRING] [BANNER]
```

Example:
```bash
go run . "Hello World" standard
```

## Banner Options

- `standard`: Classic ASCII representation
- `shadow`: Text with shadow effect
- `thinkertoy`: Minimalist style using special characters

## Implementation Details

- Reads banner templates from text files
- Handles newline characters in input strings
- Processes each character to create multi-line ASCII output

## Requirements

- Go 1.13 or higher
