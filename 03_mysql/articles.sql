-- phpMyAdmin SQL Dump
-- version 4.9.7
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Mar 01, 2021 at 01:12 AM
-- Server version: 5.7.32
-- PHP Version: 7.4.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `go_database`
--

-- --------------------------------------------------------

--
-- Table structure for table `articles`
--

CREATE TABLE `articles` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(64) NOT NULL,
  `content` text NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `articles`
--

INSERT INTO `articles` (`id`, `title`, `content`) VALUES
(1, 'The First Article', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(2, 'Lipsum', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(3, 'Golang Tutorial', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(4, 'MySQL Tutorial', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(5, 'This Is Really Really Really Long Title', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(6, 'This is <b>Title</b>', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(7, 'a', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(8, '#234', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(9, ';SHOW TABLES;', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.'),
(10, '^_^t', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis commodo imperdiet accumsan. Duis et eros luctus, aliquet libero quis, dapibus ante. Nunc maximus luctus felis vel elementum. Suspendisse maximus, neque a gravida vulputate, velit est sollicitudin dolor, sit amet mollis magna libero ut leo. Nulla commodo lectus semper felis fringilla, vel dignissim leo dapibus. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Integer consectetur interdum elementum. Morbi scelerisque fermentum massa eu lacinia. Nunc eu fermentum arcu, eu aliquam ipsum.');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `articles`
--
ALTER TABLE `articles`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `articles`
--
ALTER TABLE `articles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
