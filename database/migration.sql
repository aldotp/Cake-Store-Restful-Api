CREATE TABLE `cakes` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(300) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `rating` float DEFAULT NULL,
  `image` text DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO `cakes` (`id`, `title`, `description`, `rating`, `image`, `created_at`, `updated_at`) VALUES ('1', 'Lemon cheesecake', 'A cheesecake made of lemon', '7', 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg', '2020-02-01 10:56:31', '2020-02-13 09:30:23');

INSERT INTO `cakes` (`id`, `title`, `description`, `rating`, `image`, `created_at`, `updated_at`) VALUES ('2', 'Blue Cheese Cake', 'Good taste with blue color', '8', 'https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg', '2020-02-01 10:56:31', '2020-02-13 09:30:23');