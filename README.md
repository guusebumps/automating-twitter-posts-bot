# Twitter Bot

A bot built with Go to automate Twitter posts.

## Features
- Automatically posts tweets.
- Easy integration with the Twitter API.

## Installation
1. Clone this repository:
   ```bash
   git clone https://github.com/guusebumps/automating-twitter-posts-bot.git
   cd twitter-bot
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set your Twitter API keys in the `main.go` file.

## Usage
Run the bot:
```bash
go run main.go
```

## Dependencies
- [Go-Twitter](https://github.com/dghubble/go-twitter)
- [OAuth1](https://github.com/dghubble/oauth1)
