---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzusYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpRaMQAJpoMjsiSCio4ulZPedd15ZA
  FU8TnTHQU22Uv5v2Plv2Yk33cec9DDNunr9kqG4oxWI2oLkZ2q6Uz3IL2MIINtEc
  yruXMrxPwHlOv8Iu4zHdcHzrv7wVdEBALgpXSeyw3B7A/mlZsBArU0mCdCag3jnK
  hPnr6Z/j/VZdpvd+Em1xrTU5ycuaPzk4RCaencAbbeuFzIyO2uIs9k18DEeJVH9o
  48sGmwwbckHAY2f5jESsiTTELu7kFzDEz4KzZc0f/Kk2E6GZ9bhVk3cOIOk9qtIB
  SZJ/FmyIW54zm0Dy9numPjPsJNaPXBts8qOGbrokgI/++nZMhaerw65+wzbmkoK4
  V8AUiam4VmMz6scRXbMhld/IKTYiRrotqXPuRsbPjF9a8o6gnFQtKreLfFMAF4lf
  B8Csu3DFCGk8wdBgrG97AGTKHmv8GnhbsCZlKBfVnGfOfj5w9TCXoes96VdybhdM
  x8uBcwwAbmZVBOb19SIuZ4/hE7X/9lMtpb2Qo3lvXTlVT4EmRI6N0z6kx1LSvUdP
  LIGdSHnt/9eZMgFC6hi8lNBMExtNliL/CAJo19q9eUp8+zL/lWtqf/w2gI64eFH0
  iDyX59vNm5ebKowA5K12OJcx4GT5CDE8mA/3M/+qVJk1oSouTMjo/4R3UP7aLD9H
  bOBJe69dpb6wqIsvddpB
  =EsSU
  -----END PGP SIGNATURE-----
tags:
  - PHP
  - Symfony
  - Hasło
  - argon2
title: Przechowywanie haseł w bazie
toc: false
slug: przechowywanie-hasel-w-bazie
aliases:
  - /posts/przechowywanie-haseł-w-bazie/
description: "Wycieki haseł to kwestia czasu. Problem w tym, że większość baz nadal trzyma je tak, jakby tego czasu miało nie być."
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
