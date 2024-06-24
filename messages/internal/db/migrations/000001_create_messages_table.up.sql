CREATE TABLE messages IF NOT EXISTS (
    chatid text NOT NULL UNIQUE,
    senderid text NOT NULL UNIQUE,
    receiverid text NOT NULL UNIQUE,
    messageid text NOT NULL UNIQUE,
    val text NOT NULL,
    sentat TIMESTAMP DEFAULT CURRENT_TIMESTAMP                              
);