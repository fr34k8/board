
CREATE TABLE IF NOT EXISTS `card` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`project` INT UNSIGNED NOT NULL,
	`number` INT UNSIGNED NOT NULL,
	`title` TEXT,
	`description` TEXT,
	`createdAt` TIMESTAMP NOT NULL DEFAULT NOW(),
	`updatedAt` TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE (`project`, `number`)
);

CREATE TABLE IF NOT EXISTS `board` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`title` TEXT,
	`createdAt` TIMESTAMP NOT NULL DEFAULT NOW(),
	`updatedAt` TIMESTAMP,
	PRIMARY KEY (`id`)
);

-- One card can belong to many boards
CREATE TABLE IF NOT EXISTS `card_board` (
	`board` INT UNSIGNED NOT NULL,
	`card` INT UNSIGNED NOT NULL,
	PRIMARY KEY (`board`, `card`)
);

CREATE TABLE IF NOT EXISTS `project` (
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`key` VARCHAR(10) NOT NULL,
	`counter` INT UNSIGNED NOT NULL DEFAULT 0,
	`title` TEXT,
	PRIMARY KEY (`id`),
	UNIQUE (`key`)
);
