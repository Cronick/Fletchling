CREATE TABLE IF NOT EXISTS `nests` (
  `nest_id` bigint(20) NOT NULL,
  `lat` float NOT NULL,
  `lon` float NOT NULL,
  `name` varchar(250) NOT NULL DEFAULT 'unknown',
  `polygon` geometry NOT NULL,
  `area_name` varchar(250) DEFAULT NULL,
  `spawnpoints` smallint(5) unsigned DEFAULT 0,
  `m2` decimal(10,1) DEFAULT 0.0,
  `active` tinyint(1) DEFAULT 0,
  `pokemon_id` int(11) DEFAULT NULL,
  `pokemon_form` smallint(6) DEFAULT NULL,
  `pokemon_avg` float DEFAULT NULL,
  `pokemon_ratio` float DEFAULT 0,
  `pokemon_count` float DEFAULT 0,
  `discarded` varchar(40) DEFAULT NULL,
  `updated` int(10) DEFAULT NULL,
  PRIMARY KEY (`nest_id`),
  KEY `ix_coords` (`lat`,`lon`),
  KEY `ix_nests_updated` (`updated`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ALTER TABLE nests MODIFY nest_id BIGINT NOT NULL AUTO_INCREMENT;
