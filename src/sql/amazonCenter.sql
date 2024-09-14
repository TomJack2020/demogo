select a.account_id as account_id, short_name, a.site_code as site_code ,
       case
           when publish_type = 1 then "系统生成"
           when publish_type = 2 then "手动刊登"
           when publish_type = 3 then "JP清关补刊"
           when publish_type = 4 then "人工导入"
           when publish_type = 5 then "算法刊登"
           when publish_type = 6 then "链接翻新"
           when publish_type = 7 then "复制刊登"
           when publish_type = 10 then "导入文案智刊"
           when publish_type = 11 then "复制刊登-泛欧补FBM"
           ELSE "其他"
           end as publish_type,
       FROM_UNIXTIME(last_syn_time,'%Y-%m-%d') as calculate_date,
       count(publish_id) as publish_success_num
from yibai_sale_center_amazon.yibai_amazon_publish_success a
         left join yibai_sale_center_system.yibai_system_account b
                   on a.account_id  = b.id
where last_syn_time BETWEEN UNIX_TIMESTAMP('2024-08-01') and UNIX_TIMESTAMP('2024-08-02')
group by account_id, short_name, publish_type, site_code,calculate_date