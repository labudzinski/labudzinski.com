---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzlMACgkQXN6UakFg
  wiki1Q//duL2Ei8XFxv4FGu9xEAew+CQlg0aeoVYbOB13kPgf0RwFmnBATMQTLkw
  5+c+L0fTilQGxx70UE+kXYjD1hRNi6U9kDgylFIELKFjqbxNE2BRamhU5WZa4hcC
  NVpTfKOmmV9qe639qMhRQcpHAK/kTfYhe5qwi88b98TTSqcpKJcPyX7er3Yi3DNC
  lnx7sMFrtDEaHQCAJ/zt4kAfAgzn8JZ6hld+cmcQFtgRstwlxBCJpAEIb3Qq511Y
  fsNAo2LaRREmlWvSYb2aDqf9MDxkAxPh0flWHREOXJ4+z9szQ3PGMI7Pmwlt2Ee8
  168QVslYEVwID+DMrgWVPk6vVGtPjvwupIm1tXp9+ct2WZ12b4hxk8Teax2NdCAY
  2ldxtYg1kGstT6z5iMjEHlRRBCjeGV92lVXfh/6cVwOWDjxnTXHRIgjhwEu2E3Dt
  QPlfiiZMZr5baYUA2v0PtoYVKUwntOWdeH83qphYnSraBto0AAkUWpxvyOyLfKZB
  IQqR7sT8SJjtGiEEzisy0OwKPRZt025m2mGWTbBYN2M1Z8bt2TV4bPmp7okAqQQH
  6fj+RFgE39VYpAiJwwrgvmYYnJXU6iE0FEw/3z7hLUrT3BBJ1uw+hyVn46kkmdPy
  8vVq/1cZmANMB5yOPrPiPfcxOD8xl51tgb6S1LetkvOI7IHl7Z0=
  =h7EN
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
