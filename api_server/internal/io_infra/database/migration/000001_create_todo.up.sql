

CREATE TABLE todo
(
  id           INT(10) NOT NULL AUTO_INCREMENT,
  title     VARCHAR(50) NOT NULL,
  description TEXT NOT NULL,
  limit_time DATETIME NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  update_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  status_no INT(2) NOT NULL,
  is_activate BOOLEAN NOT NULL DEFAULT true,
  PRIMARY KEY(id)
);

-- INSERT INTO todo (title,description) VALUES ("title_1","description_1");
-- INSERT INTO todo (title,description) VALUES ("title_2","description_2");
