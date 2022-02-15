# BlogOnContainers


Kullanıcıların blog gönderilerini görüntülemesine/yönetmesine olanak tanıyan REST API'lerin oluşturulması. Mevcut tüm blog gönderilerini ve gönderi oluşturma, güncelleme ve silme işlemlerini göstermek için API yöntemleri sağlanması.


### Çalıştırma (Running)
---
`docker compose up` -> //container'ları kaldırmak için kullanabilirsiniz. <br>
`docker compose down` -> //container'ları silmek için kullanabilirsiniz. Image'lar silinmez.


### Teknik Detaylar
---
- Database olarak `mongoDB` kullanılmıştır.
- Web Framework olarak  `Gin` kullanımıştır.
- `JWT (JSON Web Tokens)` kullanılmıtır.

### Kullanım
---
`POST /user        # Kullanıcı Kayıt` <br> 
`POST /login # Giriş` <br>
`POST /api/blog # Hikaye Oluşturma`  <br> 
`POST /api/blog/:id # Hikaye Güncelleme`  <br> 
`DELETE /api/blog/:id # Hikaye Silme`  <br> 
`GET /api/blog/ # Tüm Hikayeleri Alma`  <br> 
`GET /api/blog/:id # Belirli Bir Hikaye Alma`  <br> 

##### Templates
`http://localhost:5000/loginpage # Giriş için UI`  <br> 
`http://localhost:5000/registerpage # Kayıt için UI`  <br> 
`http://localhost:5000/storypage # Hikaye için UI`  <br> 




##### Request User Example

`   curl -X POST http://localhost:5000/user
   -H 'Content-Type: application/json'
   -d '{"UserName": "admin","Password": "admin"}' 
   `
###### Response

```
{
    "Data": null,
    "Status": 200,
    "Error": null,
    "Message": "User created"
} 
```

##### Request Login Example
`   curl -X POST http://localhost:5000/login
   -H 'Content-Type: application/json'
   -d '{"UserName": "admin","Password": "admin"}'
   `
###### Response
```
{
    "Data": "{ JWT }",
    "Status": 200,
    "Error": null,
    "Message": "token created"
} 
```

##### Request Story Example
`   curl -X POST http://localhost:5000/api/blog
   -H 'Content-Type: application/json'
   -d '{"Title":"Örnek Başlık","Content":"Örnek İçerik "}'
   `
###### Response
```
{
    "Data": {
        "Title": "Örnek Başlık",
        "Content": "Örnek İçerik "
    },
    "Status": 201,
    "Error": null,
    "Message": "story created"
} 
```
