# Фильтрация

Для создания фильтра используется функция `NewFilter`
Для добавления дополнительных фильтраций к текущей используется функция `Filter.WithFilters`
Все дополнительные фильтры будут соединяться в зависимости от параметра `Operator` переданного в `Filter.SetOperator` (`OperatorAnd` или `OperatorOr`) по умолчанию устанавливается значение `OperatorAnd`

> На данный момент использование Alias не допускается

Примеры:

```go
// Фильтрация по id

// Создание фильтра по id
filter := NewFilter("id", FilterTypeEQ, "00000000-0000-0000-0000-000000000000")

// ...

builder := squirrel.Select("id").From("products")
builder = filter.UseSelectBuilder(builder)
```

```sql
SELECT id FROM products WHERE id = '00000000-0000-0000-0000-000000000000'
```

---

```go
// Фильтрация по цене меньше либо равной 1000 и рейтингу больше 3.5

// Создание фильтра по рейтингу выше 3.5
rating := NewFilter("rating", FilterTypeGT, 3.5)
// Создание фильтра по цене меньше либо равной 1000
price := NewFilter("price", FilterTypeLTE, 1000)
// Добавление фильтра с рейтингом к фильтру с ценой
price.WithFilters(rating)

// ...

builder := squirrel.Select("id").From("products")
builder = price.UseSelectBuilder(builder)
```

```sql
SELECT id FROM products WHERE (price <= 1000 AND rating > 3.5)
```

---

```go
// Фильтрация по цене меньше либо равной 1000 и рейтингу 4 или 5

// Создание фильтра по рейтингу равному 5
rating5 := NewFilter("rating", FilterTypeEQ, 5)
// Создание фильтра по рейтингу равному 4
rating4 := NewFilter("rating", FilterTypeEQ, 4)
// Установка оператора OR и добавление фильтра с рейтингом 5 к фильтру с рейтингом 4
rating4.SetOperator(OperatorOr).WithFilters(rating5)
// Создание фильтра по цене меньше либо равной 1000
price := NewFilter("price", FilterTypeLTE, 1000)
// Добавление фильтра с рейтингом 4 к фильтру с ценой
price.WithFilters(rating4)

// ...

builder := squirrel.Select("id").From("products")
builder = price.UseSelectBuilder(builder)
```

```sql
SELECT id FROM products WHERE (price <= 1000 AND (rating = 4 OR rating = 5))
```

---

```go
// Фильтрация по приблизительному названию и входжению в определенные категории и ценой меньше 1000

// Создание фильтра по цене меньше  1000
price := NewFilter("price", FilterTypeLT, 1000)
// Создание фильтра по определенным категориям
categories := NewFilter("category_id", FilterTypeEQ, []int{1, 2, 3})
// Создание фильтра по приблизительному названию
name := NewFilter("name", FilterTypeLike, "%тел%")
// Добавление фильтров с ценой и категориями к фильтру с названием
name.WithFilters(price, categories)

// ...

builder := squirrel.Select("id").From("products")
builder = name.UseSelectBuilder(builder)
```

```sql
SELECT id FROM products WHERE (name LIKE '%тел%' AND price < 1000 AND category_id IN (1,2,3))
```
