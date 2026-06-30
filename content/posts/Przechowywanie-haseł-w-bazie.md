---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6eIYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpWWoP/j8lSFd1CJfPZ5MtK/6Qf/6N
  n+2zLBk38sUOzD5qvApXo6Pd0zLn6jzAJkTxTNSpxnh4ImhxDfd4mVJ7PB20TT6s
  Mvdi4IqRTN/Gu9Ni7/fsSSKceMr5k/V8JNnzf1+6niVu9Jxkv8mZ9aVm35wQWo96
  QAvShgDpRnay0hiM5L54r5DsrqHri541wSTsziM1S6KqaivF9iepUKCHmwKc51Ii
  5rvOL/yX3UliIvuXzK50qbDk/pK9/VKptdx+DWpUBveq7ffXX6Cz2W8pIv9p90Cu
  1IGijmOZN2ldGdLL4V9LasWNAlbcM+bGtbrQuMGzB7iGn20m8tVoZI6m48SCnamn
  bcmQqOQTiuJAw51gAaJDfzBP9UsWZ3sZ3oPGPcjNlSOAwFQOusOQqpO0WtvjzEl9
  nQPMHuo9GGfcbIS5VohY+CjqACaf7VNacXq4OzHjgliuFBkdIADuciPt2+az3+0T
  z3U9kzCDB3CchPdgA2w5/fBUKZxA6n/nJuM2XxDvg/q5RAa8JnjIv/XlQTAw9w1A
  fccmKnBNohasgR1iDquRztNFtF02kD7uovCTc2B+Ge9N5GO3+aQaHJyDgfyh8aFR
  k2Gve4khFdsCyMOrueRZkVEPS9IqQ3KLiBsXefPqaMa65HG+RISRsVteN0CScMpw
  /WGQQgO/ygEv+dC4XGGQ
  =Aqi8
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
