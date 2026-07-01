---
date: 2022-06-11T09:45:32+02:00
draft: false
images: null
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQIzBAABCAAdFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmpEzlMACgkQXN6UakFg
  wimuRhAAop26S/QSaT280ptt1Xm6B72DJwRSbpaNvHNj/GwkpegunqZRED9VKFcI
  b8zRIZ3kcaNzbUsi3dHBP7TLHIO9HuvWY2+FS5aqaKAGHBAJeYLit38r14U3s+cZ
  xv9AsPHD46Gok8FXPZrjCFmdFK4uJSRzEPFDZAn9euUQubj8yWim6GjSXcVUuAG1
  Vn6Hz9UaeQFsWzMFTbezBMPMkgjmUvhsWyNkcdBOM24rn9KDFKTwt2cJImU62hwT
  ZDQ+EfcJUB7HTwV7d/BIRkAsRpV36bhCRxxmXwlCyoRC85JwJKBxdrQEteZOV1hn
  k1UDgjrGzkbz5l9PJNBk4MrHxoXI0v1RnFbbppx75PDcQjWMBdVXmvgcntVSVpsR
  707y6oGZQXFiBMjIV4ed7izb1Y+QLm+SqfsU7fEnlkUoLY7lVCwT7WQTFdJP1k8c
  h3k1MT64jevfN3iRuw45Zrq9RI5UOPfn6bbMkAFBeRuHJTWRMWBvHP1fGrDMwLyg
  +LWmUfZBIB8Y2rikeQYldtdvQoOYMR6yjn7S24cq1YvXFxBCC4grikUp08eUh7sc
  /FvpFy8BW/3DP47krkZ9KR+Jwr0bfRGHpdq8MB9Q8OpAp6bx/OEezsJ5F/uE35/2
  RnQwuVg5ArcTd4oDvP9AlT/E2g70uMCZiI7meb5cdBo6UVm1QSI=
  =5AGa
  -----END PGP SIGNATURE-----
tags:
  - BrandOriented
  - szyfrowanie
  - rsa
  - pliki
title: Szyfrowanie plików w trakcie przesyłania do aplikacji Effiana DAM+
toc: false
---

Dzisiaj skupimy się na sposobach przesyłania i przechowywania wrażliwych plików. Dlaczego? Każdego dnia BrandOriented, wdrażając system Effiana, ułatwia pracę zespołom komercyjnym w ich codziennych obowiązkach, m.in. wspierając decyzje, integrując dane w środowisku rozproszonym, jak i wspierając procesy wewnątrz organizacji. Effiana w odczuwalny sposób zmienia niemal każdy aspekt pracy zespołów oraz ich wyniki.
Istnieje jednak wiele mechanizmów, które nie są widoczne dla użytkowników, a które odpowiadają za niezwykle ważny obszar — bezpieczeństwo. Tak jest m.in. z szyfrowaniem plików, które omówimy poniżej.
<!--more--> 
Zanim zaczniemy przygotowywać się do szyfrowania plików, musimy zaplanować cały proces, a możliwości jest wyjątkowo dużo, zarówno jeżeli chodzi o algorytmy szyfrowania, jak i metody przetwarzania zaszyfrowanych plików. Na potrzeby tego materiału skupimy się na metodzie szyfrowania podobnej do tej, którą stosujemy w systemie Effiana, natomiast pokażemy ją oczywiście w dużym uproszczeniu. Przykłady mają na celu przybliżenie sposobu działania mechanizmu, a nie rzeczywistego kodu. Metoda ta zakłada, że pliki trafiają do serwera już w formie zaszyfrowanej, a dostęp do nich jest ograniczony do wybranych osób i nie wymaga ponownego szyfrowania na skutek zmiany odbiorców. Przykłady będą stanowiły pewnego rodzaju ideę, której sposoby implementacji mogą się znacznie różnić od siebie.
Zacznijmy zatem od prostego kodu wysyłającego pliki do serwera z przeglądarki użytkownika:

```
function upload() {
    const file = document.getElementById("file").files[0];
    fetch('/upload', {
        method: 'PUT',
        body: file
    });
}
```

