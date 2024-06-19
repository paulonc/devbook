CREATE TABLE publications (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(300) NOT NULL,
    author_id INT NOT NULL,
    likes INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_author
    FOREIGN KEY(author_id) 
    REFERENCES users(id) 
    ON DELETE CASCADE
);