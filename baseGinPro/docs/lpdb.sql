-- ----------------------------
-- 创建数据库
-- ----------------------------
CREATE DATABASE lpdb;
-- ----------------------------
-- 使用该数据库以后再执行下列语句
-- ----------------------------

-- ----------------------------
-- 创建序列
-- ----------------------------
CREATE SEQUENCE user_prop_everyday_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
-- ----------------------------
-- 表结构
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_prop_everyday";
CREATE TABLE "public"."user_prop_everyday" (
  "id" int4 NOT NULL DEFAULT nextval('user_prop_everyday_id_seq'::regclass),
  "game_id" int4 NOT NULL,
  "user_id" int4 NOT NULL,
  "prop_type" int4 NOT NULL,
  "prop_num" int4 NOT NULL DEFAULT 0,
  "element_type" int4 NOT NULL DEFAULT 0,
  log_date date not null default current_date,
  log_time int not null default extract(epoch from now())::int
)
;
-- ----------------------------
-- 表字段索引结构
-- ----------------------------
CREATE INDEX "user_prop_everyday_game_id_user_id_prop_type_idx" ON "public"."user_prop_everyday" USING btree (
  "game_id" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "user_id" "pg_catalog"."int4_ops" ASC NULLS LAST,
  "prop_type" "pg_catalog"."int4_ops" ASC NULLS LAST
);
-- ----------------------------
-- 主键结构
-- ----------------------------
ALTER TABLE "public"."user_prop_everyday" ADD CONSTRAINT "user_prop_everyday_pkey" PRIMARY KEY ("id");