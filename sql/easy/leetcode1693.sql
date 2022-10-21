-- 每天的领导和合伙人

SELECT  ds.date_id,ds.make_name,count(DISTINCT ds.lead_id) unique_leads,count(DISTINCT ds.partner_id) as unique_partners
 FROM  DailySales ds 
GROUP BY ds.date_id,ds.make_name