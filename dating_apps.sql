-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 17 Agu 2024 pada 12.18
-- Versi server: 8.3.0
-- Versi PHP: 8.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `dating_apps`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `customers`
--

CREATE TABLE `customers` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_uuid` varchar(255) DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL,
  `bio` varchar(500) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `level` varchar(255) DEFAULT NULL,
  `swipe_quota` bigint DEFAULT NULL,
  `profile_picture` varchar(255) DEFAULT NULL,
  `package_expiry` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `customers`
--

INSERT INTO `customers` (`id`, `created_at`, `updated_at`, `deleted_at`, `customer_uuid`, `first_name`, `last_name`, `bio`, `email`, `password`, `level`, `swipe_quota`, `profile_picture`, `package_expiry`) VALUES
(1, '2024-08-17 10:45:11.349', '2024-08-17 10:45:11.349', NULL, '4c14aa79-6e13-4724-b3f5-d7a68b69ea37', 'Joni', 'Jona', 'Hey there', 'join@example.com', '$2a$14$BHH50fT/evcNjvO9iaz0N.itBH/mXqBmZYAaSfhHQOsfFs3fBPK5y', 'FREE', 10, '', NULL),
(2, '2024-08-17 10:54:25.405', '2024-08-17 15:25:58.754', NULL, 'a209a784-a402-49ad-9bee-0380e00e596e', 'Joni 1', 'Jona 1', 'Hey there', 'join1@example.com', '$2a$14$BO2ve6IwqJgpcNSu7RgE6OwhirA5PpPfdGo5VvgvJxrhm43CFYf4W', 'PREMIUM', -1, 'assets/profile-picture/95b93146-ffff-4b9c-b13c-8b2ba6f87a0f.png', NULL);

-- --------------------------------------------------------

--
-- Struktur dari tabel `orders`
--

CREATE TABLE `orders` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `order_no` varchar(255) DEFAULT NULL,
  `customer_id` bigint UNSIGNED DEFAULT NULL,
  `customer_name` varchar(255) DEFAULT NULL,
  `customer_email` varchar(255) DEFAULT NULL,
  `packages_id` bigint UNSIGNED DEFAULT NULL,
  `packages_title` varchar(255) DEFAULT NULL,
  `packages_quota` bigint DEFAULT NULL,
  `grand_total` float DEFAULT NULL,
  `order_status` varchar(55) DEFAULT NULL,
  `payment_status` varchar(55) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `orders`
--

INSERT INTO `orders` (`id`, `created_at`, `updated_at`, `deleted_at`, `order_no`, `customer_id`, `customer_name`, `customer_email`, `packages_id`, `packages_title`, `packages_quota`, `grand_total`, `order_status`, `payment_status`) VALUES
(2, '2024-08-17 15:25:58.738', '2024-08-17 15:25:58.738', NULL, 'ORD001', 2, 'Joni', 'join1@example.com', 2, 'Package Premium', -1, 10000, 'COMPLETED', 'PAID');

-- --------------------------------------------------------

--
-- Struktur dari tabel `packages`
--

CREATE TABLE `packages` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `quota` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `packages`
--

INSERT INTO `packages` (`id`, `created_at`, `updated_at`, `deleted_at`, `code`, `title`, `price`, `quota`) VALUES
(1, NULL, NULL, NULL, 'FREE', 'Package Free', 0, 10),
(2, NULL, NULL, NULL, 'PREMIUM', 'Package Premium', 10000, -1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `swipes`
--

CREATE TABLE `swipes` (
  `id` bigint UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customers_id` bigint UNSIGNED DEFAULT NULL,
  `swipe_customers_id` bigint UNSIGNED DEFAULT NULL,
  `swipe_at` datetime(3) DEFAULT NULL,
  `swipe_type` varchar(55) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `swipes`
--

INSERT INTO `swipes` (`id`, `created_at`, `updated_at`, `deleted_at`, `customers_id`, `swipe_customers_id`, `swipe_at`, `swipe_type`) VALUES
(7, '2024-08-17 12:02:11.242', '2024-08-17 12:02:11.242', NULL, 2, 1, '2024-08-17 12:02:11.242', 'LIKE'),
(8, '2024-08-17 12:13:22.376', '2024-08-17 12:13:22.376', NULL, 2, 0, '2024-08-17 12:13:22.374', 'LIKE');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_customers_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_orders_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `packages`
--
ALTER TABLE `packages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_packages_deleted_at` (`deleted_at`);

--
-- Indeks untuk tabel `swipes`
--
ALTER TABLE `swipes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_swipes_deleted_at` (`deleted_at`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `orders`
--
ALTER TABLE `orders`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `packages`
--
ALTER TABLE `packages`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `swipes`
--
ALTER TABLE `swipes`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
