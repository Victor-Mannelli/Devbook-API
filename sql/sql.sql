CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
  id int auto_increment primary key,
  name varchar(50) not null,
  username varchar(50) not null unique,
  email varchar(50) not null unique,
  password varchar(100) not null,
  created_at timestamp default current_timestamp()
) ENGINE = INNODB;


CREATE TABLE followers(
  user_id int not null, 
  follower_id int not null, 
  created_at timestamp default current_timestamp(),

  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  primary key(user_id, follower_id)
) ENGINE = INNODB;


CREATE TABLE posts(
  post_id int auto_increment primary key,
  title varchar(50) not null,
  content varchar(300) not null,
  author_id int not null, 
  likes int default 0,
  created_at timestamp default current_timestamp(),
  
  FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = INNODB;
