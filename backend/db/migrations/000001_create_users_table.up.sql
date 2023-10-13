create table users(
                      id varchar(255) primary key,
                      FirstName varchar(255) not null,
                      LastName varchar(255),
                      UserName varchar(255) not null unique,
                      Email varchar(255) not null unique unique ,
                      Password varchar(255) not null,
                      Role ENUM('superadmin, admin, user') default 'user',
                      Images varchar(255),
                      CreatedAt timestamp default current_timestamp,
                      UpdatedAt timestamp default current_timestamp on update current_timestamp

);