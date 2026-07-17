+++
title = "Polityka prywatności aplikacji LemricFit"
linkTitle = "polityka prywatności"
description = "Polityka prywatności aplikacji LemricFit: model local first, dane treningowe na urządzeniu, RODO."
date = 2026-07-17T12:00:00+02:00
draft = false
type = "page"
toc = true
+++

**Obowiązuje od:** 17 lipca 2026 r.

Niniejsza Polityka prywatności wyjaśnia, jakie dane użytkownika są przetwarzane w związku z korzystaniem z aplikacji mobilnej **LemricFit** („Aplikacja”) oraz w jaki sposób są one chronione. Tekst opiera się na faktycznym sposobie działania Aplikacji (model *local first*: dane treningowe pozostają na urządzeniu użytkownika, bez obowiązkowego konta w chmurze i bez wymaganego serwera aplikacji).

## 1. Administrator

Administratorem w rozumieniu rozporządzenia Parlamentu Europejskiego i Rady (UE) 2016/679 z dnia 27 kwietnia 2016 r. („RODO”), w zakresie, w jakim Aplikacja lub powiązane usługi powodują przetwarzanie danych osobowych poza wyłączną kontrolą użytkownika na urządzeniu, jest:

**Dominik Labudziński**  
Adres email do spraw związanych z prywatnością: **dominik@labudzinski.com**

W praktyce większość danych treningowych i profilowych jest przechowywana **wyłącznie lokalnie** na urządzeniu użytkownika, w zaszyfrowanej bazie danych Aplikacji. Administrator nie prowadzi obowiązkowego konta użytkownika w chmurze LemricFit i nie utrzymuje centralnej bazy treningów użytkowników.

## 2. Charakter Aplikacji (local first)

LemricFit jest lokalną aplikacją fitness. Oznacza to w szczególności, że:

1. **Nie jest wymagane** założenie konta u dostawcy Aplikacji ani logowanie do serwera LemricFit.
2. **Podstawowym źródłem danych** o treningach, trasach GPS, pomiarach z czujników, celach, rekordach i ustawieniach jest zaszyfrowana baza danych na urządzeniu użytkownika.
3. Aplikacja **nie oferuje** funkcji społecznościowych, klubów ani synchronizacji społecznej między użytkownikami.
4. Funkcje sieciowe (mapy, pogoda, geokodowanie, opcjonalna diagnostyka, opcjonalna kopia zapasowa w chmurze producenta systemu) są **ograniczone i celowe**. Opisano je w dalszych punktach.

## 3. Jakie dane są przetwarzane

### 3.1. Dane przechowywane lokalnie na urządzeniu

W zależności od tego, z których funkcji korzysta użytkownik, Aplikacja może zapisywać na urządzeniu m.in.:

* **Profil:** nazwa wyświetlana, zdjęcie profilowe (ścieżka do pliku lokalnego), masa ciała, wzrost, rok urodzenia, jednostki miary, cele tygodniowe, tętno maksymalne i spoczynkowe
* **Treningi:** rodzaj aktywności, czas trwania, dystans, tempo lub prędkość, status (np. pauza), podsumowania, odcinki treningu (*splits*)
* **Lokalizacja i trasa:** ciąg punktów GPS (współrzędne, wysokość, dokładność, znaczniki czasu) w trakcie i po nagraniu aktywności
* **Czujniki:** odczyty tętna, kadencji i mocy z urządzeń Bluetooth Low Energy oraz lokalne pomiary ruchu (m.in. krokomierz)
* **Zdrowie (opcjonalnie):** dane odczytane lub zapisane za pośrednictwem Apple Health (iOS) albo Health Connect (Android), zgodnie z nadanymi uprawnieniami systemu
* **Multimedia:** zdjęcia dodane do profilu lub aktywności (pliki lokalne)
* **Środowisko aktywności:** nazwa miejsca (z geokodowania) oraz temperatura i opis pogody zapisane przy zakończeniu treningu
* **Bezpieczeństwo:** skrót (hash) kodu PIN blokady aplikacji, klucz odzyskiwania (fraza 12 słów, przechowywana zgodnie z mechanizmami bezpieczeństwa Aplikacji), klucz szyfrujący bazę (DEK) w systemowym Keychain (iOS lub macOS) lub Keystore (Android)
* **Kopie zapasowe:** zaszyfrowane kopie bazy (bez klucza DEK) w lokalnym katalogu kopii oraz, opcjonalnie, w chmurze, zgodnie z opisem poniżej
* **Ustawienia i dzienniki techniczne:** preferencje Aplikacji, lokalne wpisy dotyczące integralności i synchronizacji (bez umieszczania w nich jawnych współrzędnych w logach diagnostycznych)

