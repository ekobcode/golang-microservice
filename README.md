# golang-microservice

golang-microservice/
├── cmd/
│   └── server/
│       └── main.go                  # 🏁 Entry point (bootstrap app)
│
├── internal/
│   ├── entity/                      # 🧠 Enterprise business objects
│   │   └── user.go
│   │
│   ├── usecase/                     # 💡 Application business rules
│   │   └── user_usecase.go
│   │
│   ├── repository/                  # 💾 Data access implementation
│   │   └── user_repository.go
│   │
│   ├── delivery/                    # 🌐 Interface Adapters (Gin handlers)
│   │   └── http/
│   │       ├── handler/
│   │       │   └── user_handler.go
│   │       └── router.go
│   │
│   ├── infrastructure/              # ⚙️ Frameworks & drivers
│   │   ├── db/
│   │   │   └── postgres.go
│   │   ├── logger/
│   │   │   └── zap_logger.go
│   │   ├── config/
│   │   │   └── config.go
│   │   └── middleware/
│   │       └── auth.go
│   │
│   ├── app/                         # 🧩 Dependency injection (wire up layers)
│   │   └── app.go
│   │
│   └── utils/                       # 🛠️ Helper utilities
│       └── response.go
│
├── pkg/                             # 📦 Reusable packages (optional)
│   └── pagination/
│       └── paginate.go
│
├── test/                            # 🧪 Unit / integration tests
│   └── user_usecase_test.go
│
└── go.mod
