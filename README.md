# JLPT Grammar(for now) FlashCard Image Crawler

A Go CLI tool to fetch and download grammar flashcards for a specified JLPT level (N1 ~ N5) from [JLPT Sensei](https://jlptsensei.com/).

## Features

- **CLI Interface**: Choose the JLPT level with a flag (`--level n3`, etc.)
- **Support Pagination**: Crawls all grammar pages for the chosen level
- **Image Download**: Downloads grammar flashcard
- **Progress Bar**: Show progress of downloads in the terminal

## Usage

### 1. **Build the CLI**

```sh
go build -o jlpt
```

### 2. **Run the CLI**

```sh
# Use any JLPT level are supported
./jlpt grammar --level=n5
```

### 3. **Downloaded Files**

Downloaded flashcard will be saved in the `images/` folder
