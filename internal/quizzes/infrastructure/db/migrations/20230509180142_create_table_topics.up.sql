CREATE TABLE `topics`
(
    `id`                        bigint(20)          NOT NULL AUTO_INCREMENT,
    `title`                     text,
    `thumbnail`                 text,
    `created_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`                TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`                TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO `topics`(`title`,`thumbnail`) VALUES ('Topic 1', 'https://images.unsplash.com');
INSERT INTO `topics`(`title`,`thumbnail`) VALUES ('Topic 2', 'https://images.unsplash.com');
INSERT INTO `topics`(`title`,`thumbnail`) VALUES ('Topic 3', 'https://images.unsplash.com');