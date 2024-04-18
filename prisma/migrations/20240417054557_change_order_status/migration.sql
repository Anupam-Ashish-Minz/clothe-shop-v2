/*
  Warnings:

  - The values [COMPLETED,PENDING] on the enum `OrderStatus` will be removed. If these variants are still used in the database, this will fail.

*/
-- AlterEnum
BEGIN;
CREATE TYPE "OrderStatus_new" AS ENUM ('DELIVERED', 'PROCESSING', 'CANCLED', 'OUT_FOR_DELIVERY');
ALTER TABLE "Order" ALTER COLUMN "state" DROP DEFAULT;
ALTER TABLE "Order" ALTER COLUMN "state" TYPE "OrderStatus_new" USING ("state"::text::"OrderStatus_new");
ALTER TYPE "OrderStatus" RENAME TO "OrderStatus_old";
ALTER TYPE "OrderStatus_new" RENAME TO "OrderStatus";
DROP TYPE "OrderStatus_old";
ALTER TABLE "Order" ALTER COLUMN "state" SET DEFAULT 'PROCESSING';
COMMIT;

-- AlterTable
ALTER TABLE "Order" ALTER COLUMN "state" SET DEFAULT 'PROCESSING';
