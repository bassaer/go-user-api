DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id INT AUTO_INCREMENT,
    name VARCHAR(32) NOT NULL,
    created_at datetime default current_timestamp,
    PRIMARY KEY (id)
) DEFAULT CHARACTER SET=utf8;

INSERT INTO users (name, created_at) VALUES ('bassaer', NOW());
