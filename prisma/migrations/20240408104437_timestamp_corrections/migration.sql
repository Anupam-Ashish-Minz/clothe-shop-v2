-- AlterTable
ALTER TABLE "Order" ALTER COLUMN "date" SET DEFAULT CURRENT_TIMESTAMP,
ALTER COLUMN "date" SET DATA TYPE TIMESTAMPTZ(3);
