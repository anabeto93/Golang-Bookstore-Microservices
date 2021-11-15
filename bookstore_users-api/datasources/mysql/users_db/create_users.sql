create TABLE IF NOT EXISTS `users` (
    id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    email varchar(100) NOT NULL UNIQUE,
    `status` VARCHAR(45) NOT NULL,
    `password` VARCHAR(120) NOT NULL,
    date_created DATETIME NOT NULL,
    date_updated DATETIME DEFAULT now(),
    UNIQUE INDEX email_UNIQUE (`email` ASC)
);
