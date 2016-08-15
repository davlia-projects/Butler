CREATE TABLE IF NOT EXISTS `questions` (
  `qid` INTEGER PRIMARY KEY AUTOINCREMENT,
  `prompt` VARCHAR(2048) NULL,
  `answer` VARCHAR(2048) NULL,
  `category` VARCHAR(2048) NULL,
  `tags` VARCHAR(2048) NULL,
  `created` DATE NULL
);
