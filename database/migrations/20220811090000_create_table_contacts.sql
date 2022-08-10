-- +goose Up
CREATE TABLE `contacts` (
    `contact_id` INT unsigned NOT NULL AUTO_INCREMENT COMMENT '連絡先 ID',
    `family_name` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '姓（漢字）',
    `first_name` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '名（漢字）',
    `family_name_kana` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '姓（カナ）',
    `first_name_kana` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '名（カナ）',
    `phone_number` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '電話番号',
    `postal_code` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '郵便番号',
    `prefecture_code` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '都道府県コード',
    `city_code` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '市区町村コード',
    `address_line1` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '町丁目・番地',
    `address_line2` VARCHAR(50) COLLATE utf8mb4_bin NOT NULL COMMENT '建物名・部屋番号',
    `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT 'レコード作成日時',
    `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'レコード更新日時',
    PRIMARY KEY (`contact_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '連絡先';

-- +goose Down
DROP TABLE IF EXISTS contacts;
