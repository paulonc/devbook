package repositories

import (
	"api/src/models"
	"database/sql"
)

type (
	PublicationRepository interface {
		CreatePublication(publication models.Publication) (uint64, error)
		GetPublication(publicationID uint64) (models.Publication, error)
		GetPublications(userID uint64) ([]models.Publication, error)
		UpdatePublication(publicationID uint64, publication models.Publication) error
		DeletePublication(publicationID uint64) error
		FindByUser(userID uint64) ([]models.Publication, error)
		Like(publicationID uint64) error
		Unlike(publicationID uint64) error
	}

	publicationRepository struct {
		db *sql.DB
	}
)

func NewPublicationRepository(db *sql.DB) PublicationRepository {
	return &publicationRepository{db}
}

func (p *publicationRepository) CreatePublication(publication models.Publication) (uint64, error) {
	statement, err := p.db.Prepare(
		"INSERT INTO publications (title, content, author_id) VALUES ($1, $2, $3) RETURNING id",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var lastInsertedID uint64
	err = statement.QueryRow(publication.Title, publication.Content, publication.AuthorID).Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

func (p *publicationRepository) GetPublication(publicationID uint64) (models.Publication, error) {
	var publication models.Publication

	row, err := p.db.Query(`
		SELECT p.id, p.title, p.content, p.author_id, p.likes, p.created_at, u.nick AS author_nick
		FROM publications p
		INNER JOIN users u ON u.id = p.author_id
		WHERE p.id = $1`, publicationID)
	if err != nil {
		return publication, err
	}
	defer row.Close()

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return publication, err
		}
	}
	return publication, nil
}

func (p *publicationRepository) GetPublications(userID uint64) ([]models.Publication, error) {
	rows, err := p.db.Query(`
		SELECT DISTINCT p.*, u.nick 
		FROM publications p 
		INNER JOIN users u ON u.id = p.author_id 
		LEFT JOIN followers f ON p.author_id = f.user_id 
		WHERE u.id = $1 OR f.follower_id = $2
		ORDER BY p.id desc
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if err := rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (p *publicationRepository) UpdatePublication(publicationID uint64, publication models.Publication) error {
	statement, err := p.db.Prepare("UPDATE publications SET title = $1, content = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publicationID); err != nil {
		return err
	}

	return nil
}

func (p *publicationRepository) DeletePublication(publicationID uint64) error {
	statement, err := p.db.Prepare("DELETE FROM publications WHERE id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

func (p *publicationRepository) FindByUser(userID uint64) ([]models.Publication, error) {
	rows, err := p.db.Query(`
        SELECT p.*, u.nick FROM publications p
        JOIN users u ON u.id = p.author_id
        WHERE p.author_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var publications []models.Publication

	for rows.Next() {
		var publication models.Publication

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (p *publicationRepository) Like(publicationID uint64) error {
	statement, err := p.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}

func (p *publicationRepository) Unlike(publicationID uint64) error {
	statement, err := p.db.Prepare(`
        UPDATE publications SET likes =
        CASE
            WHEN likes > 0 THEN likes - 1
            ELSE 0
        END
        WHERE id = $1
    `)
	if err != nil {
		return err
	}

	if _, err = statement.Exec(publicationID); err != nil {
		return err
	}

	return nil
}
