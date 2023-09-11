-- Inserting data into the 'users' table
INSERT INTO "users" ("email", "username", "role", "created_at", "updated_at")
VALUES 
('admin@joystiq.com', 'admin', 'admin', NOW(), NOW()),
('org1@joystiq.com', 'org1', 'org', NOW(), NOW()),
('user1@joystiq.com', 'user1', 'user', NOW(), NOW()),
('user2@joystiq.com', 'user2', 'user', NOW(), NOW());

-- Inserting data into the 'posts' table
-- Assuming users' IDs are generated in order: 1, 2, 3, 4
INSERT INTO "posts" ("user_id", "s3_key", "title", "game", "body", "status", "likes", "created_at", "updated_at")
VALUES 
(3, 1001, 'First Post', 'Super Mario', 'This is my first post about Super Mario!', 'live', 10, NOW(), NOW()),
(3, 1002, 'Second Post', 'Zelda', 'I love playing Zelda so much!', 'live', 8, NOW(), NOW()),
(4, 1003, 'User2 First Post', 'Minecraft', 'Minecraft is so relaxing!', 'reported', 15, NOW(), NOW()),
(2, 1004, 'Organization Announcement', 'N/A', 'We have a big event coming up!', 'live', 5, NOW(), NOW());