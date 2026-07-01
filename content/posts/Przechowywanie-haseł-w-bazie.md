---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD7DAYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIp+UAP/10O1jKf4s3NwJLPge4gy82c
  YfHmy9m28VvwYPhcNhJ9SlGuxgggC1CEARk2g134SzKHTg6RH9MMiW0A9dsSwsfV
  31P7rnhjdIC8tfyG0emqmTvHuX95rN5nwNoTfk9CB6t64eA/SK8f5ORtUyAERFkj
  qSUqCF2hSr2yLFewQZ8juzu5z/xH8n19TXvTom1O56esuopw+DjYCHCCnRBAWPAi
  19Xe/h0pkwZ+IuMbztklBUjr+oylcWSPe3WtjpC1ootfhh96GQTEo65HbskX+u/W
  DecYyf5ZgfSg7riCxNuP7muSGivFTpIJU+vWX1GPKDhN96TzD1lq4XMpGopXoh79
  MVSg6EUlOyLbM3WWu737eNBvyFZrfCNl5y6FJEBi7I/AEaPJskRw7ym0jQTxe5fK
  /y83WC0VrFfNwCWG99kVqG3zs3b6K8BNlulQZeOv7ud++2oIKBvXCkGb7NobR/Hj
  W9+nR/P1nMARZCROJ90EBeNxqWCkhKdHhEUoULIrg8nGOe/anhf/96rZptvyQ+eR
  lt/jFuimsyfdiACgD6emIvvP8eP0GxeAn8eAhMJloVnKC72azyZxUMDtr8+d+324
  u/63qEbJHKNIDxBPqFmXaS0xneq1F2ynvmQ2nBxkjL7PIuH/l9FTCzP0l9VMlEGn
  OBHU+pMXJEmsgQM/9p8a
  =yrEb
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
