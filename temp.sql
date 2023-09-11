SELECT
    p.id,
    p.user_id,
    p.s3_key,
    p.title,
    p.game,
    p.body,
    p.status,
    p.created_at,
    p.updated_at,
    COUNT(lp.post_id) AS likes_count
FROM
    posts p
JOIN
    liked_posts lp ON p.id = lp.post_id
WHERE
    lp.created_at >= '2023-09-10'
GROUP BY
    p.id,
    p.user_id,
    p.s3_key,
    p.title,
    p.game,
    p.body,
    p.status,
    p.created_at,
    p.updated_at
ORDER BY
    likes_count DESC;
