-- MySQL dump 10.16  Distrib 10.1.38-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: mkultra    Database: retrosheet_events
-- ------------------------------------------------------
-- Server version	10.3.12-MariaDB-1:10.3.12+maria~bionic

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `__schema_version`
--

DROP TABLE IF EXISTS `__schema_version`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `__schema_version` (
  `version` varchar(8) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `coms`
--

DROP TABLE IF EXISTS `coms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `coms` (
  `game_id` char(12) NOT NULL,
  `idx` int(11) NOT NULL,
  `description` varchar(128) NOT NULL,
  PRIMARY KEY (`game_id`,`idx`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `data`
--

DROP TABLE IF EXISTS `data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `data` (
  `game_id` char(12) NOT NULL,
  `player_id` char(8) NOT NULL,
  `er` int(11) NOT NULL,
  PRIMARY KEY (`game_id`,`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `games`
--

DROP TABLE IF EXISTS `games`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `games` (
  `game_id` char(12) NOT NULL,
  `season` int(11) NOT NULL DEFAULT 0,
  `visteam` char(3) NOT NULL,
  `hometeam` char(3) NOT NULL,
  `game_date` date NOT NULL,
  `number` int(11) NOT NULL DEFAULT 0,
  `starttime` varchar(16) NOT NULL,
  `daynight` varchar(16) NOT NULL,
  `usedh` tinyint(1) NOT NULL DEFAULT 0,
  `pitches` varchar(7) NOT NULL,
  `umphome` char(8) DEFAULT NULL,
  `ump1b` char(8) DEFAULT NULL,
  `ump2b` char(8) DEFAULT NULL,
  `ump3b` char(8) DEFAULT NULL,
  `umplf` char(8) DEFAULT NULL,
  `umprf` char(8) DEFAULT NULL,
  `fieldcond` varchar(16) NOT NULL,
  `precip` varchar(16) NOT NULL,
  `sky` varchar(16) NOT NULL,
  `temp` int(11) NOT NULL DEFAULT 0,
  `winddir` varchar(16) NOT NULL,
  `windspeed` int(11) NOT NULL DEFAULT 0,
  `timeofgame` int(11) NOT NULL DEFAULT 0,
  `attendance` int(11) NOT NULL DEFAULT 0,
  `site` char(5) NOT NULL,
  `wp` char(8) NOT NULL,
  `lp` char(8) NOT NULL,
  `save` char(8) DEFAULT NULL,
  `gwrbi` char(8) DEFAULT NULL,
  `edittime` varchar(32) DEFAULT NULL,
  `howscored` varchar(32) DEFAULT NULL,
  `inputprogvers` varchar(32) DEFAULT NULL,
  `inputter` varchar(32) DEFAULT NULL,
  `inputtime` varchar(32) DEFAULT NULL,
  `scorer` varchar(64) DEFAULT NULL,
  `translator` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`game_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `plays`
--

DROP TABLE IF EXISTS `plays`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `plays` (
  `game_id` char(12) NOT NULL,
  `idx` int(11) NOT NULL,
  `inning` int(11) NOT NULL,
  `team` int(11) NOT NULL,
  `player_id` char(8) NOT NULL,
  `count` varchar(16) NOT NULL,
  `pitches` varchar(64) NOT NULL,
  `event` varchar(64) NOT NULL,
  PRIMARY KEY (`game_id`,`idx`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `starters`
--

DROP TABLE IF EXISTS `starters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `starters` (
  `game_id` char(12) NOT NULL,
  `player_id` char(8) NOT NULL,
  `name` varchar(64) NOT NULL,
  `team` int(11) NOT NULL,
  `batting_order` int(11) NOT NULL,
  `position` int(11) NOT NULL,
  PRIMARY KEY (`game_id`,`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `subs`
--

DROP TABLE IF EXISTS `subs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `subs` (
  `game_id` char(12) NOT NULL,
  `player_id` char(8) NOT NULL,
  `idx` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `team` int(11) NOT NULL,
  `batting_order` int(11) NOT NULL,
  `position` int(11) NOT NULL,
  PRIMARY KEY (`game_id`,`player_id`,`idx`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-03-14 17:40:54
