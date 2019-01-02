DROP TABLE IF EXISTS `dblearning`.`UserSession`;
DROP TABLE IF EXISTS `dblearning`.`User`;


CREATE TABLE `dblearning`.`User` (
  `ID` INT NOT NULL,
  `Username` VARCHAR(45) NOT NULL,
  `Fullname` VARCHAR(80) NOT NULL,
  `Hash` VARCHAR(80) NOT NULL,
  `Salt` VARCHAR(80) NOT NULL,
  `disabled` TINYINT NOT NULL,
  PRIMARY KEY (`ID`));

  CREATE TABLE `dblearning`.`UserSession` (
  `SessionKey` varchar(256) NOT NULL,
  `UserID` int(11) NOT NULL,
  `LoginTIme` datetime NOT NULL,
  `LastSeenTime` datetime NOT NULL,
  PRIMARY KEY (`SessionKey`),
  KEY `fk_UserSession_1_idx` (`UserID`),
  CONSTRAINT `fk_UserSession_1` FOREIGN KEY (`UserID`) REFERENCES `User` (`ID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
