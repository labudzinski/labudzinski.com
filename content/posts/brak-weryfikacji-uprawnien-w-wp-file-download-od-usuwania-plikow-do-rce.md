---
date: 2026-09-03T10:00:00+02:00
description: Wtyczka WP File Download posiada krytyczną podatność typu Arbitrary File Deletion. Użytkownik o uprawnieniach Subscriber może usunąć dowolny plik z serwera, w tym wp-config.php.
draft: false
pgp_key_fingerprint: 9670 B378 988E 18A5 B30E 75D7 5CDE 946A 4160 C229
pgp_signature: |
  -----BEGIN PGP SIGNATURE-----

  iQJoBAABCABSFiEElnCzeJiOGKWzDnXXXN6UakFgwikFAmqZiXwbFIAAAAAABAAO
  bWFudTIsMi41KzEuMTIsMCwzGBxkb21pbmlrQGxhYnVkemluc2tpLmNvbQAKCRBc
  3pRqQWDCKfT3D/9sWWloX+zdLjA6bCfzN07apn930AbyvlzR84BdzUVNsDTYY24N
  Yhizr6woYRAAsvmA6MDnrVurAskKrA/D0/ohwMhRJKG+gZuSfmAQ6fivP7MNKYLv
  RCfVGnG5bLYIAGs4rL7ZDUVLLsDHnMk54jFvf0UgqvHkuDR8zaPpPoxLHe0YmmTO
  BAmJmCJbVhUjK5tXFEBci2TsPM0dVUjNhqZizszAMEjcdUUNZkjlgtcJYhXrlh+g
  8pVw5CMuxxylWhvhiL4ChmXz49EMkBtReYZcNCbTu+saZ1jiTf9FM7aeLBQtyEBL
  dI87dv+vI9ORQ9GAfl35cFOpxML4bwZUXoX/IVOFdm8H4d3sslwDEfSuhb+zZnp1
  ThkbLaysDsFP70ot6fH0YGgNZuG3DxSN6YEH0XidDwJpZ87GS0rEwAId9bahecZ1
  KE/XfmkoWrc30tpZYHnwLVMuRbisSMYAXx8PQQGk0QO3VH4FIkhi9GnAe4QU5iCW
  0WqeK16/QyOjGbxrp7/sGNi1/EXV2b4HLgaleTulL/z2aC0XQbeMQrLIuBqK6Yaa
  3UfY4uTEr55Fh8pKZf5eji236yPK11whKYREHnj3yrFSkVemxii68ao1gH090E+4
  gYvw0TsUry7nnoy3sHs3E50O7e6/zYqsYqw9zQzIEVeOMDKMOhq3Tblt3A==
  =Ztti
  -----END PGP SIGNATURE-----
tags:
  - wordpress
  - cve
  - bezpieczeństwo
  - rce
  - wtyczki
title: 'Brak weryfikacji uprawnień w WP File Download: od usuwania plików do RCE'
toc: false
---

Wtyczka WP File Download dla WordPressa ma klasyczną, ale niezwykle groźną usterkę architektoniczną. CVE-2026-14982 przypisano ocenę CVSS v3.1 na poziomie 8.1. W zasadzie każdy zarejestrowany użytkownik – nawet z najniższą rolą Subskrybenta (Subscriber) – może wykorzystać kombinację dwóch akcji wtyczki do usunięcia dowolnego pliku na serwerze, do którego proces serwera www ma prawa zapisu.

Problem nie polega na jednym niedopatrzonym wywołaniu. To efekt braku elementarnych mechanizmów kontroli dostępu w logice biznesowej wtyczki.

## Mechanizm: dwuetapowe wstrzyknięcie ścieżki i unlink

Atak składa się z dwóch kroków, które wykorzystują dwa osobne punkty końcowe (endpoints) udostępniane przez wtyczkę.

