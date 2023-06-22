-- create the databases
CREATE DATABASE IF NOT EXISTS golara;

-- create the users for each database
CREATE USER 'golarauser'@'%' IDENTIFIED BY '^^9o9!44mzI^';
GRANT CREATE, ALTER, INDEX, LOCK TABLES, REFERENCES, UPDATE, DELETE, DROP, SELECT, INSERT ON `golara`.* TO 'golarauser'@'%';

FLUSH PRIVILEGES;