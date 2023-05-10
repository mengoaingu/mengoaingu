CREATE TABLE `quizzes`
(
    `id`                        bigint(20)          NOT NULL AUTO_INCREMENT,
    `topic_id`                  bigint(20),
    `name`                      text,
    `description`               text,
    `is_free`                   boolean,
    `created_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`                TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (1, 'Quiz 1', 'Description 1', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (1, 'Quiz 2', 'Description 2', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (1, 'Quiz 3', 'Description 3', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (2, 'Quiz 4', 'Description 4', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (2, 'Quiz 5', 'Description 5', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (2, 'Quiz 6', 'Description 6', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (3, 'Quiz 7', 'Description 7', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (3, 'Quiz 8', 'Description 8', true);
INSERT INTO `quizzes`(`topic_id`, `name`,`description`, `is_free`)  VALUES (3, 'Quiz 9', 'Description 9', true);