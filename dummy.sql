-- Users
INSERT INTO "users" ("email", "username", "role") VALUES
('user1@email.com', 'generalUser1', 'gamer'),
('user2@email.com', 'generalUser2', 'gamer'),
('user3@email.com', 'generalUser3', 'gamer'),
('user4@email.com', 'generalUser4', 'gamer'),
('user5@email.com', 'generalUser5', 'gamer'),
('user6@email.com', 'generalUser6', 'gamer'),
('user7@email.com', 'generalUser7', 'gamer'),
('user8@email.com', 'generalUser8', 'gamer'),
('org9@email.com', 'orgUser9', 'org'),
('org10@email.com', 'orgUser10', 'org'),
('apical@example.com', 'apical', 'gamer'),
('cotts@example.com', 'naratocotto', 'gamer'),
('nooch@example.com', 'nooch', 'gamer'),
('blackcat@example.com', 'black cat', 'gamer');

-- Follows
INSERT INTO "follows" ("following_user_id", "followed_user_id", "created_at") VALUES
(4, 2, now()),
(4, 3, now()),
(5, 2, now()),
(7, 5, now()),
(7, 6, now()),
(7, 8, now()),
(7, 9, now()),
(7, 10, now()),
(11, 13, now());

-- Content Posts
INSERT INTO "posts" ("user_id", "media", "game", "body", "status", "likes") VALUES
(1, 'none', 'GameA', 'Content for first post by User1', 'live', 7),
(1, 'none', 'GameA', 'Content for second post by User2', 'live', 6),
(1, 'none', 'GameA', 'Content for first post by User3', 'live', 5),
(2, 'none', 'GameG', 'Content for first post by User2', 'live', 4),
(2, 'none', 'GameH', 'Content for second post by User2', 'live', 3),
(2, 'none', 'Valorant', 'Content for post by User2', 'live', 3),
(3, 'none', 'Valorant', 'Content for post by User3', 'live', 3),
(3, 'none', 'Valorant', 'Content for post by User3', 'live', 3),
(5, 'none', 'Valorant', 'Content for post by User5', 'live', 3),
(5, 'none', 'Valorant', 'Content for post by User', 'live', 3),
(6, 'none', 'Valorant', 'Content for post by User6', 'live', 0),
(6, 'none', 'Valorant', 'Content for post by User6', 'live', 0),
(7, 'none', 'Valorant', 'Content for post by User7', 'live', 0),
(7, 'none', 'Valorant', 'Content for post by User7', 'live', 0),
(8, 'none', 'Valorant', 'Content for post by User8', 'live', 0),
(8, 'none', 'Valorant', 'Content for post by User8', 'live', 0),
(9, 'none', 'GameA', 'Content for first post by Org9', 'live', 0),
(9, 'none', 'GameB', 'Content for second post by Org9', 'archived', 0),
(10, 'none', 'GameC', 'Content for first post by Org10', 'reported', 0),
(10, 'none', 'GameD', 'Content for second post by Org10', 'live', 0),
(13, 'none', 'GameD', 'This new val update sucks', 'live', 0),
(13, 'video', 'GameD', 'This isnt gaming content', 'live', 0);

-- Discussion Posts
INSERT INTO "posts" ("user_id", "s3_key", "media", "game", "body", "status", "likes") VALUES
(13, 'RPReplay_Final1646032101.mov', 'video', 'Valorant', 'This is a sick fucking clip', 'live', 0),
(13, 'Valorant 2022.05.17 - 02.10.12.02.DVR_Trim.mov', 'video', 'GameD', 'Check out this clip!', 'live', 0),
(13, 'Valorant 2022.05.19 - 21.12.05.02.DVR_Trim.mov', 'video', 'GameD', 'Check out this clip!', 'live', 0),
(13, 'Valorant 2022.06.14 - 22.36.27.03.DVR_Trim.mov', 'video', 'GameD', 'Check out this clip!', 'live', 0),
(13, 'Valorant 2022.08.20 - 23.00.43.04.DVR_Trim.mov', 'video', 'GameD', 'Check out this clip!', 'live', 0),
(13, 'Valorant 2022.08.20 - 22.51.58.03.DVR_Trim.mov', 'video', 'GameD', 'Check out this clip!', 'live', 0),
(13, 'Valorant 2022.07.28 - 16.43.08.02.DVR_Trim_Trim.mov', 'video', 'GameD', 'Check out this clip!', 'live', 0);




INSERT INTO liked_posts (user_id, post_id) VALUES 
(2, 1),
(3, 1),
(4, 1),
(5, 1),
(6, 1),
(7, 1),
(8, 1),
(1, 2),
(2, 2),
(3, 2),
(4, 2),
(5, 2),
(6, 2),
(1, 3),
(2, 3),
(3, 3),
(4, 3),
(5, 3),
(2, 4),
(3, 4),
(4, 4),
(5, 4),
(2, 5),
(3, 5),
(4, 5),
(1, 6),
(2, 6),
(3, 6),
(1, 7),
(2, 7),
(3, 7),
(1, 8),
(2, 8),
(3, 8),
(1, 9),
(2, 9),
(3, 9),
(1, 10),
(2, 10),
(3, 10);
