---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD2CMYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpgvEP/3/fTBz/f+QUeyqJ8CyvHqPb
  XJC/pXAhBvcKgsdr0ti4bcZ2onnGDLfMi0+6FM2nMtmMK15Z022HSmfr6vFHcW0N
  BMpb4hIi1tmY/YFA39RLGZn53fV/zFR50+EZJJKRXgs8+/0H+QdVfsiJzTaauqjg
  YCZ7SEH8SGLUJ9dUB+D9w8hG1JSXqWjK+am9/Josksmc5h0k9rfKIGCnRAeMXrnE
  ngfooTfb9OMnSNefxpfBQFfczMu15EHtj4idOUPLCJjshCqb7svuapuhNZdHDpH4
  Cc6TBKn0sZJAj4f1xWe+G+wiI6E4EiugzdipkDRbJ6WH32bkizDV+J2Icttoy2K6
  /8I+u1JS44rGrBQ0+EZU47SpIW/B/bxcoKfju448Sbie6T1Z/WV0/5e7phpDSSND
  6Cee2xME01T7CBrcCeKym9k99Y4LnSkHMIh81qNuw3vwcS1Lk1avt6fUWwiYGRCC
  ZPu03ItEeTc5nSnM4wzuGNNBvgKD5hd+V0KyiBcn9g1J0Q33c1zez0m6UUWONg8u
  2rEkaIu766K7iAyk3owsyAcJoZvfeLVcKWBp7UqeeB4oMGld2k86JxN+ViXlRM/V
  vOGIxTI0ZPVCBC70fn2trZW0kgw9X/DtRGc8EpiUZQDnxbi3J98YfT5plQ0fSLFd
  ZNYLcTXvNU6QTiYRaya9
  =V1ur
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
