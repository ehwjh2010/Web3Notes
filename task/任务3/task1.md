
假设有一个名为 students 的表，</br>
包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。</br>
要求 ：</br>
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。</br>
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。</br>
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。</br>
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。</br>

```sql
INSERT INTO `students`(`name`, `age`, `grade`) VALUES ('张三', 20, '三年级');
SELECT * FROM `students` where age > 18;
UPDATE students set grade = '四年级' where name = '张三';
delete from students where age < 15;
```


