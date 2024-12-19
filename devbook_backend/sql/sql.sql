CREATE DATABASE IF NOT EXISTS DEVBOOK;
USE DEVBOOK;

DROP TABLE IF EXISTS LIKES;
DROP TABLE IF EXISTS POSTS;
DROP TABLE IF EXISTS FOLLOWERS;
DROP TABLE IF EXISTS USERS;

CREATE TABLE USERS(
  ID CHAR(36) PRIMARY KEY DEFAULT (UUID()),
  NAME VARCHAR(255) NOT NULL,
  NICK VARCHAR(255) NOT NULL UNIQUE,
  EMAIL VARCHAR(255) NOT NULL UNIQUE,
  `PASSWORD` VARCHAR(255) NOT NULL,
  CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB; 

CREATE TABLE FOLLOWERS(
  FOLLOWING_ID CHAR(36) NOT NULL,
  FOREIGN KEY (FOLLOWING_ID) REFERENCES USERS(ID) ON DELETE CASCADE,

  FOLLOWED_ID CHAR(36) NOT NULL,
  FOREIGN KEY (FOLLOWED_ID) REFERENCES USERS(ID) ON DELETE CASCADE,

  PRIMARY KEY(FOLLOWING_ID, FOLLOWED_ID)
) ENGINE=INNODB;

CREATE TABLE POSTS(
  ID CHAR(36) PRIMARY KEY DEFAULT (UUID()),
  TITLE VARCHAR(255) NOT NULL,
  CONTENT VARCHAR(255) NOT NULL,
  AUTHOR_ID CHAR(36) NOT NULL,
  AUTHOR_NICK VARCHAR(255) NOT NULL,
  LIKES INT DEFAULT 0, 
  CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (AUTHOR_ID) REFERENCES USERS(ID) ON DELETE CASCADE,
  FOREIGN KEY (AUTHOR_NICK) REFERENCES USERS(NICK) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE LIKES(
	POST_ID CHAR(36) NOT NULL, 
  USER_ID CHAR(36) NOT NULL,
  
  FOREIGN KEY (POST_ID) REFERENCES POSTS(ID) ON DELETE CASCADE,
  FOREIGN KEY (USER_ID) REFERENCES USERS(ID) ON DELETE CASCADE,
  
  PRIMARY KEY(POST_ID, USER_ID)
) ENGINE=INNODB;








