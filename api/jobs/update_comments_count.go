package jobs

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
)

func UpdateCommentCounts(db *sql.DB) error {
	// Create the Squirrel statement builder with PostgreSQL dollar format
	var sb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	// Get the list of post IDs from the posts table
	var postIds []int
	rows, err := db.Query("SELECT id FROM posts")
	if err != nil {
		log.Println(err)
		return err
	}
	for rows.Next() {
		var postId int
		if err := rows.Scan(&postId); err != nil {
			log.Println(err)
			rows.Close() // Close rows in case of an error
			return err
		}
		postIds = append(postIds, postId)
	}
	rows.Close()

	// Iterate over each postId and count its comments then update the posts table
	for _, postId := range postIds {
		// Count comments for the current postId from the comments table
		var commentCount int
		err = db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id=$1", postId).Scan(&commentCount)
		if err != nil {
			log.Println(err)
			return err
		}

		// Build the UPDATE query to set the comment count for the current postId
		query := sb.Update("posts").Set("comments", commentCount).Where(sq.Eq{"id": postId})

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