**Aplikacja nie wymaga** podania adresu email, numeru telefonu ani danych płatniczych do korzystania z podstawowych funkcji treningowych.

### 3.2. Dane przekazywane poza urządzenie (tylko w określonych sytuacjach)

* **Wyświetlanie map (Android):** żądania kafelków mapy / użycie zestawu SDK Map Google zgodnie z konfiguracją systemu i kluczem API → Google (Maps SDK)
* **Wyświetlanie map (iOS lub macOS):** renderowanie mapy przez Apple MapKit → Apple (w ramach systemu)
* **Geokodowanie nazwy miejsca przy końcu treningu:** współrzędne geograficzne → systemowe API geokodowania platformy
* **Pogoda przy końcu treningu:** szerokość i długość geograficzna → serwis Open Meteo
* **Apple Health / Health Connect:** treningi, tętno i powiązane typy danych wskazane w uprawnieniach → Apple Health albo Health Connect (dane pozostają w ekosystemie wybranym przez użytkownika)
* **Kopia zapasowa w chmurze (iOS):** zaszyfrowany plik kopii **bez** klucza DEK → iCloud Drive (przestrzeń Aplikacji), jeśli użytkownik ma włączoną usługę iCloud
* **Kopia zapasowa w chmurze (Android):** zaszyfrowany plik kopii **bez** klucza DEK → lokalna kopia na urządzeniu oraz, po dobrowolnym połączeniu konta Google, obszar danych aplikacji Google Drive (*app data*)
* **Opcjonalna diagnostyka awarii:** zdarzenia błędów po wcześniejszym usunięciu lub zamaskowaniu danych wrażliwych; bez domyślnego przesyłania danych osobowych identyfikujących → Sentry, **wyłącznie** gdy w danej kompilacji włączono identyfikator DSN
* **Udostępnianie przez użytkownika:** pliki lub treści wybrane przez użytkownika (np. podsumowanie, zdjęcie, klucz odzyskiwania) → aplikacja lub usługa wskazana przez użytkownika w systemowym oknie udostępniania

Aplikacja **nie sprzedaje** danych użytkowników i **nie wyświetla reklam** opartych na profilowaniu behawioralnym.

## 4. Cele i podstawy prawne

W zakresie, w jakim ma zastosowanie RODO:

* **Świadczenie lokalnych funkcji Aplikacji** (nagrywanie treningu, mapa, statystyki, cele, blokada PIN): wykonanie umowy o korzystanie z Aplikacji / niezbędność do świadczenia usługi na żądanie użytkownika (art. 6 ust. 1 lit. b RODO) oraz, w odniesieniu do danych o zdrowiu i lokalizacji, **zgoda** wyrażona poprzez systemowe uprawnienia (art. 6 ust. 1 lit. a oraz art. 9 ust. 2 lit. a RODO, o ile mają zastosowanie)
* **Zapis treningu do Apple Health / Health Connect:** **zgoda** użytkownika w systemie Health / Health Connect
* **Pogoda i nazwa miejsca przy podsumowaniu:** prawnie uzasadniony interes polegający na wzbogaceniu lokalnego podsumowania treningu (art. 6 ust. 1 lit. f RODO) albo zgoda wynikająca z korzystania z funkcji wymagającej lokalizacji; współrzędne są wysyłane wyłącznie w tym celu
* **Mapy:** niezbędność do wyświetlenia trasy / przeglądania mapy; na Androidzie także warunki usług Google dotyczące Maps SDK
* **Kopia w iCloud / Google Drive:** **zgoda** / decyzja użytkownika (włączenie iCloud albo połączenie konta Google w ustawieniach kopii)
* **Diagnostyka Sentry** (jeśli włączona w kompilacji): prawnie uzasadniony interes polegający na stabilności Aplikacji (art. 6 ust. 1 lit. f RODO), przy usuwaniu lub maskowaniu danych wrażliwych

