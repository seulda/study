-- 모든 레코드 조회하기
SELECT * FROM ANIMAL_INS ORDER BY ANIMAL_ID;



-- 역순 정렬하기
SELECT NAME, DATETIME FROM ANIMAL_INS ORDER BY ANIMAL_ID DESC;



-- 아픈 동물 찾기
SELECT ANIMAL_ID, NAME FROM ANIMAL_INS WHERE INTAKE_CONDITION = 'Sick' ORDER BY ANIMAL_ID;



-- 어린 동물 찾기
SELECT ANIMAL_ID, NAME FROM ANIMAL_INS where INTAKE_CONDITION != 'Aged' order by ANIMAL_ID;



-- 동물의 아이디와 이름
SELECT animal_id, name from animal_ins order by animal_id;



-- 여러 기준으로 정렬하기
SELECT animal_id, name, datetime from animal_ins order by name asc, datetime desc;



-- 상위 n개 레코드
SELECT name from animal_ins order by datetime asc limit 1;




-- 서울에 위치한 식당 목록 출력하기
SELECT
    ri.REST_ID
    ,ri.REST_NAME
    ,ri.FOOD_TYPE
    ,ri.FAVORITES
    ,ri.ADDRESS
    ,ROUND(AVG(rr.REVIEW_SCORE),2) AS SCORE
FROM REST_INFO ri
    JOIN (SELECT REST_ID, REVIEW_SCORE FROM REST_REVIEW) rr
    ON ri.REST_ID = rr.REST_ID
WHERE
    ri.ADDRESS LIKE '서울%'
GROUP BY REST_ID
ORDER BY SCORE DESC, FAVORITES DESC



-- 오프라인/온라인 판매 데이터 통합하기
SELECT
    DATE_FORMAT(SALES_DATE, '%Y-%m-%d') AS SALES_DATE
    ,PRODUCT_ID
    ,USER_ID
    ,SALES_AMOUNT
FROM ONLINE_SALE
WHERE DATE_FORMAT(SALES_DATE, '%Y-%m') = '2022-03'
UNION ALL (
    SELECT
        DATE_FORMAT(SALES_DATE, '%Y-%m-%d') AS SALES_DATE
        ,PRODUCT_ID
        ,NULL AS USER_ID
        ,SALES_AMOUNT
    FROM OFFLINE_SALE
    WHERE DATE_FORMAT(SALES_DATE, '%Y-%m') = '2022-03'
    ORDER BY SALES_DATE, PRODUCT_ID, USER_ID
)
ORDER BY SALES_DATE, PRODUCT_ID, USER_ID
