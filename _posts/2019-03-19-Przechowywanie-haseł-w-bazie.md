---
layout: post
title: Przechowywanie haseł w bazie
---

Od kilku lat obserwuje ciągłe wycieki haseł z baz danych różnych podmiotów, w tym tych największych. Nie ma w tym niczego dziwnego, bowiem złamanie zabezpieczeń w IT to tylko kwestia czasu i pieniędzy. 
Nie mniej jednak ze zmartwieniem i niekrytym smutkiem przyjmuję, że większość haseł przechowywanych obecnie w bazach danych nie jest dostatecznie zabezpieczona. O fakcie tym świadczy chociażby eksperyment studentów z Uniwersytetu Bonn ([https://net.cs.uni-bonn.de/fileadmin/user_upload/naiakshi/Naiakshina_Password_Study.pdf]), w którym zlecili freelancerom stworzenie prostych aplikacji, które min. miały przechowywać hasła w bazie danych. Wyniki eksperymentu nie pozostawiają złudzeń, większość programistów nie dba o bezpieczeństwo, a co za tym idzie świadomie naraża zleceniobiorców na zagrożenie. 

W dzisiejszym wpisie przedstawię sposób na szybka implementację bezpiecznego sposobu przechowywania haseł.

Niewątpliwie najlepszym algorytmem jest bcrypt, ma on jednak pewną wadę, która nie pozwala implementować go produkcyjnie. W przypadku użycia zbyt długiego hasła potrafi zawiercić aplikację używa, a w kompilacji z limitem długości wprowadzanego ciągu znaków pozwala na użycie tylko 72 bajtów.

Jednakże eliminowanie tego algorytmu, z powodu tak błahej ułomności byłoby nadużyciem, a w konsekwencji działaniem na szkodę klienta. Nie mniej jednak, aby bcrypt nie sprawiał nam problemów trzeba go przekonać, że nawet przy długich hasłach, zawsze będzie dokonywał prawidłowego hashowania. W tym celu użyjemy funkcji skrótu SHA-2 (SHA-512).

> Hasło od użytkownika → SHA-512 → bcrypt(10)

Aby lepiej zabezpieczyć hasło, dodajemy do powyższego schematu unikalna sól dla każdego rekordu.

Cały proces będzie wyglądał następująco:
> Hasło od użytkownika → SHA-512 → bcrypt(per user salt, strength 10)

Dla większości systemów będzie to wystarczające zabezpieczenie, nie mniej jednak śledząc trendy i rozwój technologii nie jest to rozwiązanie kompletne. Na samym końcu powinniśmy zastosować algorytm szyfrujący. Dlaczego szyfrujący, a nie algorytm skrótu? Każdorazowe wykonanie algorytmu bcrypt zwraca zupełnie inny ciąg znaków. Zastosowanie na wynikowym ciągu znaków funkcji skrótu spowodowałoby niemożność porównania wprowadzanego hasła z tym z bazy, co w efekcie eliminowałoby skuteczność takiego mechanizmu.

W tym wypadku użyjemy algorytmu AES, o długości klucza 256 bitów, który coraz częściej wspierany jest przez procesory.
Nasz proces dostanie kolejny etap, a co za tym idzie kolejna barierę, którą musi pokonać cracker, aby złamać nasze hasła. Do szyfrowanie dodamy globalną sól, która powinna być zapisana w kodzie aplikacji, konfiguracji, a najlepiej w kluczu USB.

> Hasło od użytkownika → SHA-512 → bcrypt(per user salt, strength 10) → AES265(global salt) → base64

Tak zapisane hasło jest na tyle bezpieczne, że jestem skłonny pokusić się o załączenie do tego wpisu hasła, którego używam na co dzień do prywatnej poczty: 

> jeUOmFUp2G3/qBndDwlG2k75qdBxQhFeFsjJS6LZm0Yt9OIFsO9oPFkmRGjV2enPCOE3VxYXEA726fr+vMb2o+pW/P9HoiJR3EuFkCp5zfQCz1c5k9AjikdtbiQbOSrgOL2Z0qNj2FXrJdHyJubAmSU=

== TODO implementacja argon2, przykładowy kod w PHP/Java