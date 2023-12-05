<a name="readme-top"></a>

<h3 align="center">AocDiscordBot</h3>

  <p align="center">
  A Discord bot designed to monitor and track progress on a private Advent of Code (AOC) leaderboard. It automatically fetches and checks for updates, then notifies members when someone on the leaderboard earns a new star.
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

The AocDiscordBot is a specialized Discord bot designed to enhance the experience of participating in the Advent of Code (AOC) challenges among a private group or community. Advent of Code is an annual series of programming puzzles that span a variety of skill sets and challenge types.

Key Features:

* Automated Updates: The bot automatically fetches the latest stars and scores from the Advent of Code leaderboard.

* Real-Time Notifications: Members are notified in Discord when someone on the leaderboard earns a new star, keeping everyone up-to-date and engaged.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

To add this bot to your own server, follow these steps.

### Prerequisites

You need to create your own Discord app through their [Devloper Portal](https://discord.com/developers/docs/intro)

### Installation

1. Clone the reop

   ```sh
   git clone https://github.com/PaytonWebber/AocDiscordBot.git
   ```

2. Create a .env file at the root of the repo

   ```sh
   touch .env
   ```

3. Enter the following information in the .env file you just created

   ```ini
   SESSION_COOKIE="<YOUR COOKIE>"
   LEADERBOARD_ID="<YOUR LEADERBOARD ID>"
   DISCORD_TOKEN="<YOUR BOT's TOKEN>"
   CHANNEL_ID="<THE CHANNEL YOU WANT THE BOT TO MONITOR>"
   ```

4. Build the project

   ```sh
   go build cmd/bot/main.go
   ```

5. Start the bot

   ```sh
   ./main

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

Currently, the bot is mostly autonomous. You can manually check the leaderboard status using the !leaderboard command.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
