# Check_in_usingMyface
 
Step 1: Install Go

Before you can start building the bot, you need to have Go installed on your system. You can download the latest version of Go from the official website: https://golang.org/dl/

Follow the installation instructions for your operating system to install Go.

Step 2: Install the required dependencies

To build the bot, you need to install the following dependencies:

Telegram Bot API wrapper: go get github.com/go-telegram-bot-api/telegram-bot-api
Gocv: go get -u -d gocv.io/x/gocv
OpenCV: You can follow the official Gocv installation instructions for your operating system: https://gocv.io/getting-started/
Make sure to install OpenCV before installing Gocv.

Step 3: Set up a Telegram bot

To create a new Telegram bot, you need to follow these steps:

Open Telegram and search for the "BotFather" bot.
Start a chat with the BotFather and follow the instructions to create a new bot.
Once you have created the bot, you will receive an HTTP API token. Save this token, as you will need it later to authenticate with the Telegram API.
Step 4: Build the bot

Create a new Go file, for example checkin_bot.go, and copy the code from this GitHub Gist: https://gist.github.com/marlon360/cc8e06f3ba1dd03c557ff9d3fa8c821e

In the code, replace YOUR_TELEGRAM_BOT_TOKEN with the token you received from the BotFather.

Save the file and open a terminal in the directory containing the file.

Build the bot with the following command:

sh

<code>go build checkin_bot.go</code>

This should create an executable file named checkin_bot in the same directory.

Step 5: Run the bot

To run the bot, simply execute the checkin_bot executable file:

sh

<code>./checkin_bot</code>

The bot should now be running and listening for incoming messages.

Step 6: Test the bot

To test the bot, open a chat with the bot in Telegram and send the /checkin command. The bot should respond with a message asking you to send a photo of your face.

Send a photo of your face and the bot should process the photo and respond with a check-in status.

If you're having issues with the bot, make sure to check the console output for error messages.
