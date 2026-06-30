---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6wcYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpBH4QAJsVFZVXzyl4gx9gfZwwXRdz
  ZRlyQJVmwNKZrdaO9M/35Z24vov+SAmOFv6RFlA2xQzos/DjZib7Tt4SttRKXKL7
  3stxxR7nlOpyCSTI8QuzORmo0+oUIu5uVfw3xnKi7WW0kt0sZhRvdzfsDej8MHyr
  N2V9Hq5aLD06yi4LRdjXC78Oal/HpuCJmOZnVnbIGF9x7UOClpwQRITiI5T0OFZf
  Qxvh1GtG+yXJWAAWGA/Zq4UmcC5CqxeafFRzn1eEzMkBIK9tTjaM6DqG83ibxzA3
  530LUJ0JlcktzUOBHEv2niGI5xMeszTogy5OcDw/39Avw2y19aKDyekXyJrvbeAr
  2DbgTUuHMZpAExJIBUfQpL8oeCjIbppB/aarHd3aPk8fYNz4HMpMhemCH/A25wYo
  SY7QN9jelKcz3YRbz90hCo7sUwQL44lirLBo7tbCL56Yrcl8N20Lfj1SuAeAxv6z
  FqgUJ9lOprg7Z08o+JzTZv9zAWF7Xs29rafKMPCdRV1GSN1b38PvwmaAqrFghsd9
  GAGrvK2zyQ70jd1CHn36V9tb+jX2sahU6SP6EnkmOONRfZZtYxoNWzZsdrCG8mx9
  2CtFnRldACVs6K5XGDWdm/kFf/5w15ftezSsWPsSlcTlwyGGW1KKtC3xVPLBNV3O
  ytMTYO6hH/orqVWWEFU9
  =jsBA
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