Użytkownik może w każdej chwili cofnąć zgody systemowe (Lokalizacja, Bluetooth, Zdrowie, Ruch, Aparat, Powiadomienia, Face ID / biometria) w ustawieniach systemu operacyjnego. Cofnięcie zgody może uniemożliwić działanie danej funkcji.

## 5. Uprawnienia systemu operacyjnego

Aplikacja prosi o dostęp wyłącznie do funkcji, których potrzebuje:

* **Lokalizacja** (w tym w tle): nagrywanie trasy GPS przy zablokowanym ekranie
* **Bluetooth**: czujniki tętna, kadencji i mocy
* **Ruch / rozpoznawanie aktywności**: propozycja rozpoczęcia treningu oraz wspomaganie pauzy automatycznej
* **Zdrowie** (Apple Health / Health Connect): odczyt i zapis wskazanych typów danych treningowych
* **Aparat i zdjęcia**: zdjęcie profilowe oraz multimedia aktywności
* **Powiadomienia**: stan treningu (w tym powiadomienie usługi działającej na pierwszym planie na Androidzie oraz Live Activity na iOS)
* **Biometria**: odblokowanie blokady Aplikacji (PIN + Face ID / odcisk palca)

Bez przyznania uprawnienia dana funkcja nie jest realizowana albo działa w ograniczonym zakresie.

## 6. Szyfrowanie, kopie zapasowe i bezpieczeństwo

1. Baza danych treningowych jest **szyfrowana**. Klucz szyfrujący (DEK) jest przechowywany w **Keychain** (Apple) lub **Keystore** (Android) i **nie jest dołączany** do plików kopii zapasowej.
2. Kopia zapasowa bez DEK jest praktycznie bezużyteczna dla osoby trzeciej, która nie posiada klucza z urządzenia użytkownika (lub poprawnie wprowadzonego klucza odzyskiwania, jeśli użytkownik go skonfigurował).
3. Blokada Aplikacji (PIN, opcjonalnie biometria) chroni dostęp do treści Aplikacji na odblokowanym urządzeniu.
4. Logi diagnostyczne są oczyszczane z danych wrażliwych (m.in. współrzędne, tętno, PIN, DEK i fraza odzyskiwania nie są celowo wysyłane w postaci jawnej do Sentry).
5. Na Androidzie wyłączono systemowy Auto Backup treści Aplikacji w sposób uniemożliwiający automatyczne wyciągnięcie zaszyfrowanej bazy przez mechanizmy kopii producenta systemu w domyślnej konfiguracji Aplikacji.
6. Użytkownik może wykonać **lokalne usunięcie wszystkich danych** z poziomu Aplikacji. Obejmuje to bazę, klucze, PIN oraz lokalne i (w miarę dostępności) chmurowe kopie zapasowe zarządzane przez Aplikację.

## 7. Okres przechowywania

* Dane lokalne są przechowywane **do czasu ich usunięcia przez użytkownika**, odinstalowania Aplikacji (z zastrzeżeniem kopii systemowych lub chmurowych poza Aplikacją, których Administrator nie kontroluje w pełni) albo wykonania funkcji usunięcia wszystkich danych lokalnych.
* Zaszyfrowane lokalne i chmurowe kopie zapasowe istnieją, dopóki użytkownik ich nie usunie albo nie wyczyści powiązanego konta lub magazynu (iCloud, Google Drive).
* Dane w Apple Health / Health Connect podlegają polityce i ustawieniom Apple albo Google oraz decyzjom użytkownika w tych systemach.
* Dane przekazane do Open Meteo lub dostawców map są przetwarzane zgodnie z ich własnymi politykami; Aplikacja nie prowadzi u Administratora archiwum tych zapytań.

## 8. Odbiorcy i powierzenie

Administrator nie sprzedaje danych. Odbiorcami mogą być wyłącznie:

