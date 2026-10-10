(function () {
  'use strict';
  var meta = document.querySelector('meta[name="ga-id"]');
  var banner = document.querySelector('.cookie-banner');
  if (!meta || !/^G-[A-Z0-9]+$/.test(meta.content) || !banner) return;

  var id = meta.content;
  var storageKey = 'blog-analytics-consent-v1';
  var maxAge = 180 * 24 * 60 * 60 * 1000;
  var memoryChoice = null;
  var transientChoice = false;
  var started = false;
  var opener = null;
  var closeButton = banner.querySelector('[data-cookie-close]');
  var status = banner.querySelector('[data-cookie-status]');

  window.dataLayer = window.dataLayer || [];
  window.gtag = function () { window.dataLayer.push(arguments); };
  window['ga-disable-' + id] = true;
  window.gtag('consent', 'default', {
    analytics_storage: 'denied', ad_storage: 'denied',
    ad_user_data: 'denied', ad_personalization: 'denied'
  });

  function readChoice() {
    if (transientChoice) return memoryChoice;
    var value;
    try { value = JSON.parse(localStorage.getItem(storageKey)); } catch (_) { return memoryChoice; }
    if (!value || value.version !== 1 || !['granted', 'denied'].includes(value.analytics) ||
        !Number.isFinite(value.savedAt) || value.savedAt > Date.now() || Date.now() - value.savedAt >= maxAge) {
      return null;
    }
    return value.analytics;
  }

  function saveChoice(choice) {
    memoryChoice = choice;
    try {
      localStorage.setItem(storageKey, JSON.stringify({ version: 1, analytics: choice, savedAt: Date.now() }));
      transientChoice = false;
      return true;
    } catch (_) { transientChoice = true; return false; }
  }

  function clearAnalyticsCookies() {
    var names = document.cookie.split(';').map(function (part) { return part.split('=')[0].trim(); });
    var hosts = location.hostname.split('.');
    var domains = [''];
    for (var i = 0; i < hosts.length - 1; i++) {
      domains.push(hosts.slice(i).join('.'));
    }
    var paths = ['/'];
    var segments = location.pathname.split('/').filter(Boolean);
    for (var j = 1; j <= segments.length; j++) {
      paths.push('/' + segments.slice(0, j).join('/'));
      paths.push('/' + segments.slice(0, j).join('/') + '/');
    }
    names.filter(function (name) { return /^(_ga(?:_|$)|_gid$|_gat(?:_|$))/.test(name); }).forEach(function (name) {
      domains.forEach(function (domain) {
        paths.forEach(function (path) {
          document.cookie = name + '=; Max-Age=0; path=' + path + (domain ? '; domain=' + domain : '') + '; SameSite=Lax';
        });
      });
    });
  }

  function applyChoice(choice) {
    var granted = choice === 'granted';
    window['ga-disable-' + id] = !granted;
    if (!granted) {
      clearAnalyticsCookies();
      if (started) {
        window.gtag('consent', 'update', {
          analytics_storage: 'denied', ad_storage: 'denied',
          ad_user_data: 'denied', ad_personalization: 'denied'
        });
      }
      return;
    }
    window.gtag('consent', 'update', {
      analytics_storage: 'granted', ad_storage: 'denied',
      ad_user_data: 'denied', ad_personalization: 'denied'
    });
    if (started) return;
    started = true;
    window.gtag('js', new Date());
    window.gtag('config', id, {
      allow_google_signals: false,
      allow_ad_personalization_signals: false,
      cookie_expires: 180 * 24 * 60 * 60,
      cookie_update: false
    });
    var script = document.createElement('script');
    script.async = true;
    script.src = 'https://www.googletagmanager.com/gtag/js?id=' + encodeURIComponent(id);
    document.head.appendChild(script);
  }

  function showBanner(choice, focus) {
    banner.hidden = false;
    closeButton.hidden = !choice;
    status.textContent = choice ? 'Aktualny wybór: analityka ' + (choice === 'granted' ? 'włączona.' : 'wyłączona.') : '';
    if (focus) banner.querySelector('[data-cookie-choice="denied"]').focus();
  }

  function hideBanner() {
    var returnFocus = banner.contains(document.activeElement);
    banner.hidden = true;
    if (returnFocus) {
      var target = opener || document.querySelector('[data-cookie-settings]');
      if (target) target.focus({ preventScroll: true });
    }
    opener = null;
  }

  document.querySelectorAll('[data-cookie-settings]').forEach(function (button) {
    button.hidden = false;
    button.addEventListener('click', function () {
      opener = button;
      showBanner(readChoice(), true);
    });
  });
  banner.querySelectorAll('[data-cookie-choice]').forEach(function (button) {
    button.addEventListener('click', function () {
      var choice = button.dataset.cookieChoice;
      var saved = saveChoice(choice);
      applyChoice(choice);
      hideBanner();
      /* Unload the active Google tag after withdrawal, including enhanced events. */
      if (choice === 'denied' && started && saved) location.reload();
    });
  });
  closeButton.addEventListener('click', hideBanner);
  banner.addEventListener('keydown', function (event) {
    if (event.key === 'Escape' && readChoice()) hideBanner();
  });

  function refreshChoice() {
    var choice = readChoice();
    applyChoice(choice);
    if (!choice) showBanner(null, false);
    else if (opener) showBanner(choice, false);
    else hideBanner();
    if (choice !== 'granted' && started && !transientChoice) location.reload();
  }
  window.addEventListener('storage', function (event) {
    if (event.key === storageKey || event.key === null) {
      memoryChoice = null;
      transientChoice = false;
      refreshChoice();
    }
  });
  window.addEventListener('pageshow', refreshChoice);
  window.addEventListener('focus', refreshChoice);
  document.addEventListener('visibilitychange', function () {
    if (!document.hidden) refreshChoice();
  });
  refreshChoice();
})();
