# Что это 
Выполнение тестового задания https://github.com/avito-tech/qa-trainee-general

## Как запустить 
Необходимо на вход передать файл, в котором лежит токен и список каналов для рассылки  
**Пример вызова**
```
go run main.go -file ./example.json
```

**Пример формата файла**
```json
{
 "bot_token": "xoxb-2462481993190-2466240188101-pNUSuGFwZWP4GYDwQ25im4sz",
 "channels": [
   {
     "channel": "test1",
     "text": "Hello, world!"
   },
   {
     "channel": "test2",
     "text": "Hello, world?"
   },
   {
     "channel": "test3",
     "text": "Hello, world :)"
   }
 ]
}
```


