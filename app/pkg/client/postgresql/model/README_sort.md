# Сортировки

Для создания сортировки используется функция `NewSort`

При неверном типе сортировки переданном в `NewSort` будет использоваться тип сортировки `ASC`

> На данный момент использование Alias не допускается

Примеры:

```go
// Сортировка по дате создания

// Создание сортировки по дате создания
sort := NewSort("created_at", "asc")

// ...

builder := squirrel.Select("id").From("products")
builder = sort.UseSelectBuilder(builder)

```

```sql
SELECT id FROM products ORDER BY created_at ASC
```

---

```go
// Сортировка по рейтингу и цене

// Создание сортировки по рейтингу
rating := NewSort("rating", "asc")
// Создание сортировки цене
price := NewSort("price", "desc")

// ...

builder := squirrel.Select("id").From("products")
builder = rating.UseSelectBuilder(builder)
builder = price.UseSelectBuilder(builder)

```

```sql
SELECT id FROM products ORDER BY rating ASC, price DESC
```
