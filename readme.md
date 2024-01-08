# EmpathiCare
**EmpathiCare** dimulai dari sebuah gagasan sebagai respons terhadap kebutuhan mendesak akan dukungan kesehatan mental yang mudah diakses dan terjangkau.

**EmpathiCare** tidak hanya menjembatani pengguna dengan psikolog profesional secara online, tetapi juga menyediakan sumber daya edukatif yang menyeluruh untuk memberikan panduan dan dukungan kepada pengguna.

## 3rd Party App
1. Midtrans untuk Payment Gateway
2. Cloudinary untuk Penyimpanan File dan Gambar
3. OpenAI untuk Chatbot
4. WebSocket untuk Real-Time Chat
5. Database menggunakan MySql dan Mongo
6. OAuth untuk Login Google Mail (Development)

## Link Pendukung 
[Github Project](https://github.com/orgs/Capstone-Mental-Health-Care-Group-6/projects/12/views/1)
[API Documentation POSTMAN](https://documenter.getpostman.com/view/30980878/2s9YXh6NZz)
[Database Design](https://app.diagrams.net/#G1kej_yzLmiJ9pcl238qr_C9B7lN9VHKFw#%7B%22pageId%22%3A%22p-vC0pzHpxmxFeUUWXhf%22%7D)

## REST API Design
Request
GET {{BASE_URL}}/counseling/methods/1

Response 
```json
{
  "data":[
    {
    "id":1,
    "name":"Chat",
    "additional_price":20000
    }
  ],
  "message":"Success Get Data"
}
```

## Project Structure
Clean Architecture

## Product Success End User
MVP | Nilai Plus | Tambahan 
--- | --- | ---
View Available Doctors | Manual Transfer | View Package Counseling
Counseling | Payment Gateway | 
Chatbot AI | Chat With Doctor |
Chatbot CS | | 
Article | |

## Product Success Doctor
MVP | Nilai Plus | Tambahan 
--- | --- | ---
Manage Patients | Chat With End User | Withdraw Balance 
View Patient Complaints | Notification | 
Chatbot AI | | 
Article | | 

## Product Success Admin (Additional)
|Fitur | Fitur|
|---|---|
| Manage Users | Manage Transactions |
| Manage Doctors | Manage Withdraw |
| Manage Article | Manage Bundle Counseling |
