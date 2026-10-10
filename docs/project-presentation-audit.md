# Audyt prezentacji projektów — 10 października 2026

Zakres: wszystkie cztery projekty, ich karty i strony szczegółowe.

| Element | Ustalenie | Wdrożenie |
| --- | --- | --- |
| Licencja | `license` była używana tylko w niewidocznym znaczniku `meta`. | Widoczna nazwa licencji w nagłówku i na karcie; opcjonalny link do pełnego tekstu przez `license_url`. |
| BatchRequest | Nagłówek danych wskazywał MIT, opis nadal GPL 3.0. Aktualne `composer.json` i `LICENSE` w repozytorium potwierdzają MIT. | Opis oraz widoczna etykieta MIT są spójne; link prowadzi do pliku LICENSE. |
| Logo | Szablon zawsze wyświetlał ogólną ikonę rodzaju projektu. | Wspólna obsługa `logo` i `logo_alt`, zachowanie proporcji oraz ikona zastępcza, gdy brak logo. |
| Dokumentacja BatchRequest | Czytelnik musiał przejść do repozytorium i sam znaleźć instrukcję. | Osobny przycisk do README projektu. |
| Instalacja | Composer i Go miały osobne, poprawne bloki instalacji. | Zachowane polecenia i przewijanie długiego kodu; polecenie w opisie BatchRequest sformatowane jako kod. |
| Linki platform | App Store, Google Play i GitHub są wspólnie obsługiwane. | Zachowane odznaki, dostępne opisy linków i oznaczenie nowej karty. |
| Dostępność | Logo i licencja nie miały widocznej prezentacji. | Opisy alternatywne logo, semantyczna etykieta licencji, fokus klawiatury i niezależny link licencji na karcie. |

## Materiały i zakres danych

- BatchRequest: oryginalny `doc/logo.webp` z lokalnego repozytorium
  BatchRequest. Licencję zweryfikowano w aktualnych publicznych plikach
  [composer.json](https://github.com/Lemric/BatchRequest/blob/main/composer.json)
  i [LICENSE](https://github.com/Lemric/BatchRequest/blob/main/LICENSE).
- LemricFit: logo z `fit.lemric.com/static/images/lemric-fit.png`.
- DataLeakTracker: znak z `dataleaktracker/public/assets/favicon.svg`.
- Event Dispatcher for Go: po wskazaniu przez właściciela repozytorium
  `lemric/eventdispatcher-go` zweryfikowano aktualny plik LICENSE przez
  GitHub API. Projekt ma licencję własnościową; dodano widoczną etykietę.
  Repozytorium nie jest dostępne anonimowo, więc nie dodano publicznego
  przycisku do kodu ani odnośnika licencji prowadzącego do błędu 404.
  Brak dedykowanego logo w sprawdzonych materiałach; zachowana ikona Go.
- Pozostałym projektom nie przypisano licencji, których nie podano lub
  nie można było jednoznacznie potwierdzić.

Grafiki są przechowywane w `static/images/projects/`. Wyświetlenie logo
nie wymaga zapytania do zewnętrznego serwisu.
