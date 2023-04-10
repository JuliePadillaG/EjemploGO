  -- MySQL Workbench Forward Engineering

  -- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
  -- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
  -- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

  -- -----------------------------------------------------
  -- Schema mydb
  -- -----------------------------------------------------
  -- -----------------------------------------------------
  -- Schema bgow6s464
  -- -----------------------------------------------------

  -- -----------------------------------------------------
  -- Schema bgow6s464
  -- -----------------------------------------------------
  CREATE SCHEMA IF NOT EXISTS `bgow6s464` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
  USE `bgow6s464` ;

  -- -----------------------------------------------------
  -- Table `bgow6s464`.`buyers`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`buyers` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `card_number_id` TEXT NOT NULL,
    `first_name` TEXT NOT NULL,
    `last_name` TEXT NOT NULL,
    PRIMARY KEY (`id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`warehouses`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`warehouses` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `address` TEXT NULL DEFAULT NULL,
    `telephone` TEXT NULL DEFAULT NULL,
    `warehouse_code` TEXT NULL DEFAULT NULL,
    `minimum_capacity` INT NULL DEFAULT NULL,
    `minimum_temperature` INT NULL DEFAULT NULL,
    PRIMARY KEY (`id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`employees`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`employees` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `card_number_id` TEXT NOT NULL,
    `first_name` TEXT NOT NULL,
    `last_name` TEXT NOT NULL,
    `warehouse_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_employees_warehouses1_idx` (`warehouse_id` ASC) VISIBLE,
    CONSTRAINT `fk_employees_warehouses1`
      FOREIGN KEY (`warehouse_id`)
      REFERENCES `bgow6s464`.`warehouses` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`locality`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`locality` (
    `id` INT NOT NULL,
    `locality_name` VARCHAR(45) NULL,
    `province_name` VARCHAR(45) NULL,
    `country_name` VARCHAR(45) NULL,
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    PRIMARY KEY (`id`))
  ENGINE = InnoDB;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`seller`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`seller` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `cid` INT NOT NULL,
    `company_name` TEXT NOT NULL,
    `address` TEXT NOT NULL,
    `telephone` VARCHAR(15) NOT NULL,
    `locality_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_sellers_locality_idx` (`locality_id` ASC) VISIBLE,
    CONSTRAINT `fk_sellers_locality`
      FOREIGN KEY (`locality_id`)
      REFERENCES `bgow6s464`.`locality` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`products`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`products` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `description` TEXT NOT NULL,
    `expiration_rate` FLOAT NOT NULL,
    `freezing_rate` FLOAT NOT NULL,
    `height` FLOAT NOT NULL,
    `length` FLOAT NOT NULL,
    `netweight` FLOAT NOT NULL,
    `product_code` TEXT NOT NULL,
    `recommended_freezing_temperature` FLOAT NOT NULL,
    `width` FLOAT NOT NULL,
    `product_type_id` INT NOT NULL,
    `seller_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_products_sellers1_idx` (`seller_id` ASC) VISIBLE,
    CONSTRAINT `fk_products_seller1`
      FOREIGN KEY (`seller_id`)
      REFERENCES `bgow6s464`.`seller` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB
  AUTO_INCREMENT = 6
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`sections`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`sections` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `section_number` INT NOT NULL,
    `current_temperature` INT NOT NULL,
    `minimum_temperature` INT NOT NULL,
    `current_capacity` INT NOT NULL,
    `minimum_capacity` INT NOT NULL,
    `maximum_capacity` INT NOT NULL,
    `warehouse_id` INT NOT NULL,
    `id_product_type` INT NOT NULL,
    PRIMARY KEY (`id`))
  ENGINE = InnoDB
  DEFAULT CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`carries`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`carries` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `cid` VARCHAR(45) NULL,
    `company_name` VARCHAR(45) NULL,
    `address` VARCHAR(45) NULL,
    `telephone` VARCHAR(45) NULL,
    `locality_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_carries_locality1_idx` (`locality_id` ASC) VISIBLE,
    CONSTRAINT `fk_carries_locality1`
      FOREIGN KEY (`locality_id`)
      REFERENCES `bgow6s464`.`locality` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`product_batches`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`product_batches` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `batch_number` VARCHAR(45) NULL,
    `current_quantity` INT NULL,
    `current_temperature` INT NULL,
    `due_date` DATETIME NULL,
    `initial_quantity` INT NULL,
    `manufacturing_date` DATE NULL,
    `manufacturing_hour` VARCHAR(45) NULL,
    `minimum_temperature` INT NULL,
    `sections_id` INT NOT NULL,
    `products_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_product_batches_sections1_idx` (`sections_id` ASC) VISIBLE,
    INDEX `fk_product_batches_products1_idx` (`products_id` ASC) VISIBLE,
    CONSTRAINT `fk_product_batches_sections1`
      FOREIGN KEY (`sections_id`)
      REFERENCES `bgow6s464`.`sections` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
    CONSTRAINT `fk_product_batches_products1`
      FOREIGN KEY (`products_id`)
      REFERENCES `bgow6s464`.`products` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`product_records`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`product_records` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `last_update_date` DATETIME NULL,
    `purchase_price` FLOAT NULL,
    `sale_price` FLOAT NULL,
    `products_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_product_records_products1_idx` (`products_id` ASC) VISIBLE,
    CONSTRAINT `fk_product_records_products1`
      FOREIGN KEY (`products_id`)
      REFERENCES `bgow6s464`.`products` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`inbound_orders`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`inbound_orders` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `order_date` DATETIME NULL,
    `order_number` VARCHAR(45) NULL,
    `employee_id` INT NOT NULL,
    `warehouse_id` INT NOT NULL,
    `product_batch_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_inbound_orders_employees1_idx` (`employee_id` ASC) VISIBLE,
    INDEX `fk_inbound_orders_warehouses1_idx` (`warehouse_id` ASC) VISIBLE,
    INDEX `fk_inbound_orders_product_batches1_idx` (`product_batch_id` ASC) VISIBLE,
    CONSTRAINT `fk_inbound_orders_employees1`
      FOREIGN KEY (`employee_id`)
      REFERENCES `bgow6s464`.`employees` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
    CONSTRAINT `fk_inbound_orders_warehouses1`
      FOREIGN KEY (`warehouse_id`)
      REFERENCES `bgow6s464`.`warehouses` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
    CONSTRAINT `fk_inbound_orders_product_batches1`
      FOREIGN KEY (`product_batch_id`)
      REFERENCES `bgow6s464`.`product_batches` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB;


  -- -----------------------------------------------------
  -- Table `bgow6s464`.`purchase_orders`
  -- -----------------------------------------------------
  CREATE TABLE IF NOT EXISTS `bgow6s464`.`purchase_orders` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `order_number` VARCHAR(45) NULL,
    `order_date` DATETIME NULL,
    `tracking_code` VARCHAR(45) NULL,
    `buyers_id` INT NOT NULL,
    `product_records_id` INT NOT NULL,
    `order_status_id` INT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_buyers1_idx` (`buyers_id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_product_records1_idx` (`product_records_id` ASC) VISIBLE,
    CONSTRAINT `fk_purchase_orders_buyers1`
      FOREIGN KEY (`buyers_id`)
      REFERENCES `bgow6s464`.`buyers` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE,
    CONSTRAINT `fk_purchase_orders_product_records1`
      FOREIGN KEY (`product_records_id`)
      REFERENCES `bgow6s464`.`product_records` (`id`)
      ON DELETE CASCADE
      ON UPDATE CASCADE)
  ENGINE = InnoDB;


  -- SET SQL_MODE=@OLD_SQL_MODE;
  -- SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
  -- SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;