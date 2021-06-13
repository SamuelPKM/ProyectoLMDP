CREATE DATABASE IF NOT EXISTS tienda;
USE tienda;

-- -----------------------------------------------------
-- Table `tienda`.`clientes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tienda`.`clientes` (
  `ID_Cliente` INT NOT NULL AUTO_INCREMENT,
  `Nombre` VARCHAR(45) NULL,
  `Apellido` VARCHAR(45) NULL,
  `Correo` VARCHAR(45) NULL,
  `Telefono` VARCHAR(20) NULL,
  PRIMARY KEY (`ID_Cliente`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tienda`.`productos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tienda`.`productos` (
  `ID_Producto` INT NOT NULL AUTO_INCREMENT,
  `Nombre` VARCHAR(45) NULL,
  `Marca` VARCHAR(45) NULL,
  `Precio` INT NULL,
  PRIMARY KEY (`ID_Producto`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tienda`.`carrito`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tienda`.`carrito` (
  `ID_Carrito` INT NOT NULL AUTO_INCREMENT,
  `ID_Cliente` INT NOT NULL,
  `ID_Producto` INT NOT NULL,
  PRIMARY KEY (`ID_Carrito`),
  INDEX `fk_carrito_clientes_idx` (`ID_Cliente` ASC),
  INDEX `fk_carrito_productos1_idx` (`ID_Producto` ASC),
  CONSTRAINT `fk_carrito_clientes`
    FOREIGN KEY (`ID_Cliente`)
    REFERENCES `mydb`.`clientes` (`ID_Cliente`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_carrito_productos1`
    FOREIGN KEY (`ID_Producto`)
    REFERENCES `mydb`.`productos` (`ID_Producto`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `tienda`.`contactos`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `tienda`.`contactos` (
  `ID_Contacto` INT NOT NULL AUTO_INCREMENT,
  `Nombre` VARCHAR(45) NULL,
  `Apellido` VARCHAR(45) NULL,
  `Correo` VARCHAR(45) NULL,
  `Telefono` VARCHAR(20) NULL,
  `Descripcion` LONGTEXT NULL,
  PRIMARY KEY (`ID_Contacto`))
ENGINE = InnoDB;


INSERT INTO `productos` (`Nombre`,`Marca`,`Precio`)
				 VALUES ('Auriculares Gamer','ASTRO Gaming',3200),
						('Videojuego Gears 5','Microsoft',1200),
						('Kit de Carga','Microsoft',700),
						('Mando Xbox One','Microsoft',999),
						('Videojuego Mario Kart 8','Nintendo',999),
						('Tarjeta de regalo Xbox','Microsoft',500),
						('The Last of Us','Sony',1200),
						('Visores VR','Sony',999),
						('Xbox One X','Microsoft',5000);
                        
SELECT * FROM clientes;
SELECT * FROM contactos;