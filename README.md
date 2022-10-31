#testProject

Для збору і запуску проекту потрібно

    docker compose build
    docker compose up

Щоб завантажити *csv файл в бд потрібно передати файл в параметри form body URL з методом PUT


http://localhost:4000/uploadfile 

Фільтрація контенту метод (GET):

    пошук по TransactionId відбувається по url 
    http://localhost:4000/search?transactionid=<номер>
    пошук по TerminalId відбувається по url 
    http://localhost:4000/search?terminalid=<номер>,<номер>,<номер>... 
    //можна вказати скільки завгодно змінних для TerminalId через ','
    пошук по Status відбувається по url 
    http://localhost:4000/search?status=<текст>
    пошук по PaymentType відбувається по url 
    http://localhost:4000/paymenttype?status=<текст>
    пошук по DatePost відбувається по url 
    http://localhost:4000/search?datepost=<2022-08-30>,<2022-09-28> // формат: рік-місяць-день
    пошук по PaymentNarrative відбувається по url 
    http://localhost:4000/search?paymentnarrative=<текст>

Запросити всю таблицю разом


    http://localhost:4000/search


Вказати декілька параметрів для пошуку одночасно, наприклад знайти стрічку де Status = "declined" і PaymentType = "card" 


    http://localhost:4000/search?status=declined&paymenttype=card


SWAGGER url: 


    http://localhost:4000/documentation/

Додатково можна запросити результати пошуку у csv файлі по url


    http://localhost:4000/searchcsv?<параметр>=<значення>&<параметр>=<значення>&<параметр>=<значення> ... 

Всю таблицю разом


    http://localhost:4000/searchcsv