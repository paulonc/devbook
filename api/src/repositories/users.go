package repositories

import (
	"api/src/models"
	"database/sql"
	"time"
)

type (
	UserRepository interface {
		CreateUser(user models.User) (uint64, error)
		GetUser(id uint64) (models.User, error)
		GetUsers() ([]models.User, error)
		GetUserByEmail(email string) (models.User, error)
		UpdateUser(id uint64, user models.User) error
		DeleteUser(id uint64) error
		FollowUser(userID, followerID uint64) error
		UnfollowUser(userID, followerID uint64) error
		GetFollowers(userID uint64) ([]models.User, error)
		GetFollowing(userID uint64) ([]models.User, error)
		GetPassword(userID uint64) (string, error)
		UpdatePassword(userID uint64, password string) error
	}

	userRepository struct {
		db *sql.DB
	}
)

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id uint64
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userRepository) GetUser(id uint64) (models.User, error) {
	var user models.User

	row, err := u.db.Query("SELECT id, name, nick, email FROM users WHERE id = $1", id)
	if err != nil {
		return user, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email); err != nil {
			return user, err
		}
	}

	return user, nil
}

func (u *userRepository) GetUsers() ([]models.User, error) {
	rows, err := u.db.Query("SELECT id, name, nick, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var id uint64
		var name, nick, email string
		if err := rows.Scan(&id, &name, &nick, &email); err != nil {
			return nil, err
		}
		users = append(users, models.User{
			ID:    id,
			Name:  name,
			Nick:  nick,
			Email: email,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	row, err := u.db.Query("SELECT id, password FROM users WHERE email = $1", email)
	if err != nil {
		return user, err
	}
	defer row.Close()

	if row.Next() {
		if err := row.Scan(&user.ID, &user.Password); err != nil {
			return user, err
		}
	}
	return user, nil

}

func (u *userRepository) UpdateUser(id uint64, user models.User) error {
	statement, err := u.db.Prepare(
		"UPDATE users SET name = $1, nick = $2, email = $3 WHERE id = $4",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Nick, user.Email, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) DeleteUser(id uint64) error {
	statement, err := u.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FollowUser(userID, followerID uint64) error {
	statement, err := u.db.Prepare("INSERT INTO followers (user_id, follower_id) VALUES ($1, $2) ON CONFLICT DO NOTHING")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil

}

func (u *userRepository) UnfollowUser(userID, followerID uint64) error {
	statement, err := u.db.Prepare("DELETE FROM followers WHERE user_id = $1 AND follower_id = $2")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil

}

func (u *userRepository) GetFollowers(userID uint64) ([]models.User, error) {
	rows, err := u.db.Query("SELECT u.id, u.name, u.nick, u.email, u.created_at FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var id uint64
		var name, nick, email string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &nick, &email, &createdAt); err != nil {
			return nil, err
		}
		followers = append(followers, models.User{
			ID:        id,
			Name:      name,
			Nick:      nick,
			Email:     email,
			CreatedAt: createdAt,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func (u *userRepository) GetFollowing(userID uint64) ([]models.User, error) {
	rows, err := u.db.Query("SELECT u.id, u.name, u.nick, u.email, u.created_at FROM users u INNER JOIN followers f ON u.id = f.user_id WHERE f.follower_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var id uint64
		var name, nick, email string
		var createdAt time.Time
		if err := rows.Scan(&id, &name, &nick, &email, &createdAt); err != nil {
			return nil, err
		}
		followers = append(followers, models.User{
			ID:        id,
			Name:      name,
			Nick:      nick,
			Email:     email,
			CreatedAt: createdAt,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func (u *userRepository) GetPassword(userID uint64) (string, error) {
	row, err := u.db.Query("SELECT password FROM users WHERE id = $1", userID)
	if err != nil {
		return "", err
	}
	defer row.Close()

	var password string
	if row.Next() {
		if err := row.Scan(&password); err != nil {
			return "", err
		}
	}
	return password, nil
}

func (u *userRepository) UpdatePassword(userID uint64, password string) error {
	statement, err := u.db.Prepare("UPDATE users SET password = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(password, userID)
	if err != nil {
		return err
	}

	return nil
}
