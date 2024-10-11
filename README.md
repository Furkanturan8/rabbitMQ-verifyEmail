# RabbitMQ ile Email Doğrulama Projesi

RabbitMQ kullanarak kullanıcıların e-posta adreslerini doğrular. Kullanıcı kaydı sırasında, verilen e-posta adresine bir doğrulama bağlantısı gönderilmektedir.

## Bileşenleri

- **Producer**: Kullanıcı kayıtlarını alır ve doğrulama bağlantısını RabbitMQ kuyruğuna gönderir.
- **Consumer**: RabbitMQ kuyruğundan gelen mesajları dinler ve işleme alır (örneğin, e-posta göndermek için).

### Register Get Request:
![foto1](https://github.com/Furkanturan8/rabbitMQ-verifyEmail/blob/main/fotos/register.png)

### Producer log:
![foto2](https://github.com/Furkanturan8/rabbitMQ-verifyEmail/blob/main/fotos/producer.png)

### Consumer log:
![foto3](https://github.com/Furkanturan8/rabbitMQ-verifyEmail/blob/main/fotos/consumer.png)

### Verify:
![foto3](https://github.com/Furkanturan8/rabbitMQ-verifyEmail/blob/main/fotos/verify.png)

