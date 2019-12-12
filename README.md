# VoteBot 
A discord bot for voting on karaoke songs for VRChat worlds in a discord channel.

## About
VRChat is a PC game for taking talking with other people using with or without a virtual reality headset ([VRChat](https://www.vrchat.com/)). One of the things people get up to in this game is karaoke.
VoteBot is a discord bot that lets people use discord to vote on which songs should be present in the karaoke world. It gathers votes in the form of reactions from a discord text channel and sends SQL queries to a database managing the 

## Usage
VoteBot can be triggered by a user with the proper role.      
| Commands             | Arguments | Required Role | Description                                                                                                                 |  
| -------------------- | --------- | ------------- | --------------------------------------------------------------------------------------------------------------------------- |  
| !openVote            |           | SongMaster    | Posts a list of all available songs in a channel as individual messages and lets user add reactions to the songs they want. |  
| !displayList         |           | KarokeUser    | Post the list of all available songs.                                                                                       |  
| !vote                | int       | KarokeUser    | Vote for the song with id.                                                                                                  |  
| !search              | string    | KarokeUser    | Search for a song where song tile contains the sub-string.                                                                  |  
| !closeVote           |           | SongMaster    | Counts all votes, posts result in the text channel and sends the set number of most voted songs to the data base.           |  
| !setSongLimit        | int       | SongMaster    | Change the number of songs are sent to the server.                                                                          |  
| !setVoteCount        | int       | SongMaster    | Change the number of votes each user can submit.                                                                            |  
| !setPatreonVoteCount | int       | SongMaster    | Change the number of additional votes users with the patreon role gets.                                                     |  

## Roadmap
VoteBot is currently in development.
- [ ] !helloWorld command.
- [ ] !displayList command.
- [ ] !openVote command, with song data from database.
- [ ] !vote command.
- [ ] !closeVote command.
- [ ] !search command.
- [ ] !setSongLimit command. 
- [ ] !setVoteCount command.
- [ ] !setPatreonVoteCount command.

## Installation
``` bash
git clone repo
cd VoteBot
make
# Config file setup.
bin/Votebot
.
.
```