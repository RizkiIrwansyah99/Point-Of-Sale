create table users(
                      id varchar(255) primary key,
                      firstname varchar(255) not null,
                      lastname varchar(255),
                      username varchar(255) not null unique,
                      email varchar(255) not null unique unique ,
                      password varchar(255) not null,
                      role ENUM('superadmin, admin, user') default 'user',
                      image varchar(255),
                      created_at timestamp default current_timestamp,
                      updated_at timestamp default current_timestamp on update current_timestamp

);