Powyższy kod wysyła plik do serwera w najprostszy możliwy sposób bez szyfrowania. Naszym celem jest jednak przesyłanie zaszyfrowanego pliku, który nigdy nie trafi do serwera w postaci jawnej. Dodajmy zatem do funkcji zapisującej plik mechanizm szyfrowania, w tym celu użyjmy biblioteki cryptojs:

```
function upload() {
    const file = document.getElementById("file").files[0];
    saveFile(file);
}
async function saveFile (file) {
  const reader = new FileReader();
  reader.addEventListener("load", () => {
        fetch('/upload', {
            method: 'PUT',
            body: CryptoJS.AES.encrypt(reader.result, "<pass>")
        });
    }, false);
    if (file) {
       reader.readAsText(file);
    }
}
```

W takiej formie funkcja saveFile wyśle plik zaszyfrowany bezpośrednio do serwera. Powyższy przykład zawiera jednak poważny błąd. Kluczem szyfrującym jest zapisane w kodzie hasło. Oczywiście można użyć w tym miejscu różnych mechanizmów, tj. sesje, czy cookie, natomiast w dalszym ciągu będzie to mało bezpieczne, gdyż hasło, którym zaszyfrujemy nasz plik, będzie w dalszym ciągu jedynym zabezpieczeniem. W tym celu zastosujemy szyfrowanie asymetryczne RSA i użyjemy biblioteki cryptico, a na potrzeby tego artykułu zapiszemy klucz publiczny w kodzie — docelowo powinien on być zwracany z serwera, ale na tym etapie przyjmijmy, że tak jest.

```
const publicKeyString = 'MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGFxQyrqquJh+W/5wtrD3lrEWaCp
EDjcHreBGU37v58GB2i3Vl/3VBP1KUqlVmkre/iAXSOKKKnS2k8cqftPZMGf8+r6
BJKX7PJkKIEDkxBjReoaIZsibqGEKYF1ZLKmrvXlBO2+EyMvymv8A8QnJtHWdznw
yitAM25LKNRnPVWbAgMBAAE=';
function upload() {
    const file = document.getElementById("file").files[0];
    saveFile(file);
}
async function saveFile (file) {
  const reader = new FileReader();
  reader.addEventListener("load", () => {
        fetch('/upload', {
            method: 'PUT',
            body: cryptico.encrypt(
                   reader.result, publicKeyString
               ).cipher
        });
    }, false);
    if (file) {
       reader.readAsText(file);
    }
}
```

Podsumowując, mamy już mechanizm, który szyfruje nam plik, zanim trafi on do serwera, używamy w tym celu szyfrowania asymetrycznego RSA, który do szyfrowania używa klucza publicznego, a do odszyfrowania klucza prywatnego. Pora zatem przejść na drugą stronę i przyjrzeć się jak po stronie serwera wyglądać będzie cały mechanizm.
Na początku musimy wygenerować parę kluczy i tu pojawia się ważne pytanie:
Czy każdy przesyłany plik powinien być szyfrowany tym samym kluczem?
Odpowiedź jest tylko teoretycznie prosta, gdyż w zależności od tego, jaka zostanie udzielona zależy to, czy wymagać będzie dodatkowego nakładu pracy i zasobów. Zespoły developerskie często starają się tworzyć rozwiązania jak najmniejszym nakładem pracy, co w przypadku tematyki bezpieczeństwa rzadko jest dobrym pomysłem. Wielokrotnie w ostatnich latach, a nawet miesiącach, słyszymy o wyciekach haseł, które po weryfikacji okazują się słabo albo wcale zabezpieczone.
Moje podejście do tematyki bezpieczeństwa zawsze zakłada maksymalizację wysiłków i środków, aby osiągnąć zadowalający rezultat, tym bardziej w kwestii tak newralgicznej, jak przechowywanie plików, które nierzadko mogą być dokumentami poufnymi.
Dlatego w dalszej części skupimy się na maksymalizacji efektu szyfrowania, a więc dla każdego pliku będzie generowana osobna para kluczy.
Na początek poprawimy kod javascript, który będzie pobierał przed każdym wysłaniem pliku na serwer nowy klucz prywatny i identyfikator pliku, który pomoże przypisać odpowiednie klucze.

