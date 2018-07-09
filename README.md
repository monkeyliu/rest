# rest<br>
GET :根据url带的id查询user信息 <br>
url="http://localhost:8080/api/v1/userinfo/"<br>

请求地址为http://localhost:8080/api/v1/userinfo/<br>

DELETE :根据url带的id删除用户信息<br>
url="http://localhost:8080/api/v1/userinfo/"<br>

请求地址为http://localhost:8080/api/v1/userinfo/<br>

PUT :根据url的id，请求传的json数据，修改对应的信息<br>
url="http://localhost:8080/api/v1/userinfo/{id}"<br>

请求地址为http://localhost:8080/api/v1/userinfo/<br>
{<br>
"ID":"27"<br>
   "NAME": "tom",<br>
        "AGE": 19,<br>
        "SEX": 7<br>
        }<br>
        
POST :根据请求传的json参数，增加对应信息<br>
url="http://localhost:8080/api/v1/userinfo<br>

请求地址为http://localhost:8080/api/v1/userinfo<br>
{<br>
"ID":"27"<br>
   "NAME": "tom",<br>
        "AGE": 19,<br>
        "SEX": 7<br>
        }<br>
        
   mysql的建表文件在dbinit文件夹的schema.sql文件中，您需要将改文件复制到clcos(没有自行创建一个同名数据库)数据库中使用改建表语句进行建表。
