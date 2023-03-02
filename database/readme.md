# Database Design



## List Table
- User
- Product
- Cart
- History  -> copy from cart



```json
// table user

user:{
    "id": "u1",
    "name": "anton",
    "email": "anton@gmail.com",
    "password": "rahasia234",
    "address": "solo",
    "gender": "male",
    "phone": 082323847923,
}
```
```json
// table product

product:{
    "id": "p1",
    "name": "headset logitech",
    "description": "weojfaofdsfhaf",
    "price": 400000,
    "category": "headset",
    "rating": 5,
    "stock": 49, // addition 
    "sold": 32, // addition
    "image": "url",
    "uidImage": "url"
}


product:{
    "id": "p2",
    "name": "headset logitech",
    "description": "weojfaofdsfhaf",
    "price": 400000,
    "category": "headset",
    "rating": 5,
    "stock": 49, // addition 
    "sold": 32, // addition
    "image": "url",
    "uidImage": "url"
}


product:{
    "id": "p3",
    "name": "headset logitech",
    "description": "weojfaofdsfhaf",
    "price": 400000,
    "category": "headset",
    "rating": 5,
    "stock": 49, // addition 
    "sold": 32, // addition
    "image": "url",
    "uidImage": "url"
}


product:{
    "id": "p4",
    "name": "headset logitech",
    "description": "weojfaofdsfhaf",
    "price": 400000,
    "category": "headset",
    "rating": 5,
    "stock": 49, // addition 
    "sold": 32, // addition
    "image": "url",
    "uidImage": "url"
}


product:{
    "id": "p5",
    "name": "headset logitech",
    "description": "weojfaofdsfhaf",
    "price": 400000,
    "category": "headset",
    "rating": 5,
    "stock": 49, // addition 
    "sold": 32, // addition
    "image": "url",
    "uidImage": "url"
}


```
## Product 
- Memiliki `stock` karena ketika user melakukan pembayaran maka `stock` akan berkurang. Admin sewaktu-waktu bisa menambahkan `stock` product dengan fitur update
- Memiliki `sold` karena ketika user melakukan pembayarn maka `sold` akan bertambah. `sold` tidak bisa diupdate manual

```json
// table cart

cart:{
    "id": "c1",
    "id_user": "u1",
    "id_product": "p1",
    "qty": 7,
    "is_checked": true
}

cart:{
    "id": "c2",
    "id_user": "u1",
    "id_product": "p2",
    "qty": 3,
    "is_checked": true
}

cart:{
    "id": "c3",
    "id_user": "u1",
    "id_product": "p3",
    "qty": 2,
    "is_checked": false
}

cart:{
    "id": "c4",
    "id_user": "u1",
    "id_product": "p4",
    "qty": 1,
    "is_checked": false
}

```

## Cart
- Menjadi table penghubung antara table user dengan table product (many to many)
- Pertama melakukan findAll `cart` yang memiliki ID user yang telah login
- User menambahkan product ke `cart` dengan memilih ID product
- Memiliki default `is_checked` = false
- total harga product didapat dari harga product dikali quantity
- ketika klik checkout maka data `is_checked` akan terupdate. Dan halaman checkout mengambil data cart yang memiliki `is_checked` = true. Oleh karena itu table checkout tidak dibutuhkan
- total harga keseluruhan didapat dari penjumlahan keseluruhan yang memiliki `is_checked` = true

## Payment
- Ketika berhasil melakukan pembayaran, maka semua data cart, product, dan user akan disalin di table `history` serta menambahkan field `review` yang bisa diupdate user
- `stock` product berkurang, `sold` product bertambah, dan cart yang memiliki `is_checked` = true akan dihapus. Tujuannya adalah agar ketika harga suatu product di update maka harga yang di table `history` tidak ikut di ubah, atau istilahnya embeded.
Oleh karena itu, field `is_paid` tidak diperlukan
- informasi `not yet paid` di dashboard admin didapat dari user yang sudah checkout tapi belum dibayar, atau cart yang `is_checked`nya true


```json
// table example

history:{
    "id": "h1",
    "user_id": "u1",
    "product_id": "p1",
    "name": "headset logitech",
    "description": "weojfaofdsfhaf",
    "price": 400000,
    "category": "headset",
    "image": "url",
    "uidImage": "url",
    "qty": 7
}

history:{
    "id": "h2",
    "user_id": "u1",
    "product_id": "p2",
    "name": "headset logitech",
    "description": "weojfaofdsfhaf",
    "price": 400000,
    "category": "headset",
    "image": "url",
    "uidImage": "url",
    "qty": 3
}

```

## History
- Tanpa memiliki field `stock` dan `sold` karena sewaktu-waktu bisa berubah, dan datanya tidak terlalu penting untuk table history
- Menyalin semua data dari cart, product, dan user
- Informasi `not yet paid` berkurang sejumlah product yang sudah dibayar
- Table history hanya melakukan relasi dengan table user
