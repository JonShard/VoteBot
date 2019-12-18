CREATE DATABASE votebotdb;
USE votebotdb;

CREATE TABLE songs
(
    ID          SMALLINT     NOT NULL AUTO_INCREMENT,
    link        varchar(256) NOT NULL,
    cover       varchar(128) NOT NULL, -- Path to image file.
    timestamp   datetime     NOT NULL,
    title       varchar(64)  NOT NULL,
    artist      varchar(64),
    album       varchar(64),
    PRIMARY KEY(ID)
);

CREATE TABLE users
(
    ID          SMALLINT NOT NULL AUTO_INCREMENT,
    discordID   char(32) NOT NULL,
    username    char(32) NOT NULL,
    PRIMARY KEY(ID)
);

CREATE TABLE votes
(
    ID          SMALLINT NOT NULL AUTO_INCREMENT,
    iteration   SMALLINT NOT NULL, -- Increments by one every time !openVotes is called by a songMaster.
    userID      SMALLINT NOT NULL,
    songID      SMALLINT NOT NULL,
    timestamp   datetime NOT NULL,
    PRIMARY KEY(ID),
    FOREIGN KEY(userID) REFERENCES users (ID),
    FOREIGN KEY(songID) REFERENCES songs (ID)
);

