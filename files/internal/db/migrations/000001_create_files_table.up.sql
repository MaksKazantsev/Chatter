CREATE TABLE IF NOT EXISTS files
(
    filelink text NOT NULL UNIQUE,
    fileid text NOT NULL UNIQUE,
    userid text NOT NULL,
    uploadedat TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);