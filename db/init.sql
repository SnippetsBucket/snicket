CREATE TABLE snippets (
  snippet_id int(11) NOT NULL AUTO_INCREMENT,
  snippet_title varchar(255) NOT NULL,
  snippet_text text NOT NULL,
  user_id int(11) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  PRIMARY KEY (snippet_id)
);
