CREATE DATABASE IF NOT EXISTS DEVBOOK;
USE DEVBOOK;

DROP TABLE IF EXISTS USERS;

CREATE TABLE USERS(
  ID CHAR(36) PRIMARY KEY DEFAULT (UUID()),
  `NAME` VARCHAR(64) NOT NULL,
  NICK VARCHAR(64) NOT NULL UNIQUE,
  EMAIL VARCHAR(64) NOT NULL UNIQUE,
  `PASSWORD` VARCHAR(64) NOT NULL,
  CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB; 