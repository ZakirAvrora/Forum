CREATE TABLE Users
(
    username VARCHAR(150) PRIMARY KEY UNIQUE,
    email VARCHAR(150) UNIQUE NOT NULL,
    passwd VARCHAR(150) NOT NULL
);

CREATE TABLE Posts
(
    post_id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_title TEXT DEFAULT 'Untitled',
    post_user VARCHAR(150) NOT NULL,
    post_content TEXT NOT NULL,
    post_likes INTEGER DEFAULT 0,
    post_dislikes INTEGER DEFAULT 0,
    post_createdAt TEXT NOT NULL,
    FOREIGN KEY (post_user) REFERENCES Users(username)
);

CREATE TABLE Comments
(
    comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
    comment_user VARCHAR(150) NOT NULL,
    comment_post INTEGER NOT NULL,
    comment_text TEXT NOT NULL ,
    comment_createdAt TEXT NOT NULL,
    FOREIGN KEY (comment_user) REFERENCES Users(username),
    FOREIGN KEY (comment_post) REFERENCES Posts(post_id)
);

CREATE TABLE Categories
(
    category_id INTEGER PRIMARY KEY AUTOINCREMENT,
    category_name varchar(150) NOT NULL
);

CREATE TABLE Post_Category
(
    pc_id INTEGER PRIMARY KEY AUTOINCREMENT,
    pc_post_id INTEGER,
    pc_category_id INTEGER,
    FOREIGN KEY (pc_post_id) REFERENCES Posts(post_id),
    FOREIGN KEY (pc_category_id) REFERENCES Categories(category_id)
);

CREATE TABLE Post_Comments
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    post_id INTEGER ,
    comment_id INTEGER,
    FOREIGN KEY (post_id) REFERENCES Posts(post_id),
    FOREIGN KEY (comment_id) REFERENCES Comments(comment_id)
);