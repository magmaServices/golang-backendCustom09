/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

DROP TABLE IF EXISTS `heroes`;
CREATE TABLE IF NOT EXISTS `heroes` (
  `hero_id` int(11) NOT NULL AUTO_INCREMENT,
  `hero_name` varchar(50) NOT NULL,
  `player_id` int(11) NOT NULL,
  `hero_stats` text NOT NULL,
  PRIMARY KEY (`hero_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;

/*!40000 ALTER TABLE `heroes` DISABLE KEYS */;
INSERT INTO `heroes` (`hero_id`, `hero_name`, `player_id`, `hero_stats`) VALUES
	(9, 'FirstHero', 9, '{"c_ft":"","c_team":"","c_hrc":"","c_hrs":"","c_skc":"","c_ltp":"9337.0000","c_ltm":"9337.0000","c_fhrs":"","c_slm":"2.0000","cdm":"","edm":"","c_kit":"","c_wallet_hero":"","c_wallet_valor":"","games":"","elo":"","level":"","xp":"","ct":"","ki":"","dt":"","su":"","win":"","los":"","fi":"","hi":"","rs":"","ts":"","ss":"","cs":"","prs":"","ppt":"","c_tut":"","awybt":"","dmc":"","gsco":"","expts":"","bnspt":"","aw":{},"mid":{"0":"6000.0000","1":"6000.0000","2":"6000.0000"},"c_mid":{},"c_cmid":{},"m0c":{},"m1c":{},"m2c":{},"c_wmid":{"0":"0.0000","1":"0.0000","2":"0.0000"},"startLVL":"","roundXP":"","roundBXP":"","roundVP":"","roundBVP":"","roundHP":"","roundPP":"","roundBPP":"","totalPP":"","cpc":"","cpa":"","cpd":"","rc":"","ks":"","ds":"","ft_rs":{},"ft_ki":{},"ft_dt":{},"ft_win":{},"ft_los":{},"fc_rs":{},"fc_ki":{},"fc_dt":{},"fc_win":{},"fc_los":{},"m_ct":{},"m_win":{},"m_los":{},"tv":{},"kv":{},"dfv":{},"kvr":{},"dstrv":{},"div":{},"tw":{},"twk":{},"kw":{},"dfw":{},"sw":{},"hw":{},"dww":{},"kk":{},"kkb":{},"ka":"","he":"","drka":"","c_apr":["10","1330"],"c_emo":[],"c_eqp":["3167","0","0","3155","0","0","0","0","0","0"],"c_items":[],"ige":{}}'),
	(10, 'SecondHero', 9, '{"c_ft":"","c_team":"","c_hrc":"","c_hrs":"","c_skc":"","c_ltp":"","c_ltm":"","c_fhrs":"","c_slm":"","cdm":"","edm":"","c_kit":"","c_wallet_hero":"","c_wallet_valor":"","games":"","elo":"","level":"","xp":"","ct":"","ki":"","dt":"","su":"","win":"","los":"","fi":"","hi":"","rs":"","ts":"","ss":"","cs":"","prs":"","ppt":"","c_tut":"","awybt":"","dmc":"","gsco":"","expts":"","bnspt":"","aw":{},"mid":{},"c_mid":{},"c_cmid":{},"m0c":{},"m1c":{},"m2c":{},"c_wmid":{},"startLVL":"","roundXP":"","roundBXP":"","roundVP":"","roundBVP":"","roundHP":"","roundPP":"","roundBPP":"","totalPP":"","cpc":"","cpa":"","cpd":"","rc":"","ks":"","ds":"","ft_rs":{},"ft_ki":{},"ft_dt":{},"ft_win":{},"ft_los":{},"fc_rs":{},"fc_ki":{},"fc_dt":{},"fc_win":{},"fc_los":{},"m_ct":{},"m_win":{},"m_los":{},"tv":{},"kv":{},"dfv":{},"kvr":{},"dstrv":{},"div":{},"tw":{},"twk":{},"kw":{},"dfw":{},"sw":{},"hw":{},"dww":{},"kk":{},"kkb":{},"ka":"","he":"","drka":"","c_apr":[],"c_emo":[],"c_eqp":[],"c_items":[],"ige":{}}'),
	(11, 'OtherPlayerHero', 11, '{"c_ft":"","c_team":"","c_hrc":"","c_hrs":"","c_skc":"","c_ltp":"9337.0000","c_ltm":"9337.0000","c_fhrs":"","c_slm":"2.0000","cdm":"","edm":"","c_kit":"","c_wallet_hero":"","c_wallet_valor":"","games":"","elo":"","level":"","xp":"","ct":"","ki":"","dt":"","su":"","win":"","los":"","fi":"","hi":"","rs":"","ts":"","ss":"","cs":"","prs":"","ppt":"","c_tut":"","awybt":"","dmc":"","gsco":"","expts":"","bnspt":"","aw":{},"mid":{},"c_mid":{},"c_cmid":{},"m0c":{},"m1c":{},"m2c":{},"c_wmid":{},"startLVL":"","roundXP":"","roundBXP":"","roundVP":"","roundBVP":"","roundHP":"","roundPP":"","roundBPP":"","totalPP":"","cpc":"","cpa":"","cpd":"","rc":"","ks":"","ds":"","ft_rs":{},"ft_ki":{},"ft_dt":{},"ft_win":{},"ft_los":{},"fc_rs":{},"fc_ki":{},"fc_dt":{},"fc_win":{},"fc_los":{},"m_ct":{},"m_win":{},"m_los":{},"tv":{},"kv":{},"dfv":{},"kvr":{},"dstrv":{},"div":{},"tw":{},"twk":{},"kw":{},"dfw":{},"sw":{},"hw":{},"dww":{},"kk":{},"kkb":{},"ka":"","he":"","drka":"","c_apr":[],"c_emo":[],"c_eqp":["3184","0","0","3155","0","0","0","0","0","0"],"c_items":[],"ige":{}}');
/*!40000 ALTER TABLE `heroes` ENABLE KEYS */;

DROP TABLE IF EXISTS `players`;
CREATE TABLE IF NOT EXISTS `players` (
  `player_id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `game_token` varchar(50) DEFAULT NULL,
  `prefer_server` int(11) DEFAULT NULL,
  `selected_hero_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`player_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;

/*!40000 ALTER TABLE `players` DISABLE KEYS */;
INSERT INTO `players` (`player_id`, `username`, `password`, `game_token`, `prefer_server`, `selected_hero_id`) VALUES
	(9, 'SomeUser', 'admin1', 'topsecret', NULL, 9),
	(11, 'OtherUse', 'password', 'moresecret', NULL, 11);
/*!40000 ALTER TABLE `players` ENABLE KEYS */;

DROP TABLE IF EXISTS `servers`;
CREATE TABLE IF NOT EXISTS `servers` (
  `server_id` int(11) NOT NULL AUTO_INCREMENT,
  `soldier_name` varchar(50) DEFAULT NULL,
  `account_username` varchar(50) DEFAULT NULL,
  `account_password` varchar(50) DEFAULT NULL,
  `api_key` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`server_id`)
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=latin1;

/*!40000 ALTER TABLE `servers` DISABLE KEYS */;
INSERT INTO `servers` (`server_id`, `soldier_name`, `account_username`, `account_password`, `api_key`) VALUES
	(123, 'Test-Server', 'Test-Server', 'Test-Server', 'SERVER-APIKEY');
/*!40000 ALTER TABLE `servers` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
