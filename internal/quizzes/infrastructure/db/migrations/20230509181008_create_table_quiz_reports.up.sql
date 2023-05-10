CREATE TABLE `quiz_reports`
(
    `id`                        bigint(20)          NOT NULL AUTO_INCREMENT,
    `user_id`                   bigint(20),
    `topic_id`                  bigint(20),
    `quiz_id`                   bigint(20),
    `percentage`                varchar(30),
    `duration`                  int,
    `total_duration`            int,
    `result`                    varchar(20),
    `start_time`                bigint(20),
    `total_quiz`                int,
    `mask_obtained`             int,
    `state`                     varchar(20),
    `answers`                   json,
    `created_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`                TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;