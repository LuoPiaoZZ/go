# good-good-study

学习笔记

常用命令行
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
go mod init xxxx

git reset hard   56e4cf5ae089a985716bcedb5fea852b8b0dbcf0

CREATE USER qipai WITH PASSWORD 'xxxxx';

psql -d postgres
pg_dump -U postgres statdb >/tmp/dbsql/b.sql
psql -d statdb -U qipai  -f statdb.sql

pg_dump -U postgres statdb >/tmp/b.sql

create user  xxxx with password ‘*****’;
ALTER ROLE xxxx CREATEROLE SUPERUSER; 

SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE datname=‘要关闭的数据库名称’ AND pid<>pg_backend_pid();

CREATE SEQUENCE user_prop_everyday_id_seq
    START WITH 1 
    INCREMENT BY 1 
    NO MINVALUE 
    NO MAXVALUE 
    CACHE 1; 
    
alter table agent_announcements alter column id set default nextval('agent_announcements_id_seq'); 

ALTER TABLE payment_channel_config add PRIMARY KEY(ID);

ALTER TABLE room_card_config ADD COLUMN price_meal jsonb NOT NULL;

redis-cli
keys *阻塞线程 
scan 0 不阻塞线程，不影响线上，数据可能更新不及时


sudo  supervisorctl restart opdataapisrv

ps -ef | grep opdataapisrv

pref top -p xxx
