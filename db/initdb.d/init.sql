DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id BINARY(36) NOT NULL,
    name VARCHAR(32) NOT NULL,
    created_at datetime default current_timestamp,
    PRIMARY KEY (id)
) DEFAULT CHARACTER SET=utf8;

INSERT INTO users (id, name, created_at) VALUES (UUID(), 'bassaer', NOW());
