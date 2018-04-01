USE gearbarter;

CREATE TABLE user (
    id SERIAL PRIMARY KEY,
    username VARCAHR(55) NOT NULL,
    password VARCHAR(55) NOT NULL, -- this will be salted

)

CREATE TABLE item (
    id SERIAL PRIMARY KEY,
    name VARCHAR(55) NOT NULL,
    description BLOB,
    user_id INT NOT NULL,
    image_id INT,

)
