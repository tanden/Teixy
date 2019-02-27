CREATE TABLE IF NOT EXISTS `books` (
  `book_id`    INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `isbn`       BIGINT UNSIGNED NOT NUll,
  `min_score`  INT UNSIGNED NOT NULL DEFAULT 0,
  `max_score`  INT UNSIGNED NOT NULL DEFAULT 990,
  `title`      TEXT NOT NULL,
  `punch_line` TEXT NOT NULL,
  `article`    TEXT NOT NULL,
  `status`     TINYINT NOT NULL,
  `mtime`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `ctime`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`book_id`),
  KEY `max_score` (`max_score`),
  KEY `min_score` (`min_score`)
) ENGINE=INNODB;
