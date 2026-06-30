---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2RoYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpf6EP/imL0vaLzPBLeGQqz7CVfXH+
  8AscbIZPB3yu4dHrQXD+2TcrUqGbSCeNhg7cxsVOf2UBNj3PrdmtmfjsBGj2O0xS
  9PMHctnzj4rYDnXa8sjzFZXbXYfY1SJ303UQeDB8KKX/m/g8Vc+3BoKS+S6auVwM
  CltLDCaZO4Ct1UHpBQDPlN2SEkETm99KhKaawZE/jIvuKceC0tg/IQWOhMMTWMJo
  p1XRqbTq5pYYQU583LO8s/oPNl3jYxPDkJ4nVkTuujA4f7i+rRgM4c3jv6VtI7+6
  53W4a3zjHdjxj3K48ayOk1JslbM6yVYWXhD86DZdGzvscFxw6k24E2AjkqePDlbZ
  0+JhPqZfU3kJl+Si994uVOyCV+zj6Q7jjWsoq3ZdvVhlypit9KewvGBc9IyPqGwX
  Pl51xxufePh1kg7ABbO8+a53AEtF4JOhambTyvADV6mD78Uq/DmYhg8Yl6+d8/yN
  LLpak8xOcimFoasxMDBag2UeisTuanp+Ta7BDtsJ9YTmBGgl/5aq1Xx0qhKytXK5
  U1TXjorSfAn6dFVSeR4Un+TUnvzgMdKd3UXvUZ4hfdvDwmXJkeoolzLZmwESVX0m
  5slOa9UfTNf5r/GptN+RTWkuLLWjv5sjBcX/xtkvrb18SUfaCSreMSdyy17Kh8hX
  /VL0pd+Wa/fo9k5RMmTb
  =nbAj
  -----END PGP SIGNATURE-----
tags:
  - PHP
  - Symfony
  - Hasło
  - argon2
title: Przechowywanie haseł w bazie
toc: false
---

Od kilku lat obserwuje ciągłe wycieki haseł z baz danych różnych podmiotów, w tym tych największych. Nie ma w tym niczego dziwnego, bowiem złamanie zabezpieczeń w IT to tylko kwestia czasu i pieniędzy. 
Nie mniej jednak ze zmartwieniem i niekrytym smutkiem przyjmuję, że większość haseł przechowywanych obecnie w bazach danych nie jest dostatecznie zabezpieczona. O fakcie tym świadczy chociażby [eksperyment studentów z Uniwersytetu Bonn](https://net.cs.uni-bonn.de/fileadmin/user_upload/naiakshi/Naiakshina_Password_Study.pdf),

w którym zlecili freelancerom stworzenie prostych aplikacji, które min. miały przechowywać hasła w bazie danych. Wyniki eksperymentu nie pozostawiają złudzeń, większość programistów nie dba o bezpieczeństwo, a co za tym idzie świadomie naraża zleceniobiorców na zagrożenie. 

W dzisiejszym wpisie przedstawię sposób na szybka implementację bezpiecznego sposobu przechowywania haseł. Teoretycznie zastosowanie funkcji takich jak SHA powinno być bezpieczne, jednakże mają one swego rodzaju ułomność. Nowoczesne procesory, w tym specjalne klastry GPU są w stanie wygenerować miliardy hashy na sekundę, co niejako zwiększa szansę na to aby wygenerować wartość odpowiadającą hash hasła w bazie.

W ciągu kilku ostatnich lat rozwijałem metodę haszowania haseł, starając się aby sposób ich łamania był możliwie najtrudniejszy. W tym celu używam kombinacji kilku algorytmów.

Cały mechanizm opieram o argon2 mieszany z globalnym kluczem dodawanym do hasła.

Na początku hasło jest przekształcane za pomocą SHA512. Stosując SHA512, możemy szybko konwertować naprawdę długie hasła na stałą wartość 512 bitów, co pozwala na działanie algorytmu z tą samą prędkością dla każdej długości hasła. Następnie tak wygenerowany skrót przekazany jest wraz z solą, która jest unikalna dla każdego użytkownika, do algorytmu argon2. Ostateczny uzyskany skrót jest szyfrowany za pomocą AES256 przy użyciu klucza globalnego (wspólnego dla wszystkich), który przechowujemy w aplikacji.

Całość mechanizmy wygląda dość prosto:
> Hasło → SHA512 → argon2(unique salt) → AES256(global salt)

Oczywiście istnieje ryzyko, że podczas włamania na serwer utracimy klucz globalny, jak również unikalne klucze dla każdego użytkownika, co może skompromitować działanie mechanizmu, nie mniej jednak odtworzenie tak zaszyfrowanych haseł może być wyjątkowo trudne.

Przykład: [PasswordEncoder.php](https://github.com/Effiana/password-bundle/blob/master/src/Security/PasswordEncoder.php)
