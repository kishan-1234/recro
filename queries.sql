CREATE TABLE `user` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `email` varchar(30) UNIQUE,
  `phoneNumber` varchar(10) DEFAULT NULL,
  `meta` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `phoneNumber_idx` (`phoneNumber`)
);