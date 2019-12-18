USE votebotdb;
INSERT INTO songs(link, cover, timestamp, title, artist)        VALUES("https://www.youtube.com/watch?v=1-fgE1jIVuk", "covers/1.jpg", CURRENT_TIMESTAMP, "Grande", "Bruno Mars");
INSERT INTO songs(link, cover, timestamp, title, artist, album) VALUES("https://www.youtube.com/watch?v=kr8wPkdHFA0", "covers/2.jpg", CURRENT_TIMESTAMP, "All of me", "John Legend", "Love In The Future");
INSERT INTO songs(link, cover, timestamp, title, album)         VALUES("https://www.youtube.com/watch?v=1A3kf7dq_14", "covers/3.jpg", CURRENT_TIMESTAMP, "Mr blue sky", "Out of the Blue");

INSERT INTO users (discordID, username) VALUES ("JonShard", "191561233755799554");
INSERT INTO users (discordID, username) VALUES ("Ronald",   "346798766789086788");
INSERT INTO users (discordID, username) VALUES ("Hans",     "238407932238799118");
