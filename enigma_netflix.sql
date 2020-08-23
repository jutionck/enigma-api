-- MySQL dump 10.13  Distrib 8.0.20, for macos10.15 (x86_64)
--
-- Host: localhost    Database: enigma_netflix
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `film`
--

DROP TABLE IF EXISTS `film`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `film` (
  `id_film` int NOT NULL AUTO_INCREMENT,
  `film_title` varchar(50) NOT NULL,
  `film_duration` varchar(5) NOT NULL,
  `film_image_url` varchar(255) NOT NULL,
  `film_synopsis` text NOT NULL,
  `film_status` varchar(1) DEFAULT '1',
  PRIMARY KEY (`id_film`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `film`
--

LOCK TABLES `film` WRITE;
/*!40000 ALTER TABLE `film` DISABLE KEYS */;
INSERT INTO `film` VALUES (1,'Inglourious Basterds','153','https://m.media-amazon.com/images/M/MV5BOTJiNDEzOWYtMTVjOC00ZjlmLWE0NGMtZmE1OWVmZDQ2OWJhXkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_UX182_CR0,0,182,268_AL_.jpg','In Nazi-occupied France during World War II, a plan to assassinate Nazi leaders by a group of Jewish U.S. soldiers coincides with a theatre owner\'s vengeful plans for the same.','1'),(2,'Kill Bill: Vol. 1','111','https://m.media-amazon.com/images/M/MV5BNzM3NDFhYTAtYmU5Mi00NGRmLTljYjgtMDkyODQ4MjNkMGY2XkEyXkFqcGdeQXVyNzkwMjQ5NzM@._V1_UX182_CR0,0,182,268_AL_.jpg','After awakening from a four-year coma, a former assassin wreaks vengeance on the team of assassins who betrayed her.','1'),(3,'Pulp Fiction','154','https://m.media-amazon.com/images/M/MV5BNGNhMDIzZTUtNTBlZi00MTRlLWFjM2ItYzViMjE3YzI5MjljXkEyXkFqcGdeQXVyNzkwMjQ5NzM@._V1_UY268_CR1,0,182,268_AL_.jpg','The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.','1'),(4,'The Hateful Eight','168','https://m.media-amazon.com/images/M/MV5BMjA1MTc1NTg5NV5BMl5BanBnXkFtZTgwOTM2MDEzNzE@._V1_UX182_CR0,0,182,268_AL_.jpg','In the dead of a Wyoming winter, a bounty hunter and his prisoner find shelter in a cabin currently inhabited by a collection of nefarious characters.','1'),(5,'The World of IT','130','https://images-na.ssl-images-amazon.com/images/I/41f98NHHvlL._SX389_BO1,204,203,200_.jpg','Collecting the best artwork produced during the making of both of these sophisticated and visually enthralling films—including concept art, sketches, storyboards, and behind-the-scenes photography—The World of IT explores the films’ singular aesthetic and meticulous world-building. This compendium includes commentary from director Andy Muschietti; producer Barbara Muschietti; the acclaimed ensemble cast; and other creative players who helped bring a new, disturbing vision of King’s perennial bestseller to life. ','1'),(6,'Stranger Things','51','https://m.media-amazon.com/images/M/MV5BMjEzMDAxOTUyMV5BMl5BanBnXkFtZTgwNzAxMzYzOTE@._V1_UX182_CR0,0,182,268_AL_.jpg','When a young boy disappears, his mother, a police chief and his friends must confront terrifying supernatural forces in order to get him back.','1');
/*!40000 ALTER TABLE `film` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-08-23 20:33:55
