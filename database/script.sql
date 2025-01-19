DROP DATABASE IF EXISTS `golang_api`;
CREATE DATABASE `golang_api` CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `golang_api`;

-- ================================================================================================
-- ================================================================================================
-- ================================================================================================

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `last_name` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ix_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ================================================================================================

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ix_roles_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ================================================================================================

DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `role_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ix_user_roles_user_id` (`user_id`),
  KEY `ix_user_roles_role_id` (`role_id`),
  CONSTRAINT `fk_user_roles_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_user_roles_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ================================================================================================

DROP TABLE IF EXISTS `branches`;
CREATE TABLE `branches` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(255) NOT NULL,
  `description` VARCHAR(255) NOT NULL COMMENT "Nombre de la sucursal",
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ix_roles_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ================================================================================================

DROP TABLE IF EXISTS `user_branches`;
CREATE TABLE `user_branches` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `branch_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ix_user_branches_user_id` (`user_id`),
  KEY `ix_user_branches_branch_id` (`branch_id`),
  CONSTRAINT `fk_user_branches_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_user_branches_branch_id` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ================================================================================================

DROP TABLE IF EXISTS `activity_log_types`;
CREATE TABLE `activity_log_types` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `description` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ================================================================================================

DROP TABLE IF EXISTS `activity_logs`;
CREATE TABLE `activity_logs` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `activity_log_type_id` INT UNSIGNED NOT NULL,
  `user_id` BIGINT UNSIGNED NOT NULL,
  `branch_id` INT UNSIGNED NOT NULL,
  `description` VARCHAR(512) NOT NULL COMMENT "Texto que describe brevemente el evento",
  `metadata` JSON NOT NULL COMMENT "Información adicional del evento",
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `ix_activity_logs_activity_log_type_id` (`activity_log_type_id`),
  KEY `ix_activity_logs_user_id` (`user_id`),
  KEY `ix_activity_logs_branch_id` (`branch_id`),
  CONSTRAINT `fk_activity_logs_activity_log_type_id` FOREIGN KEY (`activity_log_type_id`) REFERENCES `activity_log_types` (`id`),
  CONSTRAINT `fk_activity_logs_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_activity_logs_branch_id` FOREIGN KEY (`branch_id`) REFERENCES `branches` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ================================================================================================
-- ================================================================================================
-- ================================================================================================

-- id
-- code
-- description
-- created_at
-- updated_at

INSERT INTO `roles` VALUES
  (1, "admin", "Administrador", NOW(), NULL),
  (2, "operator", "Operario", NOW(), NULL);

-- ================================================================================================

-- id
-- code
-- description
-- created_at
-- updated_at

INSERT INTO `branches` VALUES
  (1, "argentina", "Argentina", NOW(), NULL),
  (2, "peru", "Peru", NOW(), NULL);

-- ================================================================================================

-- id
-- description

INSERT INTO `activity_log_types` VALUES
  (1, "login"),
  (2, "forgot_password"),
  (3, "create_user"),
  (4, "update_user"),
  (5, "delete_user"),
  (6, "change_user_role"),
  (7, "reset_user_password"),
  (8, "activate_chip"),
  (9, "activate_batch_chips");
