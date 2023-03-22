-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema first_construct_week
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `first_construct_week` ;

-- -----------------------------------------------------
-- Schema first_construct_week
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `first_construct_week` DEFAULT CHARACTER SET utf8mb3 ;
USE `first_construct_week` ;

-- -----------------------------------------------------
-- Table `first_construct_week`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `first_construct_week`.`users` ;

CREATE TABLE IF NOT EXISTS `first_construct_week`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(100) NULL DEFAULT NULL,
  `password` VARCHAR(45) NULL DEFAULT NULL,
  `deleted_at` TIMESTAMP(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `first_construct_week`.`customers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `first_construct_week`.`customers` ;

CREATE TABLE IF NOT EXISTS `first_construct_week`.`customers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  `created_by` INT NULL DEFAULT NULL,
  `users_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_customers_users1_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_customers_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `first_construct_week`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `first_construct_week`.`products`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `first_construct_week`.`products` ;

CREATE TABLE IF NOT EXISTS `first_construct_week`.`products` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NULL DEFAULT NULL,
  `price` INT NULL DEFAULT NULL,
  `stock` INT NULL DEFAULT NULL,
  `deleted_at` TIMESTAMP(3) NULL DEFAULT NULL,
  `created_by` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_products_users1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_products_users1`
    FOREIGN KEY (`created_by`)
    REFERENCES `first_construct_week`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `first_construct_week`.`transactions`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `first_construct_week`.`transactions` ;

CREATE TABLE IF NOT EXISTS `first_construct_week`.`transactions` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `invoice` INT NOT NULL,
  `transdate` DATETIME(3) NOT NULL,
  `total` INT NULL DEFAULT NULL,
  `customers_id` INT NOT NULL,
  `created_by` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `invoice_UNIQUE` (`invoice` ASC) VISIBLE,
  INDEX `fk_products_has_customers_customers1_idx` (`customers_id` ASC) VISIBLE,
  INDEX `fk_transactions_users1_idx` (`created_by` ASC) VISIBLE,
  CONSTRAINT `fk_products_has_customers_customers1`
    FOREIGN KEY (`customers_id`)
    REFERENCES `first_construct_week`.`customers` (`id`),
  CONSTRAINT `fk_transactions_users1`
    FOREIGN KEY (`created_by`)
    REFERENCES `first_construct_week`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `first_construct_week`.`products_has_transactions`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `first_construct_week`.`products_has_transactions` ;

CREATE TABLE IF NOT EXISTS `first_construct_week`.`products_has_transactions` (
  `products_id` INT NOT NULL,
  `transactions_id` INT NOT NULL,
  PRIMARY KEY (`products_id`, `transactions_id`),
  INDEX `fk_products_has_transactions_transactions1_idx` (`transactions_id` ASC) VISIBLE,
  INDEX `fk_products_has_transactions_products1_idx` (`products_id` ASC) VISIBLE,
  CONSTRAINT `fk_products_has_transactions_products1`
    FOREIGN KEY (`products_id`)
    REFERENCES `first_construct_week`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_products_has_transactions_transactions1`
    FOREIGN KEY (`transactions_id`)
    REFERENCES `first_construct_week`.`transactions` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
