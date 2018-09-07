## MYSQL

### 一、导出数据库用mysqldump命令

#### 1、 导出数据和表结构：

mysqldump -u用户名 -p密码 数据库名 > 数据库名.sql
mysqldump -uroot -p dbname > dbname .sql

#### 2、只导出表结构

mysqldump -u用户名 -p密码 -d 数据库名 > 数据库名.sql
mysqldump -uroot -p -d dbname > dbname .sql

### 二、导入数据库

#### 1、首先建空数据库

mysql>create database dbname ;

#### 2、导入数据库

方法一：

（1）选择数据库
mysql>use dbname ;
（2）设置数据库编码
mysql>set names utf8;
（3）导入数据（注意sql文件的路径）
mysql>source /home/xxxx/dbname .sql;

方法二：

mysql -u用户名 -p密码 数据库名 < 数据库名.sql


### 查看数据库编码

USE your_database_name;
show variables like "character_set_database";


**查看表格编码**

SHOW FULL COLUMNS FROM table_name;


SELECT SCHEMA_NAME 'database', default_character_set_name 'charset', DEFAULT_COLLATION_NAME 'collation' FROM information_schema.SCHEMATA;

### 设置数据库编码

ALTER DATABASE DBNAME CHARACTER SET utf8 COLLATE utf8_general_ci;