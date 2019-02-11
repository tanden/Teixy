CREATE TABLE IF NOT EXISTS `articles` (
  `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `max_score`  INT UNSIGNED NOT NULL DEFAULT 990,
  `min_score`  INT UNSIGNED NOT NULL DEFAULT 0,
  `title`      TEXT NOT NULL,
  `punch_line` TEXT NOT NULL,
  `content`    TEXT NOT NULL,
  `status`     TINYINT NOT NULL,
  `mtime`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `ctime`      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `max_score` (`max_score`),
  KEY `min_score` (`min_score`)
) ENGINE=INNODB;
