-- 없어진 기록 찾기
SELECT ANIMAL_OUTS.ANIMAL_ID, ANIMAL_OUTS.NAME FROM ANIMAL_OUTS
LEFT OUTER JOIN ANIMAL_INS ON ANIMAL_OUTS.ANIMAL_ID = ANIMAL_INS.ANIMAL_ID
WHERE ANIMAL_INS.ANIMAL_ID IS NULL
ORDER BY ANIMAL_OUTS.ANIMAL_ID


  
-- 있었는데요 없었습니다
SELECT 
    ANIMAL_INS.ANIMAL_ID
    , ANIMAL_INS.NAME 
FROM ANIMAL_INS 
    LEFT JOIN ANIMAL_OUTS 
    ON ANIMAL_INS.ANIMAL_ID = ANIMAL_OUTS.ANIMAL_ID 
WHERE ANIMAL_INS.DATETIME > ANIMAL_OUTS.DATETIME 
ORDER BY ANIMAL_INS.DATETIME


  
-- 오랜 기간 보호한 동물(1)
SELECT ANIMAL_INS.NAME, ANIMAL_INS.DATETIME FROM ANIMAL_INS
LEFT JOIN ANIMAL_OUTS ON ANIMAL_INS.ANIMAL_ID=ANIMAL_OUTS.ANIMAL_ID
WHERE ANIMAL_OUTS.DATETIME IS NULL
ORDER BY ANIMAL_INS.DATETIME LIMIT 3


  
-- 보호소에서 중성화한 동물
SELECT ANIMAL_OUTS.ANIMAL_ID, ANIMAL_OUTS.ANIMAL_TYPE, ANIMAL_OUTS.NAME 
FROM ANIMAL_OUTS
LEFT JOIN ANIMAL_INS ON ANIMAL_OUTS.ANIMAL_ID=ANIMAL_INS.ANIMAL_ID
WHERE ANIMAL_INS.SEX_UPON_INTAKE LIKE 'Intact%' 
AND (ANIMAL_OUTS.SEX_UPON_OUTCOME LIKE 'Spayed%' OR ANIMAL_OUTS.SEX_UPON_OUTCOME LIKE 'Neutered%')
ORDER BY ANIMAL_OUTS.ANIMAL_ID



-- 없어진 기록 찾기
SELECT ANIMAL_OUTS.ANIMAL_ID, ANIMAL_OUTS.NAME FROM ANIMAL_OUTS
LEFT OUTER JOIN ANIMAL_INS ON ANIMAL_OUTS.ANIMAL_ID = ANIMAL_INS.ANIMAL_ID
WHERE ANIMAL_INS.ANIMAL_ID IS NULL
ORDER BY ANIMAL_OUTS.ANIMAL_ID



-- 5월 식품들의 총매출 조회하기
SELECT
    fp.PRODUCT_ID
    ,fp.PRODUCT_NAME
    ,(fp.PRICE * fo.AMOUNT) AS TOTAL_SALES
FROM FOOD_PRODUCT fp
    JOIN (
        SELECT PRODUCT_ID, SUM(AMOUNT) AS AMOUNT
        FROM FOOD_ORDER 
        WHERE PRODUCE_DATE LIKE '2022-05%'
        GROUP BY PRODUCT_ID
    ) fo
    ON fp.PRODUCT_ID = fo.PRODUCT_ID
ORDER BY TOTAL_SALES DESC, PRODUCT_ID ASC



-- 주문량이 많은 아이스크림들 조회하기
SELECT
    fh.FLAVOR
FROM  FIRST_HALF fh
    JOIN (
        SELECT
            FLAVOR
            ,SUM(TOTAL_ORDER) AS j_sum_order
        FROM  JULY
        GROUP BY FLAVOR
    ) j
    ON fh.FLAVOR = j.FLAVOR
GROUP BY fh.FLAVOR
ORDER BY (SUM(fh.TOTAL_ORDER) + j.j_sum_order) DESC
LIMIT 3



-- 그룹별 조건에 맞는 식당 목록 출력하기
SELECT
    mp.MEMBER_NAME
    ,rr.REVIEW_TEXT
    ,DATE_FORMAT(rr.REVIEW_DATE, '%Y-%m-%d')AS REVIEW_DATE
FROM REST_REVIEW rr
    JOIN (
        SELECT MEMBER_ID, COUNT(MEMBER_ID) FROM REST_REVIEW
        GROUP BY MEMBER_ID
        HAVING COUNT(MEMBER_ID) = (
            SELECT COUNT(MEMBER_ID) as review_count
            FROM REST_REVIEW
            GROUP BY MEMBER_ID
            ORDER BY COUNT(MEMBER_ID) DESC
            LIMIT 1
        )
    ) rrc
        ON rrc.MEMBER_ID = rr.MEMBER_ID
    JOIN (SELECT MEMBER_ID, MEMBER_NAME FROM MEMBER_PROFILE) mp
        ON mp.MEMBER_ID = rr.MEMBER_ID
ORDER BY rr.REVIEW_DATE ASC, rr.REVIEW_TEXT ASC



-- 특정 기간동안 대여 가능한 자동차들의 대여비용 구하기
SELECT
    ccc.CAR_ID
    ,ccc.CAR_TYPE
    ,ROUND((ccc.DAILY_FEE * (1 - crcdp.DISCOUNT_RATE/100) * 30), 0) AS FEE
FROM CAR_RENTAL_COMPANY_CAR ccc
    JOIN (
        SELECT CAR_TYPE, DISCOUNT_RATE
        FROM CAR_RENTAL_COMPANY_DISCOUNT_PLAN 
        WHERE DURATION_TYPE = '30일 이상'
    ) crcdp
    ON ccc.CAR_TYPE = crcdp.CAR_TYPE
WHERE 
    ccc.CAR_TYPE IN ('세단', 'SUV')
    AND ccc.CAR_ID NOT IN (
        SELECT distinct CAR_ID
        FROM CAR_RENTAL_COMPANY_RENTAL_HISTORY
        WHERE START_DATE <= '2022-11-01' AND END_DATE >= '2022-11-01'
    )
    AND ROUND((ccc.DAILY_FEE * (1 - crcdp.DISCOUNT_RATE/100) * 30), 0) >= 500000 
    AND ROUND((ccc.DAILY_FEE * (1 - crcdp.DISCOUNT_RATE/100) * 30), 0) < 2000000
ORDER BY FEE DESC, ccc.CAR_TYPE ASC, ccc.CAR_ID DESC