```
function upload() {
    const file = document.getElementById("file").files[0];
    fetch('/upload', {method: 'HEAD'}).then(function(response) {
      saveFile(
          file,
          response.headers.get("File-Public-Key"),
          response.headers.get("File-Id")
      );
    });
}
async function saveFile (file, publicKeyString, fileId) {
  const reader = new FileReader();
  reader.addEventListener("load", () => {
        fetch('/upload', {
            method: 'PUT',
            body: cryptico.encrypt(
                   reader.result, publicKeyString
               ).cipher
        });
    }, false);
    if (file) {
       reader.readAsText(file);
    }
}
```

Metoda “/upload” [HEAD] powinna zwrócić w nagłówku oba potrzebne elementy które, zanim prześlemy zostaną zapisane po stronie serwera aplikacji np. w bazie danych. W tej sytuacji metoda “/upload” [HEAD] generuje parę kluczy i zapisuje ją w bazie. Gdy chcemy, aby plik po zaszyfrowaniu nie mógł zostać podmieniony i ponownie zaszyfrowany, możemy nie zapisywać kompletnej pary kluczy i pozostawić jedynie klucz prywatny, który będzie służył do odszyfrowania pliku.
Na potrzeby artykuły przyjmujemy, że przechowywane są tylko klucze prywatne, które są szyfrowane dodatkowo za pomocą algorytmu AES, dla którego klucz jest zapisany w pamięci serwera (plik, zapis w konfiguracji itp.). Ma to zabezpieczyć sam klucz prywatny przed dostępem osób nieuprawnionych, które mogą operować na bazie danych. W sytuacji, gdyby uzyskali dostęp do bazy danych, nie będą w stanie odszyfrować plików.

```
File-Public-Key: MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgGFxQyrqquJh+W/5wtrD3lrEWaCp
EDjcHreBGU37v58GB2i3Vl/3VBP1KUqlVmkre/iAXSOKKKnS2k8cqftPZMGf8+r6
BJKX7PJkKIEDkxBjReoaIZsibqGEKYF1ZLKmrvXlBO2+EyMvymv8A8QnJtHWdznw
yitAM25LKNRnPVWbAgMBAAE=
File-Id: 70ec0290-2768-42f4-af90-e378c524264d
```

```
|------------|--------------------|--------------------|
| Id         | File_ID            | File_Key           |
|------------|--------------------|--------------------| 
```

Po wprowadzeniu tych zmian w kodzie na serwerze powinien znajdować się już zaszyfrowany plik oraz klucz prywatny, służący do odszyfrowania w formie zaszyfrowanej zapisany w bazie aplikacji. 
To jednak nie koniec samego procesu szyfrowania, ponieważ brakuje tu jeszcze danych pozwalających na odszyfrowanie. Co istotne w tym momencie żaden użytkownik nie może odszyfrować pliku. Aby funkcjonalność była kompletna, po załadowaniu pliku powinniśmy nadać odpowiednie “uprawnienia” dla właściciela pliku (osoby, która wgrywa plik na serwer), czyli najogólniej mówiąc, musimy zaszyfrować dla niego klucz prywatny pliku. W tym celu do zaszyfrowanego pliku dodamy wpis z kluczem prywatnym, który zostanie zaszyfrowany indywidualnym kluczem właściciela.
Sam pomysł nie jest nowy, podobny mechanizm został użyty w PGP, natomiast zamiast ponownie szyfrować plik, do którego zostaną dodane klucze, do pliku zaszyfrowanego dodamy osobne wpisy z kluczami użytkowników. Pozwoli to uniknąć potrzeby odszyfrowywania i ponownego szyfrowania, zarówno w kontekście dużych plików, jak również w kontekście bezpieczeństwa, ponieważ plik nie będzie odszyfrowywany i zapisywany na serwerze. Nasz plik będzie wyglądał mniej więcej w taki sposób:

