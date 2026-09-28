-- Users
INSERT INTO "users" ("email", "username", "role") VALUES
('joystiq@example.com', 'JOYSTIQ', 'admin'),
('joystiq_dev@example.com', 'Joystiq_dev', 'gamer'),
('user2@email.com', 'generalUser2', 'gamer'),
('user3@email.com', 'generalUser3', 'gamer'),
('user4@email.com', 'generalUser4', 'gamer'),
('user5@email.com', 'generalUser5', 'gamer'),
('user6@email.com', 'generalUser6', 'gamer'),
('user7@email.com', 'generalUser7', 'gamer'),
('user8@email.com', 'generalUser8', 'gamer'),
('org9@email.com', 'orgUser9', 'org'),
('org10@email.com', 'orgUser10', 'org'),
('apical@example.com', 'Apical', 'gamer'),
('cotts@example.com', 'Cotts', 'gamer'),
('nooch@example.com', 'Nooch', 'gamer'),
('blackcat@example.com', 'black_cat', 'gamer'),
('livi@example.com', 'Livi', 'gamer');

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
(12, 13, now()),
(2, 13, now()),
(2, 12, now()),
(2, 2, now()),
(2, 3, now()),
(2, 4, now()),
(2, 1, now()),
(3, 1, now()),
(4, 1, now()),
(12, 1, now()),
(5, 1, now());


-- Discussion Posts
INSERT INTO "posts" ("user_id", "media", "game", "body", "status", "created_at") VALUES
(2, 'none', 'Valorant', 'This new valorant update sucks', 'live', now() - interval '4 days'),
(2, 'none', 'LOL', 'This new league update sucks', 'live', now() - interval '8 days'),
(2, 'none', 'Apex', 'This new Apex update sucks', 'live', now() - interval '12 days'),
(2, 'none', 'GameG', 'Content for first post by User2', 'live', now() - interval '22 days'),
(2, 'none', 'GameH', 'Content for second post by User2', 'live', now() - interval '2 days'),
(2, 'none', 'Valorant', 'Content for post by User2', 'live', now() - interval '3 days'),
(3, 'none', 'Valorant', 'Content for post by User3', 'live', now() - interval '13 days'),
(3, 'none', 'Valorant', 'Content for post by User3', 'live', now() - interval '4 hours'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '4 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '5 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '6 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '7 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '8 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '9 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '1 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '2 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '3 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '43 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '10 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '11 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '12 minutes'),
(4, 'none', 'Valorant', 'Content for post by User4', 'live', now() - interval '13 minutes'),
(5, 'none', 'Valorant', 'Content for post by User5', 'live', now() - interval '19 minutes'),
(5, 'none', 'Valorant', 'Content for post by User', 'live', now() - interval '46 minutes'),
(6, 'none', 'Valorant', 'Content for post by User6', 'live', now() - interval '17 days'),
(6, 'none', 'Valorant', 'Content for post by User6', 'live', now() - interval '32 minutes'),
(7, 'none', 'Valorant', 'Content for post by User7', 'live', now() - interval '18 minutes'),
(7, 'none', 'Valorant', 'Content for post by User7', 'live', now() - interval '45 minutes'),
(8, 'none', 'Valorant', 'Content for post by User8', 'live', now() - interval '7 days'),
(8, 'none', 'Valorant', 'Content for post by User8', 'live', now() - interval '19 days'),
(9, 'none', 'GameA', 'Content for first post by Org9', 'live', now() - interval '40 minutes'),
(9, 'none', 'GameB', 'Content for second post by Org9', 'archived', now() - interval '6 days'),
(10, 'none', 'GameC', 'Content for first post by Org10', 'reported', now() - interval '18 days'),
(10, 'none', 'GameD', 'Content for second post by Org10', 'live', now() - interval '9 days'),
(13, 'none', 'GameD', 'This new val update sucks', 'live', now() - interval '11 days'),
(13, 'none', 'GameD', 'This isnt gaming content', 'live', now() - interval '33 minutes');

