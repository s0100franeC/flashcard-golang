-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Czas generowania: 25 Lis 2018, 10:16
-- Wersja serwera: 10.1.36-MariaDB
-- Wersja PHP: 7.2.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Baza danych: `flashcardsdb`
--

-- --------------------------------------------------------

--
-- Struktura tabeli dla tabeli `phrase`
--

CREATE TABLE `phrase` (
  `id` int(10) UNSIGNED NOT NULL,
  `jack` int(10) UNSIGNED NOT NULL,
  `native` varchar(100) NOT NULL,
  `translation` varchar(150) DEFAULT NULL,
  `courseName` varchar(50) DEFAULT NULL,
  `category` varchar(50) DEFAULT NULL,
  `natLanguage` varchar(50) NOT NULL,
  `counter` int(25) UNSIGNED NOT NULL DEFAULT '0',
  `correct` int(25) UNSIGNED NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Zrzut danych tabeli `phrase`
--

INSERT INTO `phrase` (`id`, `jack`, `native`, `translation`, `courseName`, `category`, `natLanguage`, `counter`, `correct`) VALUES
(11, 1, 'Soft', 'miękki', 'Easy English', 'Adjective', 'English', 0, 0),
(24, 3, 'Sofa', 'kanapa', 'Easy English', 'Furniture', 'English', 0, 0),
(25, 4, 'Chair', 'krzesło', 'Easy English', 'Furniture', 'English', 0, 0),
(26, 5, 'People', 'ludzie', 'Easy English', 'people', 'English', 0, 0),
(27, 11, 'Mercy', 'dziękuję', 'French for begginers', 'zwroty', 'French', 0, 0),
(28, 12, 'Chat', 'kot', 'French for begginers', 'Animal', 'French', 0, 0),
(29, 13, 'Parler', 'mówić', 'French for begginers', 'Verb', 'French', 0, 0),
(30, 14, 'Maison', 'dom', 'French for begginers', 'building', 'French', 0, 0),
(31, 15, 'Vie', 'życie', 'French for begginers', '', 'French', 0, 0),
(32, 21, 'der Mann', 'mężczyzna', 'die Sprache', 'people', 'German', 0, 0),
(33, 22, 'der Apfel', 'jabłko', 'die Sprache', 'fruit', 'German', 0, 0),
(34, 23, 'der Hund', 'pies', 'die Sprache', 'Animal', 'German', 0, 0),
(35, 24, 'die Zeit', 'czas', 'die Sprache', '', 'German', 0, 0),
(36, 25, 'Gesundheit!', 'na zdrowie!', 'die Sprache', 'zwroty', 'German', 0, 0),
(37, 133, 'amour', 'love', 'French for begginers', 'zsdfzsdfz', 'French', 0, 0);

--
-- Indeksy dla zrzutów tabel
--

--
-- Indeksy dla tabeli `phrase`
--
ALTER TABLE `phrase`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `jack` (`jack`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT dla tabeli `phrase`
--
ALTER TABLE `phrase`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=38;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
