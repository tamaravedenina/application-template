# TODO
* спрятать папку google
* добавить миграции
* добавить пример использования api
* добавить гернератор grpc api в makefile, добавть в make generate как отдельный handler
* научиться генерировать все proto в папке api
* добавить db.close для баз
* обновить swagger-ui до полседней версии
* ?добавить make run вместе с запуском докером
* описать deploy
* добавить /live ручку, выводить там info приложения
* добавить /ready ручку, проверить там коннект к БД

# Структура приложения
```
| - api
| - api/{module}.proto
| - bin
| - bin/{project_name}_app
| - bin/{project_name}_migration
| - cmd/{project_name}/main.go 
| - config
| - config/local.yml
| - config/staging.yml
| - migrations
| - migrations/{migration_name}-{datetime}.go
| - internal
| - internal/app
| - internal/app/{module_name}
| - internal/app/{module_name}/datastruct
| - internal/app/{module_name}/datastruct/{struct_name}.go
| - internal/app/{module_name}/service
| - internal/app/{module_name}/service/service.go
| - internal/app/{module_name}/service/{business_logic_name}.go
| - internal/app/{module_name}/repository
| - internal/app/{module_name}/repository/repository.go
| - internal/app/{module_name}/repository/{business_logic_name}.go
| - internal/pkg
| - interanl/pkg/config
| - interanl/pkg/config/config.go
| - interanl/pkg/helper
| - interanl/pkg/helper/slice.go
| - interanl/pkg/helper/merge.go
| - interanl/pkg/helper/{business_logic_name}.go
| - interanl/pkg/api
| - interanl/pkg/api/{api_name}
| - interanl/pkg/api/{api_name}/{api_name}.proto
| - interanl/pkg/api/{api_name}/{api_name}.pb
| - interanl/pkg/api/{api_name}/client.go
| - interanl/pkg/api/{api_name}/{business_logic_name}.go
| - interanl/pkg/db/db.go
| - interanl/pkg/db/interface.go
| - tools/migrations/main.go
| - tools/swagger-ui/generate.go
| - tools/{business_logic_name}/main.go
| - .gitignore
| - .gitlab-ci.yml
| - go.mod
| - go.sum
| - Makefile
| - README.md
```

# Работа с приложением

# Работа с базой
## Что не должна делать БД
* заниматься валидацией данных
* ссылаться на другие сущности через foreign key
* заимствоавать бизнес логику с помощью тригеров
* абстрагировать запросы с помощью view 
* иметь схемы, использовать только public

## Что должна делать БД
* выполнять простые select/insert/upsert запросы

## Соглашение по архитектуре таблиц
* для всех ключей использовать ключевое слово key
* для названия связующих таблиц использовать шаблон {table1}_{table2}_link

# Работа с миграции

# Работа с репозиторием
* название ветки = номер задачи, git checkout -b TAXMON-1 -t origin/master
* не использовать force push origin {branch}
* ограничений по стратегии merge/rebase нет