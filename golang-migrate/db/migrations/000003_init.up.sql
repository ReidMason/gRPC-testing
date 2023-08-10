CREATE TABLE books (
  id   BIGSERIAL PRIMARY KEY,
  authorId BIGSERIAL REFERENCES authors(id) NOT NULL,
  title text NOT NULL
);

INSERT INTO authors (id, bio)
VALUES (1, 'Some bio about the author');


INSERT INTO books (authorId, title)
VALUES (1, 'Some exciting book title');
