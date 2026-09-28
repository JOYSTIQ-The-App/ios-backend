-- Users
INSERT INTO "users" ("email", "username", "role") VALUES
('joystiq@example.com', 'JOYSTIQ', 'admin'),
('apical@example.com', 'Apical', 'gamer'),
('cotts@example.com', 'Naratocotto', 'gamer'),
('nooch@example.com', 'Nooch', 'gamer'),
('livi@example.com', 'Livi', 'gamer'),
('matt@example.com', 'Matt', 'gamer'),
('t@example.com', 'T', 'gamer'),
('jc@example.com', 'JC', 'gamer');

-- Follows
INSERT INTO "follows" ("following_user_id", "followed_user_id", "created_at") VALUES
(2, 1, now()),
(3, 1, now()),
(4, 1, now()),
(5, 1, now()),
(6, 1, now()),
(7, 1, now()),
(8, 1, now());

INSERT INTO "avatars" ("user_id", "avatar_s3_key", "environment") VALUES
(1, '7D0A0E2D-D6D4-4595-9EAC-A7530694100E', 'gameroom');
