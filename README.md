# Loyalty Point Agent: Backend

Aplikasi pembelian pulsa dan paket data dengan integrasi payment gateway dan makanisme poin. Project ini menggunakan clean architecture dengan layer handler, repository, service, dll. Dilengkapi dengan CI/CD pipeline yang memudahkan proses pengembangan, dan arsitektur yang cukup efektif--horizontal autoscaling, staging dan production server, caching, dan service yang saling terpisah.

<img src="https://firebasestorage.googleapis.com/v0/b/capstone-project-eede7.appspot.com/o/capstone.drawio_1.png?alt=media&token=4cbae5b5-5b8e-4b5a-a65d-8fc83fe4b222">

Dokumentasi RESTful API tersedia di [Swagger](https://app.swaggerhub.com/apis-docs/ANDIKABAHARI48_1/loyalty-point-agent/1.0.0#/).

## Fitur

- Registrasi dan autentikasi
- Manajemen user
- Manajemen produk (pulsa dan paket data)
- Transaksi dengan payment gateway
- Redeem pulsa, paket data, cashout, dan e-money
- Manajemen FAQ
- Verifikasi one-time password
- Ganti password dan lupa password
- Produk favorit
- Feedback pelanggan
- Analytics untuk dashboard

## Requirements

- Go v1.19
- Docker 20.10.20
- MySQL v8.0
- Redis v7.0
- Service account
- Xendit secret
- Mailjet key
- Goose v3.7.0 (opsional)
- Terraform v1.3.3 (opsional)

## Terraform Variables

Contoh pengisian file `terraform.tfvars`.

```
project = ""
region  = ""
zone    = ""

env_vars_staging = [
  { "name" : "APP_ENV", "value" : "production" },

  { "name" : "HTTP_PORT", "value" : "8080" },

  { "name" : "DB_DRIVER", "value" : "" },
  { "name" : "DB_USER", "value" : "" },
  { "name" : "DB_PASS", "value" : "" },
  { "name" : "DB_NAME", "value" : "" },
  { "name" : "DB_HOST", "value" : "" },
  { "name" : "DB_PORT", "value" : "" },

  { "name" : "AUTH_SECRET", "value" : "" },
  { "name" : "AUTH_COST", "value" : "" },
  { "name" : "AUTH_EXP_HOURS", "value" : "" },

  { "name" : "BUCKET", "value" : "" },
  { "name" : "TYPE", "value" : "" },
  { "name" : "PROJECT_ID", "value" : "" },
  { "name" : "PRIVATE_KEY_ID", "value" : "" },
  { "name" : "PRIVATE_KEY", "value" : "" },
  { "name" : "CLIENT_EMAIL", "value" : "" },
  { "name" : "CLIENT_ID", "value" : "" },

  { "name" : "MAIL_HOST", "value" : "" },
  { "name" : "MAIL_PORT", "value" : "" },
  { "name" : "MAIL_USERNAME", "value" : "" },
  { "name" : "MAIL_PASSWORD", "value" : "" },

  { "name" : "MAILJET_API_KEY", "value" : "" },
  { "name" : "MAILJET_SECRET_KEY", "value" : "" },
  { "name" : "MAILJET_SENDER_EMAIL", "value" : "" },
  { "name" : "MAILJET_SENDER_NAME", "value" : "" },

  { "name" : "XENDIT_SECRET", "value" : "" },

  { "name" : "REDIS_ADDR", "value" : "" },
  { "name" : "REDIS_PORT", "value" : "" },
  { "name" : "REDIS_PASSWORD", "value" : "" },
  { "name" : "REDIS_DB", "value" : "" },

  { "name" : "REACT_APP_BASE_URL", "value" : "" },
]

env_vars_production = [
  { "name" : "APP_ENV", "value" : "production" },

  { "name" : "HTTP_PORT", "value" : "8080" },

  { "name" : "DB_DRIVER", "value" : "" },
  { "name" : "DB_USER", "value" : "" },
  { "name" : "DB_PASS", "value" : "" },
  { "name" : "DB_NAME", "value" : "" },
  { "name" : "DB_HOST", "value" : "" },
  { "name" : "DB_PORT", "value" : "" },

  { "name" : "AUTH_SECRET", "value" : "" },
  { "name" : "AUTH_COST", "value" : "" },
  { "name" : "AUTH_EXP_HOURS", "value" : "" },

  { "name" : "BUCKET", "value" : "" },
  { "name" : "TYPE", "value" : "" },
  { "name" : "PROJECT_ID", "value" : "" },
  { "name" : "PRIVATE_KEY_ID", "value" : "" },
  { "name" : "PRIVATE_KEY", "value" : "" },
  { "name" : "CLIENT_EMAIL", "value" : "" },
  { "name" : "CLIENT_ID", "value" : "" },

  { "name" : "MAIL_HOST", "value" : "" },
  { "name" : "MAIL_PORT", "value" : "" },
  { "name" : "MAIL_USERNAME", "value" : "" },
  { "name" : "MAIL_PASSWORD", "value" : "" },

  { "name" : "MAILJET_API_KEY", "value" : "" },
  { "name" : "MAILJET_SECRET_KEY", "value" : "" },
  { "name" : "MAILJET_SENDER_EMAIL", "value" : "" },
  { "name" : "MAILJET_SENDER_NAME", "value" : "" },

  { "name" : "XENDIT_SECRET", "value" : "" },

  { "name" : "REDIS_ADDR", "value" : "" },
  { "name" : "REDIS_PORT", "value" : "" },
  { "name" : "REDIS_PASSWORD", "value" : "" },
  { "name" : "REDIS_DB", "value" : "" },

  { "name" : "REACT_APP_BASE_URL", "value" : "" },
]
```
