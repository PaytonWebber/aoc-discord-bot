<a name="readme-top"></a>

<h3 align="center">AocDiscordBot</h3>

  <p align="center">
  A Discord bot designed to monitor and track progress on a private Advent of Code (AOC) leaderboard. It automatically fetches updates and notifies members when someone on the leaderboard earns a new star.
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

The AocDiscordBot is a specialized Discord bot designed to enhance the experience of participating in the Advent of Code (AOC) challenges among a private group or community. Advent of Code is an annual series of programming puzzles that span a variety of skill sets and challenge types

Key Features:

* Automated Updates: The bot automatically fetches the latest scores and achievements from the Advent of Code leaderboard.
* Real-Time Notifications: Members are notified in Discord when someone on the leaderboard earns a new star, keeping everyone up-to-date and engaged.
* User-Friendly Interface: Simple commands and an intuitive interface make it easy for all members, regardless of their technical background, to interact with the bot and stay informed about the leaderboard status.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally. To get a local copy up and running follow these simple example steps.

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
3. Enter your Advent of Code SESSION COOKIE, Leaderboard ID, Discord Bot Token, and the Channel ID you wish to update and monitor in the .env you just created
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

Currently, the bot is mostly autonomous, and will send updates to a specified channel when a player gets a star. You can manually check the leaderboard status using the !leaderboard command.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
