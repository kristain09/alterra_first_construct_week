-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `mydb` ;

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `mydb` DEFAULT CHARACTER SET utf8mb3 ;
USE `mydb` ;

-- -----------------------------------------------------
-- Table `mydb`.`users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`users` ;

CREATE TABLE IF NOT EXISTS `mydb`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(100) NULL DEFAULT NULL,
  `password` VARCHAR(45) NULL DEFAULT NULL,
  `deleted_at` TIMESTAMP(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mydb`.`customers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`customers` ;

CREATE TABLE IF NOT EXISTS `mydb`.`customers` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL DEFAULT NULL,
  `created_by` INT NULL DEFAULT NULL,
  `users_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_customers_users1_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_customers_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `mydb`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mydb`.`products`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`products` ;

CREATE TABLE IF NOT EXISTS `mydb`.`products` (
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
    REFERENCES `mydb`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mydb`.`transactions`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`transactions` ;

CREATE TABLE IF NOT EXISTS `mydb`.`transactions` (
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
    REFERENCES `mydb`.`customers` (`id`),
  CONSTRAINT `fk_transactions_users1`
    FOREIGN KEY (`created_by`)
    REFERENCES `mydb`.`users` (`id`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `mydb`.`products_has_transactions`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `mydb`.`products_has_transactions` ;

CREATE TABLE IF NOT EXISTS `mydb`.`products_has_transactions` (
  `products_id` INT NOT NULL,
  `transactions_id` INT NOT NULL,
  PRIMARY KEY (`products_id`, `transactions_id`),
  INDEX `fk_products_has_transactions_transactions1_idx` (`transactions_id` ASC) VISIBLE,
  INDEX `fk_products_has_transactions_products1_idx` (`products_id` ASC) VISIBLE,
  CONSTRAINT `fk_products_has_transactions_products1`
    FOREIGN KEY (`products_id`)
    REFERENCES `mydb`.`products` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_products_has_transactions_transactions1`
    FOREIGN KEY (`transactions_id`)
    REFERENCES `mydb`.`transactions` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb3;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
