```
forum/
├── .dockerignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go                 # Entry point: initializes DB and starts server
├── database/               # Database initialization and migrations
│   ├── forum.db            # SQLite database file (gitignored)
│   └── schema.sql          # Initial CREATE TABLE queries
├── pkg/                    # Backend Logic
│   ├── auth/               # Cookies, Session, Bcrypt logic
│   ├── db/                 # SQLite drivers and query execution
│   └── models/             # Data structures (User, Post, Comment)
├── internal/               # Private application code
│   ├── handlers/           # HTTP Handlers (LoginHandler, PostHandler)
│   ├── repository/         # SQL SELECT/INSERT/UPDATE logic
│   └── middleware/         # Auth check, Logging, Error handling
├── ui/                     # "Frontend" files
│   ├── static/             # Assets
│   │   ├── css/            # Stylesheets
│   │   └── js/             # Vanilla JS (if needed for filtering)
│   └── templates/          # HTML Templates (Go html/template)
│       ├── base.layout.tmpl
│       ├── home.page.tmpl
│       ├── login.page.tmpl
│       └── post.page.tmpl
└── tests/                  # Unit tests for your logic
```
