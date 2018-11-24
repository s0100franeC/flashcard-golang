# flashcard-golang
Student projects

struktura projektu
C:.
│   main.exe
│   main.go
│
└───form
        Courses.tmpl
        Edit.tmpl
        Footer.tmpl
        Header.tmpl
        Index.tmpl
        Learn.tmpl
        Menu.tmpl
        New.tmpl
        Show.tmpl
  
OPIS:
Program w założeniu miał służyć do tworzenia własnych kursów np. językowych, gdzie użytkownik może wprowadzić własne frazy i ich tłumaczenia. Efekt finalny jest dla mnie zadowalający w 60% (brak systemu powtórkowego weryfikującego poprawność zapamiętanych tłumaczeń, nieintuicyjne tworzenie kursów, brak usera).

W aktualnej formie aplikacja umożliwia:
- łączenie się z wcześniej utworzoną bazą danych (nie zaimplementowałem tworzenia bazy od zera)
- obsługa w interfejsie webowym napisanym w HTML/CSS

- dodawania nowych słów do bazy danych MySQL
- szczegóły podgląd wybranych wierszy w bazie
- edycja wybranych wierszy w bazie
- wybór kursów do nauki
- losowanie fraz w formie fiszek
