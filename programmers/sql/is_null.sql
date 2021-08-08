-- 이름이 없는 동물의 아이디
SELECT animal_id from animal_ins where name is null;

-- 이름이 있는 동물의 아이디
SELECT animal_id from animal_ins where name is not null;

-- NULL 처리하기
SELECT animal_type, ifnull(name, 'No name') as name, sex_upon_intake from animal_ins order by animal_id;
