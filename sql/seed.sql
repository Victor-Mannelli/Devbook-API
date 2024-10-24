insert into users (name, username, email, password)
values 
("User 1", "user_1", "user1@gmail.com", "$2a$10$UJ1GXHF5TVlWCsRjGiubvOQpBXDevr5v7UGqU9qD2HIAmuRqtoe8e"),
("User 2", "user_2", "user2@gmail.com", "$2a$10$UJ1GXHF5TVlWCsRjGiubvOQpBXDevr5v7UGqU9qD2HIAmuRqtoe8e"),
("User 3", "user_3", "user3@gmail.com", "$2a$10$UJ1GXHF5TVlWCsRjGiubvOQpBXDevr5v7UGqU9qD2HIAmuRqtoe8e"),
("User 4", "user_4", "user4@gmail.com", "$2a$10$UJ1GXHF5TVlWCsRjGiubvOQpBXDevr5v7UGqU9qD2HIAmuRqtoe8e"),
("User 5", "user_5", "user5@gmail.com", "$2a$10$UJ1GXHF5TVlWCsRjGiubvOQpBXDevr5v7UGqU9qD2HIAmuRqtoe8e"),
("User 6", "user_6", "user6@gmail.com", "$2a$10$UJ1GXHF5TVlWCsRjGiubvOQpBXDevr5v7UGqU9qD2HIAmuRqtoe8e");

insert into followers (user_id, follower_id)
values 
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(2, 3),
(2, 4),
(2, 5),
(2, 6),
(3, 4),
(3, 5),
(3, 6),
(4, 5),
(4, 6),
(5, 6);

insert into posts (post_id, title, content, author_id, likes)
values 
(1, "title 1", "content_1", 1, 0),
(2, "title 2", "content_2", 2, 2),
(3, "title 3", "content_3", 3, 4),
(4, "title 4", "content_4", 4, 8);