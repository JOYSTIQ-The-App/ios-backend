-- Posts
select * from posts;

SELECT 
    posts.*, 
    CASE WHEN liked_posts.post_id IS NOT NULL THEN TRUE ELSE FALSE END as user_liked 
FROM posts 
LEFT JOIN liked_posts ON posts.id = liked_posts.post_id AND liked_posts.user_id = 11 
WHERE posts.user_id IN (13,12,2,3,4) 
ORDER BY created_at DESC;

SELECT 
    posts.*, 
    CASE WHEN liked_posts.post_id IS NOT NULL THEN TRUE ELSE FALSE END as user_liked 
FROM posts 
LEFT JOIN liked_posts ON posts.id = liked_posts.post_id AND liked_posts.user_id = 11 
WHERE 
    posts.id < 19 AND
    posts.user_id IN (13,12,2,3,4) 
ORDER BY created_at DESC