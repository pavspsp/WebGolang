**Отчет по оптимизации**
1. Чтение всего файла сразу неэффективно по памяти
   
![CopyQ qSQcER](https://github.com/user-attachments/assets/341fb37e-57fa-4815-8cf9-8a6d977b5c10)

 Заменяем на построчную обработку файла: io.ReadAll() -> bufio.Scanner...

![image](https://github.com/user-attachments/assets/a8c1687e-0028-4318-a0f0-13fd5848e676)

2. Для проверки браузера регулярные выражения компилировались каждую итерацию цикла
   
![image](https://github.com/user-attachments/assets/7fcc3564-0b92-494e-94ae-fd5ac0bdfb14)

  Заранее компилируем две регулярки для проверки браузера:
  androidR := regexp.MustCompile(`Android`)
	msieR := regexp.MustCompile(`MSIE`)
 
![image](https://github.com/user-attachments/assets/df29893c-1963-468c-82b9-e7f79d21439b)

3. Стандартный парсинг encode/json неэффективен,

![image](https://github.com/user-attachments/assets/701d216b-169a-4345-9fec-c6a13b415478)


   Заменяем на сгенрерированный easyjson код

![image](https://github.com/user-attachments/assets/b8b49e3b-9be1-4bad-8b03-1cca8ea43503)


**Текущее решение:**
BenchmarkFast-8     501 2183829 ns/op 583289 B/op 7740  allocs/op

Образец:
BenchmarkSolution-8 500 2782432 ns/op 559910 B/op 10422 allocs/op







