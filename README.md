## Small Go projects to practice core language concepts

---

### 1. web-server

Serves static files from `./static`, a form handler on `/form`, and `/hello` on port **8080**.

```mermaid
flowchart TD
  C[Browser / client] --> L["net/http :8080"]
  L --> P{URL path}
  P -->|"/" + files| FS["FileServer ./static"]
  P -->|"/form"| FM["formHandler: ParseForm → FormValue name / address → response"]
  P -->|"/hello"| HL["helloHandler: path + GET check → hello or 404 / 405"]
```

---

### 2. movies-crud-postman

In-memory movie list with **Gorilla Mux** on port **8000**; JSON CRUD for `/movies` and `/movies/{id}`.

```mermaid
flowchart TD
  P[Postman / JSON client] --> R["Gorilla Mux :8000"]
  R --> G1["GET /movies → all movies"]
  R --> G2["GET /movies/{id} → one by id"]
  R --> CR["POST /movies → decode body → random id → append"]
  R --> UP["PUT /movies/{id} → remove old → decode → append same id"]
  R --> DL["DELETE /movies/{id} → splice slice"]
  G1 --> S[(movies slice in memory)]
  G2 --> S
  CR --> S
  UP --> S
  DL --> S
```

---

### 3. bookstore-api

REST API with **Gorilla Mux**, **GORM**, and **MySQL** on **localhost:9010**.

```mermaid
flowchart TD
  C[HTTP client] --> S["ListenAndServe localhost:9010"]
  S --> R[mux.Router]
  R --> RT["Routes: POST+GET /book/, GET+PUT /book/{id}, DELETE /books/{id}"]
  RT --> H[Controllers: CreateBook, GetBookById, UpdateBook, DeleteBook]
  H --> U["utils.ParseBody on JSON where needed"]
  H --> M["models: GORM CRUD"]
  M --> DB[(MySQL)]
```
