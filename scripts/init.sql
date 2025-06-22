CREATE DATABASE IF NOT EXISTS todo_database; 



CREATE TABLE todo
(
  id           INT(10) NOT NULL AUTO_INCREMENT,
  title     VARCHAR(50) not null,
  description TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  update_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  is_activate BOOLEAN DEFAULT true,
  PRIMARY KEY(id)
);

INSERT INTO todo (title,description) VALUES ("title_1","description_1");
INSERT INTO todo (title,description) VALUES ("title_2","description_2");




-- MYSQL_ROOT_PASSWORD=root_password
-- MYSQL_DATABASE=todo_database
-- MYSQL_USER=mysql_user
-- MYSQL_PASSWORD=todo_database_password