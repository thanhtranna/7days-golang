# 7 days golang programs from scratch

What can be accomplished in 7 days? A gin-like web framework? A distributed cache like groupcache? Or a simple Python interpreter? Hope this repo can give you the answer.

## Web Framework - Gee

Gee is a [gin](https://github.com/gin-gonic/gin)-like framework


- Day 1 - http.Handler Interface Basic [Code](gee-web/day1-http-base)
- Day 2 - Design a Flexiable Context [Code](gee-web/day2-context)
- Day 3 - Router with Trie-Tree Algorithm [Code](gee-web/day3-router)
- Day 4 - Group Control [Code](gee-web/day4-group)
- Day 5 - Middleware Mechanism [Code](gee-web/day5-middleware)
- Day 6 - Embeded Template Support [Code](gee-web/day6-template)
- Day 7 - Panic Recover & Make it Robust [Code](gee-web/day7-panic-recover)

## Distributed Cache - GeeCache

GeeCache is a [groupcache](https://github.com/golang/groupcache)-like distributed cache

- Day 1 - LRU (Least Recently Used) Caching Strategy [Code](gee-cache/day1-lru)
- Day 2 - Single Machine Concurrent Cache [Code](gee-cache/day2-single-node)
- Day 3 - Launch a HTTP Server [Code](gee-cache/day3-http-server)
- Day 4 - Consistent Hash Algorithm [Code](gee-cache/day4-consistent-hash)
- Day 5 - Communication between Distributed Nodes [Code](gee-cache/day5-multi-nodes)
- Day 6 - Cache Breakdown & Single Flight  | [Code](gee-cache/day6-single-flight)
- Day 7 - Use Protobuf as RPC Data Exchange Type | [Code](gee-cache/day7-proto-buf)

## ORM framework - GeeORM

GeeORM is an imitation of [gorm](https://github.com/jinzhu/gorm) and [xorm](https://github.com/go-xorm/xorm) ORM framework

- Day 1：database/sql basics | [Code](gee-orm/day1-database-sql)
- Day 2：Object table structure mapping | [Code](gee-orm/day2-reflect-schema)
- Day 3：Record addition and query | [Code](gee-orm/day3-save-query)
- Day 4：Chain operations and update and delete | [Code](gee-orm/day4-chain-operation)
- Day 5：Implement Hooks | [Code](gee-orm/day5-hooks)
- Day 6：Support transaction (Transaction) | [Code](gee-orm/day6-transaction)
- Day 7：Database migration (Migrate) | [Code](gee-orm/day7-migrate)