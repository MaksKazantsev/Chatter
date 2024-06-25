CREATE TABLE IF NOT EXISTS messages
(
    chatid text NOT NULL,
    senderid text NOT NULL,
    receiverid text NOT NULL,
    messageid text NOT NULL UNIQUE,
    val text NOT NULL,
    sentat TIMESTAMP
    );
CREATE TABLE IF NOT EXISTS chats
(
  chatid text NOT NULL DEFAULT '',
  chatphoto text NOT NULL DEFAULT '',
  chatname text NOT NULL DEFAULT '',
  missed smallint NOT NULL DEFAULT 0,
  userid text NOT NULL DEFAULT ''
  );
CREATE TABLE IF NOT EXISTS chat_members
(
    userid text NOT NULL DEFAULT '',
    chatid text NOT NULL DEFAULT '',
    missed smallint NOT NULL DEFAULT 0
);