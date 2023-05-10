CREATE TABLE `users`
(
    `id`                        bigint(20)          NOT NULL AUTO_INCREMENT,
    `uid`                       varchar(50)         NOT NULL,
    `display_name`              varchar(50),
    `email`                     varchar(50),
    `avatar`                    text,
    `provider`                  varchar(50),
    `created_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`                TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;