* **dostawcy systemu operacyjnego i usług platformowych** (Apple, Google), w zakresie map, zdrowia, biometrii, powiadomień, iCloud lub Google Drive,
* **Open Meteo**, w zakresie jednorazowych zapytań o pogodę,
* **Sentry**, wyłącznie gdy dana kompilacja Aplikacji ma włączoną diagnostykę,
* **osoby lub usługi wskazane przez użytkownika** przy ręcznym udostępnianiu treści.

W zakresie usług tych podmiotów zastosowanie mają również ich regulaminy i polityki prywatności.

## 9. Przekazywanie danych poza Europejski Obszar Gospodarczy

Niektóre podmioty (np. Google, Apple, Sentry, Open Meteo) mogą przetwarzać dane na serwerach poza EOG. W takich przypadkach przekazywanie odbywa się w oparciu o mechanizmy przewidziane w RODO (m.in. standardowe klauzule umowne lub decyzje stwierdzające odpowiedni stopień ochrony, w zależności od dostawcy i aktualnego stanu prawnego). Korzystanie z funkcji map, pogody, Health, Drive/iCloud lub opcjonalnego Sentry oznacza akceptację tego ryzyka w zakresie niezbędnym do działania danej funkcji.

## 10. Prawa użytkownika

W zakresie, w jakim Administrator przetwarza dane osobowe użytkownika, użytkownikowi przysługują, stosownie do RODO, prawa do:

* dostępu do danych,
* sprostowania danych,
* usunięcia danych („prawo do bycia zapomnianym”),
* ograniczenia przetwarzania,
* przenoszenia danych (w praktyce: eksport / udostępnienie lokalnych treści przez użytkownika oraz usunięcie danych w Aplikacji),
* wniesienia sprzeciwu wobec przetwarzania opartego na art. 6 ust. 1 lit. f RODO,
* cofnięcia zgody w dowolnym momencie (bez wpływu na zgodność z prawem przetwarzania przed jej cofnięciem),
* wniesienia skargi do Prezesa Urzędu Ochrony Danych Osobowych (ul. Stawki 2, 00 193 Warszawa).

Ponieważ większość danych znajduje się wyłącznie na urządzeniu użytkownika, najskuteczniejszą formą realizacji prawa do usunięcia jest:

1. usunięcie danych w ustawieniach Aplikacji (funkcja usunięcia wszystkich danych lokalnych), oraz
2. odinstalowanie Aplikacji, oraz
3. usunięcie uprawnień i danych w Apple Health / Health Connect, iCloud lub Google Drive, według potrzeb użytkownika.

W sprawach, które wymagają działania Administratora (np. dane w opcjonalnym Sentry), należy napisać na adres wskazany w pkt 1.

## 11. Dzieci

Aplikacja nie jest skierowana do dzieci poniżej 16. roku życia. Administrator nie zbiera świadomie danych dzieci. W razie powzięcia informacji o takim zbieraniu dane zostaną usunięte w zakresie, w jakim Administrator nimi dysponuje.

## 12. Brak zautomatyzowanego podejmowania decyzji o skutkach prawnych

Aplikacja nie podejmuje wobec użytkownika zautomatyzowanych decyzji wywołujących skutki prawne w rozumieniu art. 22 RODO. Sugestie rozpoczęcia treningu, klasyfikacja aktywności czy statystyki treningowe mają charakter pomocniczy i lokalny.

## 13. Zmiany Polityki

Polityka może być aktualizowana, gdy zmieni się sposób działania Aplikacji, zakres integracji lub obowiązujące przepisy. Zaktualizowana wersja zostanie opublikowana wraz z datą obowiązywania. W razie istotnych zmian Administrator dołoży starań, aby poinformować użytkowników w Aplikacji lub w materiałach dystrybucyjnych (App Store / Google Play), o ile będzie to możliwe bez obowiązkowego konta w chmurze.

## 14. Kontakt

Pytania dotyczące prywatności: **dominik@labudzinski.com**.

*Niniejszy dokument ma charakter informacyjny i odzwierciedla działanie Aplikacji LemricFit na dzień wskazany na wstępie. Nie stanowi porady prawnej.*
