# Мы ожидаем получить ровно 50 000 операций. Если бы
# мы использовали неатомарный `ops++` для увеличения
# счетчика, мы бы, вероятно, получили другое число,
# изменяющееся между прогонами, потому что горутины
# мешали бы друг другу. Более того, мы получим сбои
# в гонке данных при работе с флагом -race.
$ go run atomic-counters.go
ops: 50000

# Далее мы рассмотрим мьютексы, еще один способ
# управления состоянием.
