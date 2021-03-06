# Запуск программы показывает, что мы выполнили
# около 90000 операций в нашем синхронизированном
# с `мьютексом`состоянии.
$ go run mutexes.go
readOps: 83285
writeOps: 8320
state: map[1:97 4:53 0:33 2:15 3:2]

# Далее мы рассмотрим реализацию той же задачи управления
# состоянием с использованием только горутин и каналов.
