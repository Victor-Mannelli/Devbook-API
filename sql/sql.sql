CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(50) not null,
  username varchar(50) not null unique,
  email varchar(50) not null unique,
  password varchar(100) not null,
  created_at timestamp default current_timestamp()
) ENGINE = INNODB;

DROP TABLE IF EXISTS followers;

CREATE TABLE followers(
  user_id int not null, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  follower_id int not null, FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  created_at timestamp default current_timestamp(),

  primary key(user_id, follower_id)
) ENGINE = INNODB;