```
-----BEGIN EFFIANA KEY-----
ID:24f1b2296b37c229789611a2293ef498bdde0787c0db3ae5ba216ac51a02c
KEY:/Oc8eQ6SJ52zxlIceORE20ZBqP+u4DhmY5i8rCeTv3GHnktDQBKxfZhqp8oAJKX3hlOvBApt54hZY4FupsxTpyKVbonleDETZcjNENs3Tk4GQqKOd9swADEWyQWC1DTOE+h20pEWldHCp9PKPw5wvg==
-----END EFFIANA KEY——
-----BEGIN EFFIANA MESSAGE-----
uAc6JIvkisAwvHXKjPuYmldoVC1sQWLY7J/Z0Ke1Dqy3yPHeMKzjyzzp8WSukWFCMG54A+9J/gdqDsA7lsScCdooSBNm7lEYuRLiopdzPEt72rqIQqKw2PxTvwCJWgfrDn3DZlFNgCs664OuF3S3mWK6cRoq9lybhCwsPCUGQPwDFHARXxMoEqucbgw5x30+UCNMzEODolWfJv0EgiOb4R55MydPOBVpUT9zRDJW95+mK4xQmjYYtr4kog1C436PdWzpMkNOtzOD49QhHS2y06jONZyD2tG4VMqjD+C9qSqK6tX/p/ZRRzNK9B2yqUALGhnH1wuSKrVNCGIIjXdk/0lCkoSCBEF+J0Nnx0erdjpqiU2AkbDtuRZNKliNIPGhj8p+61Mq4v8OO52PoLwa/rx+RvVlv8k5yJv6p1nHIDB1rnzAKdHz9sKIIY7UYeiX6Z7pGjQk5JEgQRWFnoGAoHAszgzhZe88SPAr+KDNadzxPk+U3+wuLng6w4Wn0ZxmQ=
-----END EFFIANA MESSAGE-----
```

Teraz plik jest kompletny, zawiera wszystkie potrzebne elementy, w tym identyfikator i klucz użytkownika, którymi będziemy się posługiwać przy odszyfrowywaniu pliku. Identyfikatorem może być adres email, unikalny numer użytkownika itp. W systemie Effiana używamy identyfikatora użytkownika, który przechowywany jest w postaci UUID, a dodatkowo jest on zapisany w postaci funkcji skrótu za pomocą algorytmu Keccak (SHA3).
Jak zatem wygląda proces tworzenia “EFFIANA KEY”?
Na samym początku generujemy parę kluczy dla użytkownika i pliku, gdyż na wstępie założyliśmy, że każdy plik będzie szyfrowany unikalnymi kluczami. Następnie dla zaszyfrowanego pliku pobieramy klucz prywatny, dokładnie ten sam, którego używamy do odszyfrowania. Teraz pozostało tylko odszyfrowanie klucza prywatnego za pomocą AES i zaszyfrowanie go za pomocą klucza publicznego wygenerowanego dla użytkownika. Dla bezpieczeństwa zapominamy o kluczu publicznym, a w bazie zapisujemy tylko klucz prywatny.
W tym miejscu możemy zastosować różne techniki zabezpieczenia tych kluczy, w tym np. dodatkowe szyfrowanie np. AES za pomocą przypisanego do użytkownika unikalnego tokenu. Możliwości jest wiele, natomiast najważniejsze, aby żadne klucze nie znalazły się w tym samym miejscu (serwerze, dysku itp.), w którym znajdują się zaszyfrowane nimi pliki.

> UWAGA: pliki zaszyfrowane w ten sposób nie będą mogły zostać odszyfrowane bez odpowiedniego klucza, a w przypadku skasowania ich z bazy, pliki staną się bezużyteczne i nie będzie możliwości ich odzyskania.
> Aby temu zapobiec, można dla każdego pliku szyfrować klucz publiczny i wysyłać go do zewnętrznego kontenera. W systemie Effiana stosujemy w tym celu dedykowaną bazę danych, która umożliwia tylko i wyłącznie dodanie do tabeli kluczy, ale nie pozwala na ich usuwanie, edycję i pobieranie, dodatkowo co określony czas wszystkie klucze są zapisywane w zaszyfrowanej kopii bezpieczeństwa, która jest rozsyłana do kilku miejsc docelowych.
> W przypadku bazy danych komenda może wyglądać tak:
> MySQL: "GRANT INSERT ON table TO ‘username’@'localhost’;"
> PostgreSQL: "GRANT INSERT ON table TO username;"


W następnych częściach zajmiemy się szyfrowaniem dużych plików (>1GB), pobieraniem odszyfrowanych plików i rozszerzaniem osób uprawnionych do ich odszyfrowania.
Pierwotnie opublikowano na [LinkedIn](https://www.linkedin.com/posts/brandoriented_security-encryption-activity-6940972126568910849-P1GX/)
