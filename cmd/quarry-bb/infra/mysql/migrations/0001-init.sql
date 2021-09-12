CREATE DATABASE quarry_bb;
USE quarry_bb;
CREATE TABLE user (
  id INT unsigned NOT NULL AUTO_INCREMENT, 
  full_name VARCHAR(255), 
  short_name VARCHAR(255), 
  uid VARCHAR(255) NOT NULL, 
  password_hash VARCHAR(255) NOT NULL, 
  login VARCHAR(255) NOT NULL, 
  verified BOOLEAN, 
  created TIMESTAMP, 
  PRIMARY KEY (id)
);
