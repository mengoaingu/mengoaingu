CREATE TABLE `tasks`
(
    `id`                        bigint(20)          NOT NULL AUTO_INCREMENT,
    `name`                      text                NOT NULL,
    `description`               text,
    `priority`                  varchar(50),
    `status`                    varchar(50),
    `user_id`                   varchar(50)         NOT NULL,
    `created_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`                TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;