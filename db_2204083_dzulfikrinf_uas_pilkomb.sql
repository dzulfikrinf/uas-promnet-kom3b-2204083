-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 07, 2024 at 11:13 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_2204083_dzulfikrinf_uas_pilkomb`
--

-- --------------------------------------------------------

--
-- Table structure for table `peminjamanbuku_dzulfikrinf`
--

CREATE TABLE `peminjamanbuku_dzulfikrinf` (
  `id` int(11) NOT NULL,
  `judul_buku` varchar(150) NOT NULL,
  `jumlah` int(11) NOT NULL,
  `nama_peminjam` varchar(60) NOT NULL,
  `alamat_peminjam` text NOT NULL,
  `noHp_peminjam` varchar(20) NOT NULL,
  `tanggal_pinjam` date NOT NULL,
  `tanggal_pengembalian` date NOT NULL,
  `lama_pinjam` varchar(25) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `peminjamanbuku_dzulfikrinf`
--

INSERT INTO `peminjamanbuku_dzulfikrinf` (`id`, `judul_buku`, `jumlah`, `nama_peminjam`, `alamat_peminjam`, `noHp_peminjam`, `tanggal_pinjam`, `tanggal_pengembalian`, `lama_pinjam`) VALUES
(1, 'Harry Potter and the Sorcerer\'s Stone', 2, 'John Doe', '123 Main St, City, Country', '123456789', '2023-12-01', '2023-12-15', '14 hari'),
(2, 'To Kill a Mockingbird', 2, 'Alice Johnson', '456 Elm St, Town, Country', '987654321', '2023-11-15', '2023-12-05', '20 hari'),
(3, '1984', 3, 'Bob Smith', '789 Oak St, Village, Country', '111222333', '2023-12-10', '2023-12-30', '20 hari'),
(4, 'The Great Gatsby', 2, 'Emily Brown', '567 Pine St, Hamlet, Country', '444555666', '2023-12-05', '2023-12-20', '15 hari'),
(5, 'Pride and Prejudice', 3, 'Sarah Wilson', '890 Maple St, City, Country', '777888999', '2023-11-25', '2023-12-10', '15 hari');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `peminjamanbuku_dzulfikrinf`
--
ALTER TABLE `peminjamanbuku_dzulfikrinf`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `peminjamanbuku_dzulfikrinf`
--
ALTER TABLE `peminjamanbuku_dzulfikrinf`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
