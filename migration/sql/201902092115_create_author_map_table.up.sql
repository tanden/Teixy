CREATE TABLE IF NOT EXISTS `author_map` (
  `book_id`    INT UNSIGNED NOT NULL,
  `author_id`  INT UNSIGNED NOT NULL,
  `mtime`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `ctime`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`book_id`, `author_id`),
  KEY `author_id` (`author_id`),
  FOREIGN KEY (`book_id`) REFERENCES books (`book_id`),
  FOREIGN KEY (`author_id`) REFERENCES authors (`author_id`)
) ENGINE=INNODB;
