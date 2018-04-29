-- MySQL dump 10.13  Distrib 5.5.60, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: baseball_databank_2017
-- ------------------------------------------------------
-- Server version	5.5.60-0+deb8u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `allstarfull`
--

DROP TABLE IF EXISTS `allstarfull`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `allstarfull` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `gameNum` int(11) NOT NULL,
  `gameID` varchar(255) DEFAULT NULL,
  `teamID` varchar(255) DEFAULT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `GP` int(11) DEFAULT NULL,
  `startingPos` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`gameNum`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `appearances`
--

DROP TABLE IF EXISTS `appearances`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `appearances` (
  `yearID` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `playerID` varchar(255) NOT NULL,
  `G_all` int(11) NOT NULL,
  `GS` int(11) DEFAULT NULL,
  `G_batting` int(11) DEFAULT NULL,
  `G_defense` int(11) DEFAULT NULL,
  `G_p` int(11) DEFAULT NULL,
  `G_c` int(11) DEFAULT NULL,
  `G_1b` int(11) DEFAULT NULL,
  `G_2b` int(11) DEFAULT NULL,
  `G_3b` int(11) DEFAULT NULL,
  `G_ss` int(11) DEFAULT NULL,
  `G_lf` int(11) DEFAULT NULL,
  `G_cf` int(11) DEFAULT NULL,
  `G_rf` int(11) DEFAULT NULL,
  `G_of` int(11) DEFAULT NULL,
  `G_dh` int(11) DEFAULT NULL,
  `G_ph` int(11) DEFAULT NULL,
  `G_pr` int(11) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`teamID`,`playerID`,`G_all`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awardsmanagers`
--

DROP TABLE IF EXISTS `awardsmanagers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awardsmanagers` (
  `playerID` varchar(255) NOT NULL,
  `awardID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `lgID` varchar(255) NOT NULL,
  `tie` varchar(255) DEFAULT NULL,
  `notes` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`lgID`,`awardID`,`playerID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awardsplayers`
--

DROP TABLE IF EXISTS `awardsplayers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awardsplayers` (
  `playerID` varchar(255) NOT NULL,
  `awardID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `lgID` varchar(255) NOT NULL,
  `tie` varchar(255) DEFAULT NULL,
  `notes` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`lgID`,`awardID`,`playerID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awardssharemanagers`
--

DROP TABLE IF EXISTS `awardssharemanagers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awardssharemanagers` (
  `awardID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `lgID` varchar(255) NOT NULL,
  `playerID` varchar(255) NOT NULL,
  `pointsWon` int(11) DEFAULT NULL,
  `pointsMax` int(11) DEFAULT NULL,
  `votesFirst` int(11) DEFAULT NULL,
  PRIMARY KEY (`awardID`,`yearID`,`lgID`,`playerID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `awardsshareplayers`
--

DROP TABLE IF EXISTS `awardsshareplayers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `awardsshareplayers` (
  `awardID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `lgID` varchar(255) NOT NULL,
  `playerID` varchar(255) NOT NULL,
  `pointsWon` decimal(10,5) DEFAULT NULL,
  `pointsMax` int(11) DEFAULT NULL,
  `votesFirst` decimal(10,5) DEFAULT NULL,
  PRIMARY KEY (`awardID`,`yearID`,`lgID`,`playerID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `batting`
--

DROP TABLE IF EXISTS `batting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `batting` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `stint` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `G` int(11) DEFAULT NULL,
  `AB` int(11) DEFAULT NULL,
  `R` int(11) DEFAULT NULL,
  `H` int(11) DEFAULT NULL,
  `doubles` int(11) DEFAULT NULL,
  `triples` int(11) DEFAULT NULL,
  `HR` int(11) DEFAULT NULL,
  `RBI` int(11) DEFAULT NULL,
  `SB` int(11) DEFAULT NULL,
  `CS` int(11) DEFAULT NULL,
  `BB` int(11) DEFAULT NULL,
  `SO` int(11) DEFAULT NULL,
  `IBB` int(11) DEFAULT NULL,
  `HBP` int(11) DEFAULT NULL,
  `SH` int(11) DEFAULT NULL,
  `SF` int(11) DEFAULT NULL,
  `GIDP` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`teamID`,`stint`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `battingpost`
--

DROP TABLE IF EXISTS `battingpost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `battingpost` (
  `yearID` int(11) NOT NULL,
  `round` varchar(255) NOT NULL,
  `playerID` varchar(255) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `G` int(11) DEFAULT NULL,
  `AB` int(11) DEFAULT NULL,
  `R` int(11) DEFAULT NULL,
  `H` int(11) DEFAULT NULL,
  `doubles` int(11) DEFAULT NULL,
  `triples` int(11) DEFAULT NULL,
  `HR` int(11) DEFAULT NULL,
  `RBI` int(11) DEFAULT NULL,
  `SB` int(11) DEFAULT NULL,
  `CS` int(11) DEFAULT NULL,
  `BB` int(11) DEFAULT NULL,
  `SO` int(11) DEFAULT NULL,
  `IBB` int(11) DEFAULT NULL,
  `HBP` int(11) DEFAULT NULL,
  `SH` int(11) DEFAULT NULL,
  `SF` int(11) DEFAULT NULL,
  `GIDP` int(11) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`round`,`playerID`,`teamID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `collegeplaying`
--

DROP TABLE IF EXISTS `collegeplaying`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `collegeplaying` (
  `playerID` varchar(255) NOT NULL,
  `schoolID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  PRIMARY KEY (`playerID`,`schoolID`,`yearID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `fielding`
--

DROP TABLE IF EXISTS `fielding`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fielding` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `stint` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `POS` varchar(255) NOT NULL,
  `A` int(11) DEFAULT NULL,
  `GS` int(11) DEFAULT NULL,
  `G` int(11) DEFAULT NULL,
  `PB` int(11) DEFAULT NULL,
  `InnOuts` int(11) DEFAULT NULL,
  `ZR` int(11) DEFAULT NULL,
  `PO` int(11) DEFAULT NULL,
  `WP` int(11) DEFAULT NULL,
  `CS` int(11) DEFAULT NULL,
  `E` int(11) DEFAULT NULL,
  `DP` int(11) DEFAULT NULL,
  `SB` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`teamID`,`POS`,`stint`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `fieldingof`
--

DROP TABLE IF EXISTS `fieldingof`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fieldingof` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `stint` int(11) NOT NULL,
  `Glf` int(11) DEFAULT NULL,
  `Gcf` int(11) DEFAULT NULL,
  `Grf` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`stint`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `fieldingofsplit`
--

DROP TABLE IF EXISTS `fieldingofsplit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fieldingofsplit` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `stint` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `POS` varchar(255) NOT NULL,
  `G` int(11) DEFAULT NULL,
  `GS` int(11) DEFAULT NULL,
  `InnOuts` int(11) DEFAULT NULL,
  `PO` int(11) DEFAULT NULL,
  `A` int(11) DEFAULT NULL,
  `E` int(11) DEFAULT NULL,
  `DP` int(11) DEFAULT NULL,
  `PB` int(11) DEFAULT NULL,
  `WP` int(11) DEFAULT NULL,
  `SB` int(11) DEFAULT NULL,
  `CS` int(11) DEFAULT NULL,
  `ZR` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`teamID`,`POS`,`stint`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `fieldingpost`
--

DROP TABLE IF EXISTS `fieldingpost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `fieldingpost` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `teamID` varchar(255) DEFAULT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `round` varchar(255) NOT NULL,
  `POS` varchar(255) NOT NULL,
  `G` int(11) DEFAULT NULL,
  `GS` int(11) DEFAULT NULL,
  `InnOuts` int(11) DEFAULT NULL,
  `PO` int(11) DEFAULT NULL,
  `A` int(11) DEFAULT NULL,
  `E` int(11) DEFAULT NULL,
  `DP` int(11) DEFAULT NULL,
  `TP` int(11) DEFAULT NULL,
  `PB` int(11) DEFAULT NULL,
  `SB` int(11) DEFAULT NULL,
  `CS` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`round`,`POS`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `halloffame`
--

DROP TABLE IF EXISTS `halloffame`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `halloffame` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `votedBy` varchar(255) NOT NULL,
  `ballots` int(11) DEFAULT NULL,
  `needed` int(11) DEFAULT NULL,
  `votes` int(11) DEFAULT NULL,
  `inducted` varchar(255) DEFAULT NULL,
  `category` varchar(255) NOT NULL,
  `needed_note` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`category`,`yearID`,`votedBy`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `homegames`
--

DROP TABLE IF EXISTS `homegames`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `homegames` (
  `yearkey` int(11) NOT NULL,
  `parkkey` varchar(255) NOT NULL,
  `teamkey` varchar(255) NOT NULL,
  `leaguekey` varchar(255) DEFAULT NULL,
  `attendance` int(11) DEFAULT NULL,
  `games` int(11) DEFAULT NULL,
  `spanfirst` varchar(255) DEFAULT NULL,
  `spanlast` varchar(255) DEFAULT NULL,
  `openings` int(11) DEFAULT NULL,
  PRIMARY KEY (`yearkey`,`parkkey`,`teamkey`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `managers`
--

DROP TABLE IF EXISTS `managers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `managers` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) NOT NULL,
  `inseason` int(11) NOT NULL,
  `G` int(11) DEFAULT NULL,
  `W` int(11) DEFAULT NULL,
  `L` int(11) DEFAULT NULL,
  `rank` int(11) DEFAULT NULL,
  `plyrMgr` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`lgID`,`teamID`,`playerID`,`inseason`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `managershalf`
--

DROP TABLE IF EXISTS `managershalf`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `managershalf` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) NOT NULL,
  `inseason` int(11) DEFAULT NULL,
  `half` int(11) NOT NULL,
  `G` int(11) DEFAULT NULL,
  `W` int(11) DEFAULT NULL,
  `L` int(11) DEFAULT NULL,
  `rank` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`teamID`,`lgID`,`yearID`,`half`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `parks`
--

DROP TABLE IF EXISTS `parks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `parks` (
  `parkkey` varchar(255) NOT NULL,
  `parkname` varchar(255) DEFAULT NULL,
  `parkalias` varchar(255) DEFAULT NULL,
  `alias` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`parkkey`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `people`
--

DROP TABLE IF EXISTS `people`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `people` (
  `playerID` varchar(9) NOT NULL,
  `birthYear` int(11) DEFAULT NULL,
  `birthMonth` int(11) DEFAULT NULL,
  `birthDay` int(11) DEFAULT NULL,
  `birthCountry` varchar(14) DEFAULT NULL,
  `birthState` varchar(22) DEFAULT NULL,
  `birthCity` varchar(26) DEFAULT NULL,
  `deathYear` int(11) DEFAULT NULL,
  `deathMonth` int(11) DEFAULT NULL,
  `deathDay` int(11) DEFAULT NULL,
  `deathCountry` varchar(14) DEFAULT NULL,
  `deathState` varchar(20) DEFAULT NULL,
  `deathCity` varchar(28) DEFAULT NULL,
  `nameFirst` varchar(12) DEFAULT NULL,
  `nameLast` varchar(14) DEFAULT NULL,
  `nameGiven` varchar(43) DEFAULT NULL,
  `weight` int(11) DEFAULT NULL,
  `height` int(11) DEFAULT NULL,
  `bats` varchar(255) DEFAULT NULL,
  `throws` varchar(255) DEFAULT NULL,
  `debut` varchar(255) DEFAULT NULL,
  `finalGame` varchar(255) DEFAULT NULL,
  `retroID` varchar(255) DEFAULT NULL,
  `bbrefID` varchar(9) DEFAULT NULL,
  PRIMARY KEY (`playerID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `pitching`
--

DROP TABLE IF EXISTS `pitching`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pitching` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `stint` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `W` int(11) DEFAULT NULL,
  `L` int(11) DEFAULT NULL,
  `G` int(11) DEFAULT NULL,
  `GS` int(11) DEFAULT NULL,
  `CG` int(11) DEFAULT NULL,
  `SHO` int(11) DEFAULT NULL,
  `SV` int(11) DEFAULT NULL,
  `IPouts` int(11) DEFAULT NULL,
  `H` int(11) DEFAULT NULL,
  `ER` int(11) DEFAULT NULL,
  `HR` int(11) DEFAULT NULL,
  `BB` int(11) DEFAULT NULL,
  `SO` int(11) DEFAULT NULL,
  `BAOpp` float DEFAULT NULL,
  `ERA` float DEFAULT NULL,
  `IBB` int(11) DEFAULT NULL,
  `WP` int(11) DEFAULT NULL,
  `HBP` int(11) DEFAULT NULL,
  `BK` int(11) DEFAULT NULL,
  `BFP` int(11) DEFAULT NULL,
  `GF` int(11) DEFAULT NULL,
  `R` int(11) DEFAULT NULL,
  `SH` int(11) DEFAULT NULL,
  `SF` int(11) DEFAULT NULL,
  `GIDP` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`stint`,`teamID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `pitchingpost`
--

DROP TABLE IF EXISTS `pitchingpost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pitchingpost` (
  `playerID` varchar(255) NOT NULL,
  `yearID` int(11) NOT NULL,
  `round` varchar(255) NOT NULL,
  `teamID` varchar(255) DEFAULT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `W` int(11) DEFAULT NULL,
  `L` int(11) DEFAULT NULL,
  `G` int(11) DEFAULT NULL,
  `GS` int(11) DEFAULT NULL,
  `CG` int(11) DEFAULT NULL,
  `SHO` int(11) DEFAULT NULL,
  `SV` int(11) DEFAULT NULL,
  `IPouts` int(11) DEFAULT NULL,
  `H` int(11) DEFAULT NULL,
  `ER` int(11) DEFAULT NULL,
  `HR` int(11) DEFAULT NULL,
  `BB` int(11) DEFAULT NULL,
  `SO` int(11) DEFAULT NULL,
  `BAOpp` float DEFAULT NULL,
  `ERA` float DEFAULT NULL,
  `IBB` int(11) DEFAULT NULL,
  `WP` int(11) DEFAULT NULL,
  `HBP` int(11) DEFAULT NULL,
  `BK` int(11) DEFAULT NULL,
  `BFP` int(11) DEFAULT NULL,
  `GF` int(11) DEFAULT NULL,
  `R` int(11) DEFAULT NULL,
  `SH` int(11) DEFAULT NULL,
  `SF` int(11) DEFAULT NULL,
  `GIDP` int(11) DEFAULT NULL,
  PRIMARY KEY (`playerID`,`yearID`,`round`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `salaries`
--

DROP TABLE IF EXISTS `salaries`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `salaries` (
  `yearID` int(11) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `playerID` varchar(255) NOT NULL,
  `salary` int(11) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`teamID`,`playerID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `schools`
--

DROP TABLE IF EXISTS `schools`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schools` (
  `schoolID` varchar(255) NOT NULL,
  `name_full` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `state` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`schoolID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `seriespost`
--

DROP TABLE IF EXISTS `seriespost`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `seriespost` (
  `yearID` int(11) NOT NULL,
  `round` varchar(255) NOT NULL,
  `teamIDwinner` varchar(255) DEFAULT NULL,
  `lgIDwinner` varchar(255) DEFAULT NULL,
  `teamIDloser` varchar(255) DEFAULT NULL,
  `lgIDloser` varchar(255) DEFAULT NULL,
  `wins` int(11) DEFAULT NULL,
  `losses` int(11) DEFAULT NULL,
  `ties` int(11) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`round`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `teams`
--

DROP TABLE IF EXISTS `teams`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teams` (
  `yearID` int(11) NOT NULL,
  `lgID` varchar(255) NOT NULL,
  `teamID` varchar(255) NOT NULL,
  `franchID` varchar(255) DEFAULT NULL,
  `divID` varchar(255) DEFAULT NULL,
  `Rank` int(11) DEFAULT NULL,
  `G` int(11) DEFAULT NULL,
  `Ghome` int(11) DEFAULT NULL,
  `W` int(11) DEFAULT NULL,
  `L` int(11) DEFAULT NULL,
  `DivWin` varchar(255) DEFAULT NULL,
  `WCWin` varchar(255) DEFAULT NULL,
  `LgWin` varchar(255) DEFAULT NULL,
  `WSWin` varchar(255) DEFAULT NULL,
  `R` int(11) DEFAULT NULL,
  `AB` int(11) DEFAULT NULL,
  `H` int(11) DEFAULT NULL,
  `doubles` int(11) DEFAULT NULL,
  `triples` int(11) DEFAULT NULL,
  `HR` int(11) DEFAULT NULL,
  `BB` int(11) DEFAULT NULL,
  `SO` int(11) DEFAULT NULL,
  `SB` int(11) DEFAULT NULL,
  `CS` int(11) DEFAULT NULL,
  `HBP` int(11) DEFAULT NULL,
  `SF` int(11) DEFAULT NULL,
  `RA` int(11) DEFAULT NULL,
  `ER` int(11) DEFAULT NULL,
  `ERA` float DEFAULT NULL,
  `CG` int(11) DEFAULT NULL,
  `SHO` int(11) DEFAULT NULL,
  `SV` int(11) DEFAULT NULL,
  `IPouts` int(11) DEFAULT NULL,
  `HA` int(11) DEFAULT NULL,
  `HRA` int(11) DEFAULT NULL,
  `BBA` int(11) DEFAULT NULL,
  `SOA` int(11) DEFAULT NULL,
  `E` int(11) DEFAULT NULL,
  `DP` int(11) DEFAULT NULL,
  `FP` float DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `park` varchar(255) DEFAULT NULL,
  `attendance` int(11) DEFAULT NULL,
  `BPF` int(11) DEFAULT NULL,
  `PPF` int(11) DEFAULT NULL,
  `teamIDBR` varchar(255) DEFAULT NULL,
  `teamIDlahman45` varchar(255) DEFAULT NULL,
  `teamIDretro` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`lgID`,`teamID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `teamsfranchises`
--

DROP TABLE IF EXISTS `teamsfranchises`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teamsfranchises` (
  `franchID` varchar(255) NOT NULL,
  `franchName` varchar(255) DEFAULT NULL,
  `active` varchar(255) DEFAULT NULL,
  `NAassoc` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`franchID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `teamshalf`
--

DROP TABLE IF EXISTS `teamshalf`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `teamshalf` (
  `yearID` int(11) NOT NULL,
  `lgID` varchar(255) DEFAULT NULL,
  `teamID` varchar(255) NOT NULL,
  `Half` int(11) NOT NULL,
  `divID` varchar(255) DEFAULT NULL,
  `DivWin` varchar(255) DEFAULT NULL,
  `Rank` int(11) DEFAULT NULL,
  `G` int(11) DEFAULT NULL,
  `W` int(11) DEFAULT NULL,
  `L` int(11) DEFAULT NULL,
  PRIMARY KEY (`yearID`,`teamID`,`Half`)
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

-- Dump completed on 2018-04-29 19:33:38
