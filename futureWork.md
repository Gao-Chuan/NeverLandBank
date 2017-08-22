web server of NeverLand bank
---

1. func

    - a language choose page(汉语/English)

    - users' registion(submit a issue to the managers. The account will be in use after managers' check)
    - users' log in / log out
    - users' check balance
    - users' transfer
        + auto-filled blank of recently tranfer account(a list)

    - managers' log in / log out (auto-distinguish. there shouldn't be a "manager-log-in block")
    - managers' management
        + check whether if an account can be built(show in managers' homepage after every refresh)
        + delete an acount(balance should be zero)
        + add/sub money to a specific account(submit a issue to the super-managers, and balance will change after super-managers' check)
    - mamagers' check balance
    - managers' transfer
        + auto-filled blank of recently tranfer account(a list)
    
    - super-managers' log in / log out (auto-distinguish. there shouldn't be a "super-manager-log-in block")
    - super-managers' management
        + check whether if a transfer can be execute(show in managers' homepage after every refresh. Show as a link, in which enclude the whole transfer chain)
        + give an account privilege to be a manager
    - check the whole users list(users & managers)(show as a link, in which ecnlude all message of the account)
    - check todays' transfer(each transfer show as a link, in which enclude the whole transfer chain)

2. property
    - property of a manager
        + by which super-manager
        + honest record(times of submission being rejected)

    - property of a traansfer
        + from/to
        + submited by which manager
        + ensured by which super-manager

3. sql
    - users
        + nickName
        + id
        + name
        + password(sha256)
            - salt
        + balance
        + gender
        + idCard
        + fromWhichManager
    - managers
        + nickName
        + id
        + name
        + password(sha256)
            - salt
        + gender
        + idCard
        + privateKey
        + publicKey
        + honestRecord
        + fromWhichSuperManager
    - superManagers
        + nickName
        + id
        + name
        + password(sha256)
            - salt
        + gender
        + idCard
        + privateKey
        + publicKey
        + ensuredManagers(less than 10)
    - transfer
        + transferID
        + from
        + to
        + amount
        + timeStamp
        + submitedByManager(only when "from" equal to NULL)
        + signedBySuper-manager1(only when "from" equal to NULL)
        + signedBySuper-manager2(only when "from" equal to NULL)

CREATE TABLE IF NOT EXISTS `users`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `nickName` VARCHAR(20) NOT NULL,
   `firstname` VARCHAR(20) NOT NULL,
   `lastname` VARCHAR(20) NOT NULL,
   `password` VARBINARY(32) NOT NULL,
   `salt` VARBINARY(32) NOT NULL,
   `balance` INT NOT NULL,
   `idcard` VARCHAR(20) NOT NULL,
   `gender` VARCHAR(10) NOT NULL,
   `fromWhichManager` INT UNSIGNED NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `users_temp`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `nickName` VARCHAR(20) NOT NULL,
   `firstname` VARCHAR(20) NOT NULL,
   `lastname` VARCHAR(20) NOT NULL,
   `password` VARBINARY(32) NOT NULL,
   `salt` VARBINARY(32) NOT NULL,
   `balance` INT NOT NULL,
   `idcard` VARCHAR(20) NOT NULL,
   `gender` VARCHAR(10) NOT NULL,
   `fromWhichManager` INT UNSIGNED,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `managers`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `nickName` VARCHAR(20) NOT NULL,
   `firstname` VARCHAR(20) NOT NULL,
   `lastname` VARCHAR(20) NOT NULL,
   `password` VARBINARY(32) NOT NULL,
   `salt` VARBINARY(32) NOT NULL,
   `idcard` VARCHAR(20) NOT NULL,
   `gender` VARCHAR(10) NOT NULL,
   `privateKey` VARBINARY(150) NOT NULL,
   `publicKey` VARBINARY(100) NOT NULL,
   `honestRecord` INT UNSIGNED NOT NULL,
   `fromWhichSuperManager` INT UNSIGNED NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `managers_temp`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `nickName` VARCHAR(20) NOT NULL,
   `firstname` VARCHAR(20) NOT NULL,
   `lastname` VARCHAR(20) NOT NULL,
   `password` VARBINARY(32) NOT NULL,
   `salt` VARBINARY(32) NOT NULL,
   `idcard` VARCHAR(20) NOT NULL,
   `gender` VARCHAR(10) NOT NULL,
   `fromWhichSuperManager` INT UNSIGNED NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `superManagers`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `nickName` VARCHAR(20) NOT NULL,
   `firstname` VARCHAR(20) NOT NULL,
   `lastname` VARCHAR(20) NOT NULL,
   `password` VARBINARY(32) NOT NULL,
   `salt` VARBINARY(32) NOT NULL,
   `idcard` VARCHAR(20) NOT NULL,
   `gender` VARCHAR(10) NOT NULL,
   `privateKey` VARBINARY(150) NOT NULL,
   `publicKey` VARBINARY(100) NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE IF NOT EXISTS `transfer`(
   `transferID` INT UNSIGNED AUTO_INCREMENT,
   `fromID` INT UNSIGNED,
   `toID` INT UNSIGNED NOT NULL,
   `amount` INT UNSIGNED NOT NULL,
   `time` DATE NOT NULL,
   `submitedByManger` INT UNSIGNED NOT NULL,
   `signedBySuperManager1` INT UNSIGNED NOT NULL,
   `signedBySuperManager2` INT UNSIGNED NOT NULL,
   PRIMARY KEY ( `transferID` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `transfer_temp`(
   `transferID` INT UNSIGNED AUTO_INCREMENT,
   `fromID` INT UNSIGNED,
   `toID` INT UNSIGNED NOT NULL,
   `amount` INT UNSIGNED NOT NULL,
   `time` DATE NOT NULL,
   `submitedByManger` INT UNSIGNED NOT NULL,
   `signedBySuperManager1` INT UNSIGNED,
   `signedBySuperManager2` INT UNSIGNED,
   PRIMARY KEY ( `transferID` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
