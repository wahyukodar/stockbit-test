CREATE
DATABASE IF NOT EXISTS `omdb` COLLATE 'utf8_general_ci' ;
GRANT ALL
ON `omdb`.* TO 'default'@'%' ;

CREATE TABLE `omdb`.`logs_movie`
(
    `id`          BIGINT      NOT NULL AUTO_INCREMENT,
    `log_request` TEXT        NOT NULL,
    `action`      VARCHAR(45) NOT NULL,
    `created_at`  TIMESTAMP   NOT NULL,
    PRIMARY KEY (`id`)
);

FLUSH PRIVILEGES ;