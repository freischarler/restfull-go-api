CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    first_name VARCHAR(150) NOT NULL,
    last_name VARCHAR(150) NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS posts (
    id serial NOT NULL,
    recipe_name VARCHAR(50) NOT NULL,
    recipe_type VARCHAR(50) NOT NULL,
    user_id int NOT NULL,
    ingredients text NOT NULL,
    description text NOT NULL,
    thumbnail text,
    likes int DEFAULT 0,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_notes PRIMARY KEY(id),
    CONSTRAINT fk_posts_users FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS comments (
    id serial NOT NULL,
    recipe_id int NOT NULL,
    name VARCHAR(50) NOT NULL,
    comment text,
    created_at timestamp DEFAULT now(),
    CONSTRAINT pk_comentary PRIMARY KEY(id),
    CONSTRAINT fk_comments_posts FOREIGN KEY(recipe_id) REFERENCES posts(id)
);
