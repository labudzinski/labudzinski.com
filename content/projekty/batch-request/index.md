+++
title = "BatchRequest"
description = "Biblioteka PHP do wielu wywołań API w jednym żądaniu HTTP dla Symfony i Laravel. Operacje równoległe, scalona odpowiedź JSON, obsługa błędów RFC 7807."
date = 2026-07-08T00:00:00+02:00
draft = false
github = "https://github.com/Lemric/BatchRequest"
composer = "lemric/batch-request"
license = "MIT"
license_url = "https://github.com/Lemric/BatchRequest/blob/main/LICENSE"
logo = "/images/projects/batch-request.webp"
logo_alt = "Logo BatchRequest"
docs_url = "https://github.com/Lemric/BatchRequest#readme"
tags = ["PHP", "Symfony", "Laravel", "API"]
toc = false
+++

BatchRequest to biblioteka PHP, która pozwala wykonać wiele wywołań API w jednym żądaniu HTTP. Powstała z myślą o aplikacjach Symfony i Laravel, w których liczba połączeń po stronie klienta wpływa na wydajność i przewidywalność działania.

Zamiast wysyłać serię osobnych żądań, klient przekazuje jedną partię operacji na punkt końcowy, zwykle metodą POST pod adresem /batch. Każda operacja w partii, zapisana w JSON-ie, zawiera metodę HTTP, adres względny i opcjonalne ciało. Operacje niezależne wykonują się równolegle, a zależne po kolei. Serwer zwraca tablicę odpowiedzi w tej samej kolejności, dzięki czemu łatwo przypisać wynik do konkretnego wywołania i zdecydować, które nieudane operacje ponowić.

Biblioteka współpracuje z symfony/rate-limiter i liczy każde wywołanie w partii jako osobne żądanie przy limitach API. Błąd jednej operacji nie zatrzymuje pozostałych. Nieudane odpowiedzi podrzędne mają format RFC 7807 (typ application/problem+json). Obsługiwane są różne typy treści: JSON, HTML, dane binarne zakodowane w base64 oraz odpowiedzi bez ciała. W jednym wywołaniu można też przesłać załączniki binarne w formacie multipart.

W Symfony integracja z profilerem pokazuje liczbę żądań podrzędnych, czas wykonania, zużycie pamięci i szczegóły transakcji. Pakiet działa z Symfony w wersjach od 6.4 do 8.x. W Laravel wystarczy zarejestrować dostawcę usług i ustawić limit wielkości partii w konfiguracji.

Biblioteka jest udostępniana na licencji MIT. Instalacja: `composer require lemric/batch-request`. Kod, dokumentację i przykłady znajdziesz w repozytorium na GitHubie.
