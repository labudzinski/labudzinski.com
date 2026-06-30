---
date: 2019-03-19T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJMBAABCgA2FiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpD6MUYHGRvbWluaWtA
  bGFidWR6aW5za2kuY29tAAoJEFzelGpBYMIpW+kQAKjRHQ8vmP/MSlM6FWdLlFul
  qZZg3kRstBvAt3w2oi7KLAbbTRGJ/tf9OlAPtDVgpBPlFABZjxYUAIr4zWc6E1gA
  Ds78fwUribmOMBNjVMge6ayvEEHD/PB1XIR5YqyRivFX7G+hk8kzXlcx37+xLk+x
  bVM0/CxDN9pqStzOhX+AwMxFxBrzPfCxPswtb8448ZFp6DsfVUvpJpLqyGX89cnG
  k7TlYmcyfLXXtgy4ZcNgIoWdC31TSD+bS2jO2ac0GWivIJb3dE31k2HrWvCP39zD
  MDVqAh1Iu/1zNJj77apCZRC8JV8FoDSHDsy7kbBOoIryRUReuALah47siscc3+JC
  mnqid4yrVFEEQKu4MJPVeFjPIImk9CN//FADLAZ33ZzHIdKNV182mBnQCTQ5n1/C
  ikiBbPl3Mg1jvfsMrnct6i34GGEGY7QpWlpO+olejKvRMgdb29UqzLVkZl/ud1sN
  j2Wc2HuEEi/oHfq5Q+0gxz5ztFwgGTg3J3cPMzXu/4mpETCyxUHUvk7QJiA0UY3V
  pCC6TUIH0Au6zePBDX9vM64URgyYxQtVwF718WdXcFYXf8S6sLucNFliWW7LkY3S
  9wkm+VNPkEsyHPibT42TRrGoR9SZu4O73rZpto+ZzhXLa8qIP2r2Ib6JcivHKtae
  a4VwXGuE1o0WQvw6noGb
  =cJQl
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
