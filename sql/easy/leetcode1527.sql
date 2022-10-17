-- 患某种疾病的患者


SELECT * FROM Patients  p where p.conditions like 'DIAB1%' or  p.conditions like '% DIAB1%'