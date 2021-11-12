create TABLE IF NOT EXISTS `users` (
    id BIGINT(20) AUTO_INCREMENT PRIMARY KEY,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    date_created DATETIME NOT NULL,
    date_updated DATETIME DEFAULT now()
);
