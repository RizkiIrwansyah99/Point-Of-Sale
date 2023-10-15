CREATE TABLE users (
                       id VARCHAR(255) PRIMARY KEY,
                       firstname VARCHAR(255) NOT NULL,
                       lastname VARCHAR(255),
                       username VARCHAR(255) NOT NULL UNIQUE,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       role varchar(255) NOT NULL DEFAULT 'user',
                       image VARCHAR(255),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


create table token(
    id varchar(255) primary key,
    user_id varchar(255),
    token varchar(255),
    refresh_token varchar(255)
);