-- 首先删除外键约束
ALTER TABLE "transfers" DROP CONSTRAINT IF EXISTS "transfers_to_account_id_fkey";
ALTER TABLE "transfers" DROP CONSTRAINT IF EXISTS "transfers_from_account_id_fkey";
ALTER TABLE "entries" DROP CONSTRAINT IF EXISTS "entris_account_id_fkey";

-- 然后删除索引
DROP INDEX IF EXISTS "transfers_from_account_id_to_account_id_idx";
DROP INDEX IF EXISTS "transfers_to_account_id_idx";
DROP INDEX IF EXISTS "transfers_from_account_id_idx";
DROP INDEX IF EXISTS "entris_account_id_idx";
DROP INDEX IF EXISTS "accounts_owner_idx";

-- 最后删除表
DROP TABLE IF EXISTS "transfers";
DROP TABLE IF EXISTS "entries";
DROP TABLE IF EXISTS "accounts";