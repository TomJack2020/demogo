-- 选品结果字段数据提取
a.publish_id,
if(b.short_name is null, '', b.short_name) short_name,
if(b.shop_name is null, '', b.shop_name) shop_name