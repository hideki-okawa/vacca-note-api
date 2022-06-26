USE vacca_note_db_local;
INSERT INTO notes
    (name, gender, age, vaccine_type, number_of_vaccination, max_temperature, log, remarks, good_count, created_at, updated_at)
    VALUES
    ("匿名", "1", "2", "M", 2, "6", "テストログ", "備考テスト", 1, "2019-08-08 15:00", "2019-08-08 18:00");