-- Content Posts
INSERT INTO "posts" ("user_id", "s3_key", "thumbnail_s3_key", "media", "game", "body", "status", "created_at") VALUES
(12, '6A1B424B-6CFC-4E6E-A17F-437E4F883E15.mov', 'LogoTransparent.png', 'video', 'Valorant', 'This is a sick fucking clip', 'live', now() - interval '4 days'),
(12, 'AD14CA5C-F8B2-4CF2-B9AE-60763949A847.mov', 'LogoTransparent.png', 'video', 'GameD', 'Check out this clip!', 'live', now() - interval '6 days'),
(12, '3CDE676B-9A0D-432C-A337-E0CA6E86EA86.mov', 'LogoTransparent.png', 'video', 'GameD', 'Check out this clip!', 'live', now() - interval '4 minutes'),
(12, '20169CB6-DED9-43D0-8390-CD5CDA18E9DC.mov', 'LogoTransparent.png', 'video', 'GameD', 'Check out this clip!', 'live', now() - interval '22 days'),
(1, '60774E3E-F4A7-416C-B4A0-491ABA832DF7.mov', 'LogoTransparent.png', 'video', 'GameD', 'Check out this clip!', 'live', now() - interval '23 minutes'),
(1, 'D3D5F29E-F54B-44D5-873D-6D1C60811C0C.mov', 'LogoTransparent.png', 'video', 'GameD', 'Check out this clip!', 'live', now() - interval '44 days'),
(1, '4FA52E66-2BE1-48CD-9220-559C97B09A6B.mov', 'LogoTransparent.png', 'video', 'GameD', 'Check out this clip!', 'live', now() - interval '14 days'),
(1, '9347E76F-3DF4-4826-8118-E60574D3CF5D.mov', 'LogoTransparent.png', 'video', 'GameD', 'Check out this clip!', 'live', now());

-- Liked Posts
INSERT INTO liked_posts (user_id, post_id) VALUES 
(2, 1),
(3, 1),
(4, 1),
(5, 1),
(6, 1),
(7, 1),
(8, 1),
(2, 2),
(3, 2),
(4, 2),
(5, 2),
(6, 2),
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
(2, 6),
(3, 6),
(2, 7),
(3, 7),
(2, 8),
(3, 8),
(2, 9),
(3, 9),
(2, 10),
(3, 10),
(12, 38),
(12, 41),
(12, 37),
(12, 43),
(12, 42),
(12, 39),
(12, 40),
(12, 44);



-- Comments
INSERT INTO "comments" ("user_id", "post_id", "text", "created_at") VALUES
(2, 1, 'I agree with User1!', now() - interval '3 days'),
(2, 2, 'This is even better than the first!', now() - interval '2 days'),
(3, 2, 'Well said, User2!', now() - interval '1 days'),
(3, 3, 'What a post!', now() - interval '5 hours'),
(4, 4, 'User2, always on point!', now() - interval '3 hours'),
(5, 5, 'This game is the best.', now() - interval '2 hours'),
(6, 6, 'Solid post.', now() - interval '1 hours'),
(2, 7, 'I don’t entirely agree here.', now() - interval '45 minutes'),
(3, 8, 'Interesting perspective.', now() - interval '30 minutes'),
(4, 9, 'User5 always has something good to say!', now() - interval '15 minutes'),
(5, 10, 'I’ve been saying this for years!', now() - interval '10 minutes'),
(7, 13, 'Totally disagree with this update too.', now() - interval '5 minutes'),
(8, 14, 'Why is this even here?', now() - interval '4 minutes'),
(9, 15, 'Orgs always have the best content.', now() - interval '3 minutes'),
(10, 16, 'Keep these coming!', now() - interval '2 minutes'),
(11, 17, 'Amazing clip. How do you even do that?', now() - interval '1 minute'),
(12, 18, 'I’ve seen better...', now() - interval '30 seconds'),
(12, 25, 'I’ve seen better...', now() - interval '30 seconds'),
(12, 26, 'I’ve seen better...', now() - interval '30 seconds'),
(12, 27, 'I’ve seen better...', now() - interval '30 seconds');

-- Profiles
INSERT INTO "profiles" ("user_id", "bio", "resume") VALUES
(11, 'JOYSTIQ Founder \nApex gamer, I also like league', 'Solo climbed to Diamond in Apex so yeah, I’m pretty good'),
(2, 'I am a JOYSTIQ developer. This is a bio for my profile page. I wonder how long this could be.\nNew line check', 'We built this app, so yeah were pretty cool and were pretty dammn good developers.');

-- Socials
INSERT INTO "socials" ("user_id", "social", "social_username") VALUES
(11, 'discord', 'apical'),
(11, 'twitch', 'apical'),
(11, 'xbox', 'apical'),
(2, 'xbox', 'Joystiq_dev'),
(2, 'discord', 'Joystiq_dev'),
(2, 'twitch', 'Joystiq_dev');

INSERT INTO "avatars" ("user_id", "avatar_s3_key", "environment") VALUES
(1, '7D0A0E2D-D6D4-4595-9EAC-A7530694100E', 'gameroom');
