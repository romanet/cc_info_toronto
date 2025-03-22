# Make executable 

go build

# Duckdb usage (supports PIVOT)

```sql

INSTALL sqlite;

LOAD sqlite;

--SELECT * FROM sqlite_scan('C:\\Projects\\go_projects\\cc_info_toronto\\cc_sqlite.db', 'centres');

ATTACH 'C:\\Projects\\go_projects\\cc_info_toronto\\cc-info-toronto.db' AS test (TYPE sqlite);
-- The table "tbl_name" can now be queried as if it is a regular table
 
-- Switch the active database to "test"
USE test;
-- List all tables in the file
SHOW TABLES;
 

WITH dd AS (
        SELECT *
          FROM (VALUES ('monday', 0), ('tuesday', 1), ('wednesday',2), ('thursday',3), ('friday',4), ('saturday',5), ('sunday',6)) aa(day_of_week, num)
       ),  my_loc AS (
        SELECT 45.7764411 lat,
               -75.4786722 lng
       ),
       base AS (
        SELECT round(acos(sin(my_loc.lat::float * (pi()/180)) * sin(c.lat::float * (pi()/180)) + cos(my_loc.lat::float *(pi()/180)) * cos(c.lat::float * (pi()/180)) *cos(c.lng::float*(pi()/180)-my_loc.lng::float*(pi()/180)))*6371,1) AS distance_km,
               s.centre_id,
               s.week_start_date::DATE + dd.num scheduler_day,
               s.sport,
               age,
               program_time,
               status ,
               c.name,
               c.address,
               c.phone,
               c.district,
               c.lng,
               c.lat
          FROM  program_scheduler_records s
          JOIN centres c
            ON s.centre_id = c.id
          JOIN dd
            ON s.day_of_week = dd.day_of_week
        join my_loc ON 1=1
         WHERE sport = 'Table Tennis'
           AND age not like '%60%'
         ORDER BY scheduler_day
       ) pivot base
    ON scheduler_day USING max(program_time) order by distance_km ;


```
