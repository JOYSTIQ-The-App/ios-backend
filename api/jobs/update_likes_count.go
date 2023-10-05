package jobs

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func UpdateLikesCount(db *sql.DB) error {
	// Create the Squirrel statement builder with PostgreSQL dollar format
	var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Get the list of post IDs from the posts table
	var postIds []int
	rows, err := db.Query("SELECT id FROM posts")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var postId int
		if err := rows.Scan(&postId); err != nil {
			log.Println(err)
			return err
		}
		postIds = append(postIds, postId)
	}

	// Iterate over each postId and count its likes then update the posts table
	for _, postId := range postIds {
		// Count likes for the current postId from the liked_posts table
		var likeCount int
		err = db.QueryRow("SELECT COUNT(*) FROM liked_posts WHERE post_id=$1", postId).Scan(&likeCount)
		if err != nil {
			log.Println(err)
			return err
		}

		// Build the UPDATE query to set the likes count for the current postId
		query := sb.Update("posts").Set("likes", likeCount).Where(sq.Eq{"id": postId})

		// Prepare and execute the query
		sqlQuery, args, err := query.ToSql()
		if err != nil {
			log.Println("Prepare query failed", err)
			return err
		}

		_, err = db.Exec(sqlQuery, args...)
		if err != nil {
			log.Println("Execution failed", err)
			return err
		}
	}

	return nil
}
