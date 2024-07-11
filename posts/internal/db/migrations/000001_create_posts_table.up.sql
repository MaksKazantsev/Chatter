CREATE TABLE IF NOT EXISTS posts
(
    userid text NOT NULL,
    postid text NOT NULL UNIQUE,
    posttitle text NOT NULL,
    postdesc text NOT NULL,
    postfile text NOT NULL,
    likesamount smallint NOT NULL DEFAULT 0,
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS comments
(
    userid text NOT NULL,
    postid text NOT NULL,
    commentid text NOT NULL UNIQUE,
    createdat TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    val text NOT NULL
);

CREATE TABLE IF NOT EXISTS likes
(
    userid text NOT NULL,
    postid text NOT NULL,
    likeid text NOT NULL UNIQUE
);