W pierwszym kroku atakujący wysyła żądanie do zadania `file.save`. Celem tego kroku jest zapisanie w metadanych obiektu pliku zmodyfikowanej ścieżki zawierającej sekwencję przechodzenia przez katalogi (path traversal), na przykład `../../../../wp-config.php`. Kod odpowiedzialny za przetwarzanie parametru w zadaniu `file.save` nie sanitizuje należycie wejścia pod kątem niedozwolonych znaków w ścieżce i przyjmuje ciąg znaków bez dostatecznej walidacji.

W drugim kroku atakujący wywołuje zadanie `file.delete`. Wtyczka odczytuje wcześniej zapisane metadane i przekazuje zapisaną w nich ścieżkę bezpośrednio do natywnej funkcji PHP `unlink()`.

Co najważniejsze z punktu widzenia architektonicznego: oba te punkty końcowe nie weryfikują uprawnień użytkownika (`capability checks`), ani nie wymagają poprawnego tokenu zabezpieczającego (`nonce`). Oznacza to, że warstwa uprawnień WordPressa jest tu całkowicie pomijana. Wystarczy ważna sesja dowolnego zalogowanego użytkownika.

## Konsekwencje: usuwanie plików i przejście do RCE

Usunięcie dowolnego pliku (Arbitrary File Deletion) rzadko kończy się wyłącznie na utracie danych. W środowisku WordPressa jest to powszechnie znany wektor prowadzący bezpośrednio do zdalnego wykonania kodu (RCE).

Głównym celem w takich scenariuszach staje się plik `wp-config.php`. Jeśli proces serwera HTTP ma uprawnienia do jego usunięcia, wykonanie zadania `file.delete` ze wstrzykniętą ścieżką powoduje skasowanie konfiguracji bazy danych. Gdy `wp-config.php` znika z dysku, WordPress przechodzi w stan wstępnej instalacji.

W tym momencie dowolny niezalogowany użytkownik, który jako pierwszy odwiedzi stronę, może uruchomić instalator `/wp-admin/install.php`, podać własną bazę danych lub podpiąć się pod istniejącą i ustawić nowe konto administratora. Stamtąd droga do wgrania własnej wtyczki z powłoką (web shell) i przejęcia pełnej kontroli nad systemem operacyjnym jest już prosta.

Usunięcie innych plików, takich jak `.htaccess` czy kluczowe pliki motywów i wtyczek, prowadzi z kolei do natychmiastowej odmowy usługi (DoS) lub ujawnienia zawartości katalogów.

## Stan poprawek i zalecenia

Podatność dotyczy wszystkich dotychczasowych wersji wtyczki WP File Download. W udostępnionych danych analitycznych brakuje informacji o wydanej łatce naprawiającej ten problem, co oznacza, że luka w momencie zgłoszenia pozostawała bez oficjalnego patcha. Tego publicznie nie widać, czy wydawca zdołał już spatchować kod w najnowszym wydaniu.

Jeśli w Twojej instancji WordPressa włączona jest otwarta rejestracja użytkowników na poziom Subscriber, a wtyczka WP File Download znajduje się w katalogu `/wp-content/plugins/`, system jest bezpośrednio narażony na atak.

W tej sytuacji optymalną decyzją inżynieryjną jest:

1. Tymczasowe całkowite wyłączenie wtyczki WP File Download do czasu weryfikacji wydania poprawki przez dostawcę.
2. Jeśli wtyczka jest niezbędna biznesowo, natychmiastowe zablokowanie możliwości rejestracji nowych użytkowników oraz wyłączenie kont o poziomie Subscriber i wyższym, które nie są bezwzględnie zaufane.
3. Wdrożenie na poziomie Web Application Firewall (WAF) reguł blokujących żądania zawierające sekwencje `../` w parametrach kierowanych do punktów końcowych wtyczki.

Brak weryfikacji tokenów nonce i funkcji `current_user_can()` w funkcjach wykonujących operacje I/O na dysku to powtarzający się błąd w ekosystemie wtyczek. Dopóki kod nie sprawdza uprawnień przed wywołaniem `unlink()`, jedyną bezpieczną opcją jest wycięcie kodu z procesu wykonawczego.

## Źródła

- https://nvd.nist.gov/vuln/detail/CVE-2026-14982
