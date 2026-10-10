# Przyciski projektów

Dodaj adresy do nagłówka pliku projektu. Właściwe przyciski pojawią się
automatycznie na karcie projektu i na jego stronie. Nie ustawiaj adresu,
dopóki aplikacja lub repozytorium nie są dostępne.

```toml
appstore = "https://apps.apple.com/pl/app/lemricfit/id6791898733"
github = "https://github.com/Lemric/BatchRequest"
# googleplay = "adres opublikowanej aplikacji w Google Play"
```

Obsługiwane pola:

| Pole | Przycisk |
| --- | --- |
| `appstore` (również `app_store`) | Oficjalna polska odznaka App Store |
| `googleplay` (również `google_play` lub `playstore`) | Oficjalna polska odznaka Google Play |
| `github` | Przycisk „Zobacz kod / GitHub” z oficjalnym znakiem GitHuba |

## Logo i licencja

```toml
logo = "/images/projects/batch-request.webp"
logo_alt = "Logo BatchRequest"
license = "MIT"
license_url = "https://github.com/Lemric/BatchRequest/blob/main/LICENSE"
docs_url = "https://github.com/Lemric/BatchRequest#readme"
```

Logo pojawia się na liście projektów oraz w nagłówku strony projektu.
Grafika zachowuje proporcje i kolory w obu motywach. Używaj lokalnego
pliku z oficjalnych materiałów projektu. Bez pola `logo` pozostaje ikona
rodzaju projektu.

Pole `license` wyświetla nazwę licencji w obu widokach. Opcjonalne
`license_url` dodaje odnośnik do jej pełnego tekstu i jest również używane
w metadanych strony. Bez adresu nazwa jest zwykłym tekstem. Nie przypisuj
licencji tylko na podstawie publicznej dostępności repozytorium.
`docs_url` dodaje odnośnik do dokumentacji, a `composer` i `go_module`
wyświetlają polecenia instalacji na stronie projektu.

Wspólne szablony `project_brand.html` i `project_license.html` zapewniają
taki sam sposób prezentacji na kartach i stronach projektów.

Odznaki sklepów zachowują oryginalną grafikę, proporcje, czarne tło
i odstępy ochronne. Są przechowywane lokalnie; samo wyświetlenie przycisku
nie wysyła zapytania do sklepu. App Store jest pierwszy w kolejności.
Przyciski mają opis dostępny dla czytników ekranu i oznaczenie nowej karty.

Szablon `project_platforms.html` zbiera skonfigurowane adresy, a
`platform_badge.html` odpowiada za wygląd poszczególnych platform.
Nową platformę dodaj w obu szablonach wraz z jej lokalną grafiką.

## Źródła grafik

- [Apple — wytyczne i narzędzia marketingowe](https://developer.apple.com/app-store/marketing/guidelines/)
  — polska odznaka z `https://tools.applemediaservices.com/api/badges/download-on-the-app-store/black/pl-pl?size=250x83`.
- [Google Play — materiały marketingowe](https://play.google.com/intl/en_us/badges/)
  — polska odznaka z `https://play.google.com/intl/en_us/badges/static/images/badges/pl_badge_web_generic.png`.
- [GitHub Primer Octicons — mark-github](https://github.com/primer/octicons/blob/main/icons/mark-github-16.svg)
  — licencja MIT; kopia licencji znajduje się obok grafiki.

Odznaki sklepów podlegają zasadom użycia znaków Apple i Google.
Nie zmieniaj ich kolorów, napisów ani proporcji.
