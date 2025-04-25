**Отчет по оптимизации**
1. Чтение всего файла сразу неэффективно по памяти
2. 
![CopyQ qSQcER](https://github.com/user-attachments/assets/341fb37e-57fa-4815-8cf9-8a6d977b5c10)

 Заменяем на построчную обработку файла: io.ReadAll() -> bufio.Scanner...

![image](https://github.com/user-attachments/assets/a8c1687e-0028-4318-a0f0-13fd5848e676)

3. Для проверки браузера регулярные выражения компилировались каждую итерацию цикла
4. 
![image](https://github.com/user-attachments/assets/7fcc3564-0b92-494e-94ae-fd5ac0bdfb14)

  Заранее компилируем две регулярки для проверки браузера:
  androidR := regexp.MustCompile(`Android`)
	msieR := regexp.MustCompile(`MSIE`)
 
![image](https://github.com/user-attachments/assets/df29893c-1963-468c-82b9-e7f79d21439b)



